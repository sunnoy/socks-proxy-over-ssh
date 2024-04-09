package main

import (
	_ "embed"
	"fmt"
	"github.com/blacknon/go-sshlib"
	"github.com/things-go/go-socks5"
	"golang.org/x/crypto/ssh"
	"log"
	"math/rand"
	"os"
	"sync"
)

var (

	//go:embed host
	sshHost string
	sshPort = "3390"

	//go:embed user
	sshUser string

	sshListen = "0.0.0.0:3391"

	// 将 sz 文件写入变量
	//go:embed ssh.key
	sshKey string

	wg = sync.WaitGroup{}

	socks5ListenChan = make(chan string)
)

func main() {

	wg.Add(2)

	go socks5Server()
	go sshRemoteForward()

	wg.Wait()

}

func sshRemoteForward() {

	// Create sshlib.Connect
	con := &sshlib.Connect{}

	err := os.WriteFile(".ssh.key", []byte(sshKey), 0644)
	if err != nil {
		log.Println(err)
	}
	keyAuthMethod, _ := sshlib.CreateAuthMethodPublicKey(".ssh.key", "")

	if keyAuthMethod != nil {
		err := con.CreateClient(sshHost, sshPort, sshUser, []ssh.AuthMethod{keyAuthMethod})
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	// PortForward
	socks5Listen := <-socks5ListenChan
	err = con.TCPRemoteForward(socks5Listen, sshListen)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("hello")
	}

	wg.Done()
}

func socks5Server() {
	// Create a SOCKS5 server
	server := socks5.NewServer(
	//socks5.WithLogger(socks5.NewLogger(log.New(os.Stdout, "socks5: ", log.LstdFlags))),
	)

	log.Println("world")

	port := rand.Intn(9000) + 10000
	socks5Listen := fmt.Sprintf("127.0.0.1:%d", port)

	// 将socks5Listen发送到channel
	socks5ListenChan <- socks5Listen

	if err := server.ListenAndServe("tcp", socks5Listen); err != nil {
		log.Println(err)
	}

	wg.Done()
}
