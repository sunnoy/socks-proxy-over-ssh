# socks proxy over ssh

Many intranet penetration projects have been blocked by firewalls. 
What is certain is that the ssh protocol will not be intercepted. 
So this project uses the ssh protocol as the sock channel


# build

```bash
go mod tidy


#### cross compile ####
# Linux build
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

# macos build
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go


# windows build
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go


```


# Windows build need cgo

```bash
# 1
https://sourceforge.net/projects/mingw-w64/files/Toolchains%20targetting%20Win64/Personal%20Builds/mingw-builds/8.1.0/

# 2
# download x86_64-posix-seh 

# 3 move to c disk
C:\mingw64

# 4
# add system env
```


# run

```bash
# linux macos
export SSH_HOST="you ssh server ip"
export SSH_PORT="you ssh server port"
export SSH_PW="you ssh server password"
export SOCKS_L="you local socks serer"
export SSH_L="you remote ssh server listen port map to local socks server listen"


# windows
set SSH_HOST="you ssh server ip"
set SSH_PORT="you ssh server port"
set SSH_PW="you ssh server password"
set SOCKS_L="you local socks serer"
set SSH_L="you remote ssh server listen port map to local socks server listen"


go run 
```

