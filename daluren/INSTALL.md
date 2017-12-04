# Build
## 在centos7.4上构建
* 安装golang1.9.2
```bash
https://redirector.gvt1.com/edgedl/go/go1.9.2.linux-amd64.tar.gz
```
* 配置golang
```bash
tar xvzf go1.9.2.linux-amd64.tar.gz
mkdir ~/gopkg
cat .bashrc

...
export GOPATH=~/gopkg
export PATH=${GOPATH}/bin/:~/go/bin/:PATH
...

source .bashrc
```
* 安装 nodejs8.9.1
```bash
wget https://nodejs.org/dist/v8.9.1/node-v8.9.1-linux-x64.tar.xz
```
* 配置 nodejs
```bash
tar xvJf node-v8.9.1-linux-x64.tar.xz

cat .bashrc
...
export GOPATH=~/gopkg
export PATH=~/node-v8.9.1-linux-x64/bin/:${GOPATH}/bin/:~/go/bin/:PATH
...

source .bashrc
```
* 安装git
```bash
yum install git -y
```
* 安装bee及bee命令

bee 框架版本为1.9.1
```bash
go get github.com/astaxie/beego
go get github.com/beego/bee
```
* 安装polymer、bower
```bash
npm --registry https://registry.npm.taobao.org install  -g polymer-cli
npm --registry https://registry.npm.taobao.org install  -g bower
```
* 将源代码复制到goapth中
```bash
cd $GOPATH
git clone https://dfsfds
cd gitclone_dir
cp -r * ~/gopkg/src/
```
* 安装polymer元件、公历转农历js库
```bash
cd $GOPATH/src/daluren/static && bower install
npm install lunar-calendar
```
* build polymer app
```bash
polymer build
```
* 调整bee 的static目录
```bash
mv  build/ /tmp/
rm -rf *
cp -r /tmp/build/default/* .
```
* 调整 bee的运行模式
```bash
cd  ../conf/
cat app.conf
...
runmode = prod
...
```
* 发布bee app ,得到daluren.tar.gz
```bash
bee  pack
scp daluren.tar.gz username@yourserver
```

## 在windows10 下使用golang 64位构建32位版本
* 安装golang1.9.2
```bash
https://redirector.gvt1.com/edgedl/go/go1.9.2.windows-amd64.msi
```

* 在win10的bash中安装 nodejs8.9.1
```bash
wget https://nodejs.org/dist/v8.9.1/node-v8.9.1-linux-x64.tar.xz
```
* 在win10的bash中 配置 nodejs
```bash
tar xvJf node-v8.9.1-linux-x64.tar.xz

cat .bashrc
...
export GOPATH=~/gopkg
export PATH=~/node-v8.9.1-linux-x64/bin/:${GOPATH}/bin/:~/go/bin/:PATH
...

source .bashrc
```
* 在win10的bash中安装git
```bash
apt-get install git -y
```
* 在win10中安装bee及bee命令

bee 框架版本为1.9.1
```bash
# 在cmd中执行下面的命令
go get github.com/astaxie/beego
go get github.com/beego/bee
```
* 在win10的bash中安装polymer、bower
```bash
npm --registry https://registry.npm.taobao.org install  -g polymer-cli
npm --registry https://registry.npm.taobao.org install  -g bower
```
* 将源代码复制到goapth中,以下命在powershell中执行
```bash
cd %GOPATH%
git clone https://dfsfds
cd gitclone_dir
cp -r * ~/gopkg/src/
```
* 在win10的bash中安装polymer元件、公历转农历js库
```bash
cd $GOPATH/src/daluren/static && bower install
npm install lunar-calendar
```
* build polymer app
```bash
polymer build
```
* 调整bee 的static目录,在win10的bash中
```bash
mv  build/ /tmp/
rm -rf *
cp -r /tmp/build/default/* .
```
* 调整 bee的运行模式
```bash
cd  ../conf/
cat app.conf
...
runmode = prod
...
```
* 发布bee app ,得到daluren.tar.gz，在cmd中执行
```bash
set GOARCH=386
bee  pack
scp daluren.tar.gz username@yourserver
```

# 使用
* 解压daluren.tar.gz
* 运行daluren
* 在浏览器访问http://127.0.0.1:8080