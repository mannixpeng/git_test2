.PHONY: clean all build run gotool test  conver help
BINARY="bluebell"
all: gotool build
build:
	CGO_ENALBED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}
run:
	@go run ./server/server.go
	@go run ./client/clieng.go
gotool:
	go fmt ./
	go vet ./
test:
	go test ./...
clean:
	@if [ -f ${BINARY} ]; then rm -rfv ${BINARY} ${BINARY}.html ${BINARY}.out; fi
cover:
	go test -coverprofile ${BINARY}.out
	go tool cover -html=${BINARY}.out -o ./${BINARY}.html
docker:
#    @docker build -t leo/hello:latest .

help:
	@echo "make - 格式化GO代码，并编译生成二进制文件"
	@echo "make build - 编译GO代码，生成二进制文件"
	@echo "make run - 直接运行GO代码"
	@echo "make clean - 移除二进制文件合vim swap files"
	@echo "make gotool - 运行GO工具 'fmt' and 'vet'"
	@echo "make lint 执行代码检查"
	@echo "make docker 构建docker镜像"