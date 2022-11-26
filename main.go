package main

import (
	"github.com/things-go/go-socks5"
	"log"
	"os"
	"sync"
)

var (
	sshHost     = ""
	sshPort     = "22"
	sshUser     = "root"
	sshPassword = ""

	socks5Listen   = "localhost:1989"
	socks5User     = ""
	socks5Password = ""

	sshListen = "0.0.0.0:3000"
)

func main() {
	getEnv()

	wg := sync.WaitGroup{}
	wg.Add(2)

	// do ray sock5 proxy
	go func() {
		credentials := socks5.StaticCredentials{
			socks5User: socks5Password,
		}

		cator := socks5.UserPassAuthenticator{Credentials: credentials}

		// Create a SOCKS5 server
		server := socks5.NewServer(
			socks5.WithAuthMethods([]socks5.Authenticator{cator}),
		)

		log.Println("开始启动服务端！")

		if err := server.ListenAndServe("tcp", ":1989"); err != nil {
			panic(err)
		}

		wg.Done()
	}()
}

// get args from env
func getEnv() {

	sshHost = os.Getenv("SSH_HOST")
	sshPort = os.Getenv("SSH_PORT")
	sshPassword = os.Getenv("SSH_PW")
	socks5Listen = os.Getenv("SOCKS_L")
	socks5User = os.Getenv("SOCKSU")
	socks5Password = os.Getenv("SOCKSPW")
	sshListen = os.Getenv("SSH_L")

}
