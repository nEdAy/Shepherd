# Sheep服务端Golang(Gin）

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