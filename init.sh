#!bin/sh

echo "Write your repository without https://, example format: <account_name>/<repo_name>.git"
read a

export GOPATH=$HOME/go             
export PATH=$PATH:/usr/local/go/bin
rm ./go.mod
go mod init $a

