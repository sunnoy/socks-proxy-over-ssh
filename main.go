package main

import (
	"fmt"
	"github.com/blacknon/go-sshlib"
	"github.com/things-go/go-socks5"
	"golang.org/x/crypto/ssh"
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

	go func() {

		// Create sshlib.Connect
		con := &sshlib.Connect{}

		// Create ssh.AuthMethod
		authMethod := sshlib.CreateAuthMethodPassword(sshPassword)

		// todo 添加key认证
		keyAuthMethod, _ := sshlib.CreateAuthMethodPublicKey("key file path", "key password")

		if keyAuthMethod != nil {
			err := con.CreateClient(sshHost, sshPort, sshUser, []ssh.AuthMethod{keyAuthMethod})
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		// Connect ssh server
		if authMethod != nil {
			err := con.CreateClient(sshHost, sshPort, sshUser, []ssh.AuthMethod{authMethod})
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		// PortForward
		err := con.TCPRemoteForward(socks5Listen, sshListen)
		if err != nil {
			log.Println(err)
		}

		// Set terminal log
		//con.SetLog(termlog, false)

		// Create session
		session, err := con.CreateSession()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		//data, _ := session.CombinedOutput("watch -n 1 'netstat -antp | grep ssh | grep 3000' | tail -f -")
		//
		//if err != nil {
		//	log.Fatalf("failed to call CombinedOutput(): %v", err)
		//}
		//log.Printf("output: %s", data)

		log.Println("开启启动加密隧道！")
		// Start ssh shell
		err = con.CmdShell(session, "tail -f /dev/null")
		if err != nil {
			log.Println(err)
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
