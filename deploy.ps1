$current_datetime = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
$password = "pass"

git add .
git commit -m "deploy $current_datetime"
git push origin master

$sshCommand = @"
cd /home/ys/ai4u-api-main/
git pull origin master
docker build -t lucatest .
docker-compose up -d
"@
ssh ys@192.168.2.228 $sshCommand

Write-Host "deploy finished"
