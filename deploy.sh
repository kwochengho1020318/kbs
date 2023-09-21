current_datetime=$(date '+%F %T')
git add .
git commit -m "deploy $current_datetime"
git push origin master