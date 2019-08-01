.PHONY: all
all: build
FORCE: ;

.PHONY: build

build: build-auth build-feedback build-vote

build-auth:
	cd auth; go build -o ../bin/auth main.go

build-feedback:
	cd feedback; go build -o ../bin/feedback main.go

build-vote:
	cd vote; go build -o ../bin/vote main.go
