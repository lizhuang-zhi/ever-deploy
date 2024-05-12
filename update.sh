#!/bin/sh -xe
echo $PATH
cd /server2  
rm -rf gin-server  # 删除之前目录中构建的gin-server应用
/usr/local/go/bin/go build -o gin-server main.go || { echo "Go build failed"; exit 1; }  # 构建最新版本的服务

# 等待一会儿
sleep 5

######### 开始不停服更新 ############

echo "Deploying new version of the application...."

HOST=10.1.27.86  # 执行部署的服务器地址

######### 部署server1 ############
echo "Update server1 project..."

# 设置nginx权重（server1无权重）
scp ./nginx/server.1.conf ${HOST}:/etc/nginx/conf.d/server.conf

# 关闭server1进程
ssh ${HOST} ' nginx -t; systemctl reload nginx; supervisorctl stop prod:server1'
# 更新server1内容
scp ./server2/gin-server ${HOST}:/home/ever-deploy/server1/
# 开启server1进程
ssh ${HOST} ' supervisorctl start prod:server1'

######### 部署server2 ############
echo "Update server2 project..."

# 设置nginx权重（server2无权重）
scp ./nginx/server.2.conf ${HOST}:/etc/nginx/conf.d/server.conf
# 关闭server2进程
ssh ${HOST} ' nginx -t; systemctl reload nginx; supervisorctl stop prod:server2'
# 更新server2内容
scp ./server2/gin-server ${HOST}:/home/ever-deploy/server2/
# 开启server2进程
ssh ${HOST} ' supervisorctl start prod:server2'

# 恢复nginx权重
scp ./nginx/server.conf ${HOST}:/etc/nginx/conf.d/server.conf
ssh ${HOST} ' nginx -t; systemctl reload nginx'

echo "Update complete."