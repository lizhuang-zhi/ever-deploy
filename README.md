# ever-deploy
不停服更新demo

# 描述
server1目录下构建的gin-server 模拟v1版本的gin-server服务
server2目录下构建的gin-server 模拟v2版本的gin-server服务

v1版本已经部署到第二台虚拟机的11110和11111端口
现在通过本项目重新打包构建v2版本并不停服更新到11110和11111端口

通过update.sh完成不停服更新步骤

> tip: 该仓库省略在虚拟机10.1.27.86上配置supervisor