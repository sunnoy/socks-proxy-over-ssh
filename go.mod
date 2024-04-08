module socks-proxy-over-ssh

go 1.19

require (
	github.com/blacknon/go-sshlib v0.1.10
	github.com/things-go/go-socks5 v0.0.5
	golang.org/x/crypto v0.18.0
)

require (
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/ScaleFT/sshkeys v1.2.0 // indirect
	github.com/ThalesIgnite/crypto11 v1.2.5 // indirect
	github.com/armon/go-socks5 v0.0.0-20160902184237-e75332964ef5 // indirect
	github.com/dchest/bcrypt_pbkdf v0.0.0-20150205184540-83f37f9c154a // indirect
	github.com/lunixbochs/vtclean v1.0.0 // indirect
	github.com/miekg/pkcs11 v1.1.1 // indirect
	github.com/moby/term v0.0.0-20221128092401-c43b287e0e0f // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/thales-e-security/pool v0.0.2 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/term v0.16.0 // indirect
)

replace github.com/ThalesIgnite/crypto11 v1.2.5 => github.com/blacknon/crypto11 v1.2.6
