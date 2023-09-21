current_datetime=$(date '+%F %T')
set password "your_password"

git add .
git commit -m "deploy $current_datetime"
git push origin master
ssh ys@192.168.2.228 "cd /home/ys/ai4u-api-main/ ;git pull origin master"
expect "Password:"

send "$password\r"

expect eof