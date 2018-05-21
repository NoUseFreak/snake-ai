

default:
	go get || true
	go run *.go

run_docker:
	docker run -ti -v `pwd`:/go/src/app -w /go/src/app golang:latest make
