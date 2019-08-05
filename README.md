# Sheep服务端Golang(Gin）

- [Sheep](https://github.com/nEdAy/Sheep/tree/master) —— Android(Kotlin)

## 前期准备

**云主机：** 阿里云、腾讯云、[Vultr](https://www.vultr.com/?ref=8113323-4F) —— 47.93.228.168

**域名+证书 + DNS：**[新网](http://xinnet.com/)  + [Letsencrypt.org](https://letsencrypt.org/) 、 [Buypass.com](https://www.buypass.com/)  + DNS —— https://www.neday.cn/ping

**系统及服务 ：** CentOS + [LNMP(Nginx/MySQL/PHP)](https://lnmp.org/)  + Redis

​	[LAMP(Apache/MySQL/PHP)](https://lamp.sh) 、[LNMPA(Nginx/MySQL/PHP/Apache)](https://lnmp.org/lnmpa.html)

​ [PHP探针](http://47.93.228.168/p.php)

​ [phpmyadmin](http://47.93.228.168/phpmyadmin/ )  + Navicat Premium 12

​	[数据库无法远程访问](https://bbs.vpser.net/thread-13563-1-1.html)： 云服务器安全组 + CentOS防火墙iptables + phpmyadmin user 配置 %

​ ssh连接工具：** Xshell 6 、MobaXterm

​	常用第三方包指令： lrzsz(rz,sz) 、[screen](https://linuxize.com/post/how-to-use-linux-screen/)

[大淘客API](http://www.dataoke.com/pmc/api-market.html) <——> [Shepherd Swagger](https://www.neday.cn/swagger/index.html) <——> [Sheep](https://github.com/nEdAy/Confidence/tree/master/Sheep)


main.go  

- build —— *.ini, ssl, 可执行文件
- controllor
- docs
- middleware
  - jwt
- model —— [GORM](https://jasperxu.github.io/gorm-zh/)
- pkg
  - logger
  - redis
  - config
  - jwt
  - scrypt
  - response
- router

[MySQL命名、设计及使用规范](https://www.biaodianfu.com/mysql-best-practices.html)

Golang交叉编译 & Makefile

supervisor 管理进程 —— [http://47.93.228.168:9001](http://47.93.228.168:9001/)

[存储密码方案](https://astaxie.gitbooks.io/build-web-application-with-golang/zh/09.5.html)

hash 、 hash + salt  、dk := scrypt.Key([]byte("password"), []byte("salt"), 16384, 8, 1, 32)

