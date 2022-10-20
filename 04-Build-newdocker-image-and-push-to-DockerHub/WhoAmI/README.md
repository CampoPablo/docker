[![Docker Image](http://img.shields.io/badge/Docker-Image-blue.svg)](https://registry.hub.docker.com/u/flowerpot/goif/)

WHOAMI
======

Simple go webserver, that prints hostname and all network interface ip
adresses.

I use it as a dummy application in cluster deployments, for example when
testing out configuration management or orchestration, so I can see that load
balancing is working. The advantage of this is no depencencies and only one
binary/docker-image to deploy.

Try with

	docker run --rm -p 8080:8080 flowerpot/goif

If you also want to try sending the logs to a syslog server

	docker run --rm -p 8080:8080 --log-driver=syslog --log-opt syslog-address=udp://syslog-server:514 flowerpot/goif

Even more comfortable using docker-compose

	docker-compose up

> Hint: make sure you uncomment and change the `syslog-server` hostname
> placeholder to your desired remote syslog server as well as the port
