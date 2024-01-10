echo "设置权限"
chmod 777 ./es/data
chmod +x ../main

\cp settings.yaml ../

echo "安装docker-compose"

chmod +x ./docker/docker-compose
\cp ./docker/docker-compose /usr/local/bin/docker-compose

echo "执行docker-compose up -d"
docker-compose up -d

echo "移动nginx配置文件"

\cp ./nginx.conf ../../


echo "下载supervisor"

# 安装第三方源，CentOS本身的yum源中没有supervisor，需要更换第三方源：
yum install epel-release -y
# 安装supervisor：
yum install -y supervisor

# 启动supervisor 服务、查看supervisor 服务状态
systemctl start supervisord
# 开机启动
systemctl enable supervisord