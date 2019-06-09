VERSION := $(shell git tag | grep ^v | sort -V | tail -n 1)
LDFLAGS = -ldflags "-X github.com/AgentZombie/go-embed-version/cmd.Version=${VERSION}"
TIMESTAMP := $(shell date +%Y%m%d-%H%M)
DEVLDFLAGS = -ldflags "-X github.com/AgentZombie/go-embed-version/cmd.Version=dev-${TIMESTAMP}"

SERVBIN=/tmp/versserv

dev-serv: server/*
	go build ${DEVLDFLAGS} -o ${SERVBIN} server/*.go

serv: server/*
	go build ${LDFLAGS} -o ${SERVBIN} server/*.go
