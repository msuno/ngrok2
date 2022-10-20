# 部署步骤

### 下载并编译
```
git clone https://gitee.com/msuno/ngrok2.git
make all
```

### 构建容器
```
docker build -t ngrok:latest
```

### 运行容器
```
mkdir ~/.ngrok
docker run -p 8001:8001 -p 8002:8002 -p 4443:4443 -p 8000:8000 -v ~/.ngrok/info.log:/log -d --name "ngrok" ngrok:1.1
```