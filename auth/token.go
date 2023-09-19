package auth

import (
	"fmt"
	"main/config"
	"main/gojdb"
	"os"
	"strconv"

	"github.com/hunterhug/gosession"
	"github.com/hunterhug/gosession/kv"
)

func GetTokenManage() (gosession.TokenManage, error) {
	redisconf := config.NewConfig("appsettings.json").Redis
	boolValue, _ := strconv.ParseBool(os.Getenv("DOCKER_CONTAINER"))
	fmt.Println(boolValue)
	var redisConfig *kv.MyRedisConf
	if boolValue {
		redisConfig = gosession.NewRedisSessionSingleModeConfig(redisconf.DockerAddr, redisconf.Db, redisconf.Password)
	} else {
		redisConfig = gosession.NewRedisSessionSingleModeConfig(redisconf.Addr, redisconf.Db, redisconf.Password)
	}
	tokenManage, err := gosession.NewRedisSession(redisConfig)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	tokenManage.ConfigDefaultExpireTime(600)
	tokenManage.ConfigUserKeyPrefix("go-user")
	tokenManage.ConfigTokenKeyPrefix("go-token")

	fn := func(id string) (user *gosession.User, err error) {
		db := gojdb.NewGOJDB()
		param := make(map[string][]string)
		param["user_id"] = []string{id}
		bureau, _ := db.QueryData("Select * from Users where user_id =@user_id", param)
		return &gosession.User{
			Id:     id,
			Detail: bureau,
		}, nil
	}
	tokenManage.ConfigGetUserInfoFunc(fn)
	return tokenManage, nil
}

func SetAndGettoken(userid string) (string, error) {
	tokenManage, err := GetTokenManage()
	if err != nil {
		return "", err
	}
	var tokenExpireTimeAlone int64 = 600
	token, err := tokenManage.SetToken(userid, tokenExpireTimeAlone)
	tokenManage.AddUser(userid, tokenExpireTimeAlone)
	if err != nil {

		return "", err
	}
	return token, err
}
func CheckTokenExists(token string) (*gosession.User, bool, error) {
	TokenManage, err := GetTokenManage()
	if err != nil {

		return nil, false, err
	}
	var userExpireTimeAlone int64 = 10
	u, exists, err := TokenManage.CheckTokenOrUpdateUser(token, userExpireTimeAlone)
	if u == nil {
		return nil, false, err
	} else {

		if err != nil {
			panic(err)
		}
	}
	if err != nil {
		fmt.Println(err.Error())
		return nil, false, err
	}
	return u, exists, err
}
