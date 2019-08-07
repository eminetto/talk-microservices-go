.PHONY: all
all: build
FORCE: ;

.PHONY: build

build: build-auth build-feedback build-vote

build-auth:
	cd auth; go build -o ../bin/auth main.go

build-feedback:
	cd feedbacks; go build -o ../bin/feedbacks main.go

build-vote:
	cd votes; go build -o ../bin/votes main.go
