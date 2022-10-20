compile:
	go build -a -installsuffix cgo -o main .
build-docker:
	docker build -t pcsys/whoami .
release-docker: build-docker
	docker push pcsys/whoami
