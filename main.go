package main

import "os"

var (
	sshHost     = ""
	sshPort     = "22"
	sshUser     = "root"
	sshPassword = ""

	socks5Listen = "localhost:1989"

	sshListen = "0.0.0.0:3000"
)

func main() {
	getEnv()

}

func getEnv() {
	sshHost = os.Getenv("SSH_HOST")
	sshPort = os.Getenv("SSH_PORT")
	sshPassword = os.Getenv("SSH_PW")
	socks5Listen = os.Getenv("SOCKS_L")
	sshListen = os.Getenv("SSH_L")
}
