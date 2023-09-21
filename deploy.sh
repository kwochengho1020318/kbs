current_datetime=$(date '+%F %T')
password = "pass"
git pull origin simplify
git add .
git commit -m "deploy $current_datetime"
git push origin simplify
ssh_address="ys@192.168.2.228"

# 远程目录
remote_directory="/home/ys/ai4u/user1-project1"

# 检查目录是否存在
ssh $ssh_address "[ -d '$remote_directory' ]"
if [ $? -eq 0 ]; then
    # 如果目录存在，执行一系列命令
    ssh $ssh_address "cd '$remote_directory' && git pull origin master && docker build -t lucatest . && docker-compose up -d"
    echo "Commands executed successfully."
else
    echo "Directory does not exist on the remote server."
fiecho "deploy finished"
