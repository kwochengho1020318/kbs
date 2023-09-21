current_datetime=$(date '+%F %T')
password = "pass"

git add .
git commit -m "deploy $current_datetime"
git push origin master
ssh ys@192.168.2.228 "cd /home/ys/ai4u-api-main/ ;git pull origin master"<<EOF
$password
EOF
