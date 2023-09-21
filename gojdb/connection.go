package gojdb

import (
	"database/sql"
	"errors"
	"fmt"
	"main/config"
	"os"
	"strconv"
	"sync"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbInstance  *sql.DB
	InstanceErr error
	once        sync.Once
)

func SqlConnection() (*sql.DB, *error) {
	once.Do(func() {
		myconfig := config.NewConfig("appsettings.json")
		boolValue, _ := strconv.ParseBool(os.Getenv("DOCKER_CONTAINER"))
		var connString string
		if boolValue {
			connString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
				myconfig.Database.DockeHost,
				myconfig.Database.User,
				myconfig.Database.Password,
				myconfig.Database.Port,
				myconfig.Database.DB_Name)
		} else {
			connString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
				myconfig.Database.Server,
				myconfig.Database.User,
				myconfig.Database.Password,
				myconfig.Database.Port,
				myconfig.Database.DB_Name)
		}
		db, err := sql.Open(myconfig.Database.Driver, connString)

		// See "Important settings" section.
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
		db.Ping()
		dbInstance = db
		if err != nil {
			InstanceErr = errors.New("DB Error")
		}

	})
	return dbInstance, &InstanceErr
}
