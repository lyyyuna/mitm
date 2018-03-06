os=$1
arch=$2

export GOPATH=`pwd`
GOOS=$os GOARCH=$arch go build -o bin/httpmitm$os$arch src/main/*.go
