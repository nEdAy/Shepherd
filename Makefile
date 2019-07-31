.PHONY: default build make windows_build_386 linux_build clean

default:
	@echo 'Usage of make: [ build | clean | windows_build | linux_build ]'

build:
	@go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`' -X main.GIT_HASH=`git rev-parse HEAD`" -o ./build/Shepherd ./

linux_build:
	@rm -f ./build/Shepherd
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`' -X main.GIT_HASH=`git rev-parse HEAD` -s" -o ./build/Shepherd ./

windows_build:
	@rm -f ./build/Shepherd.exe
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`' -X main.GIT_HASH=`git rev-parse HEAD` -s" -o ./build/Shepherd.exe ./

windows_build_386:
	@rm -f ./build/Shepherd.exe
	@CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`' -X main.GIT_HASH=`git rev-parse HEAD` -s" -o ./build/Shepherd.exe ./

clean: 
	@rm -f ./build/Shepherd
	@rm -f ./build/Shepherd.exe
	@rm -f ./build/logs/*.log