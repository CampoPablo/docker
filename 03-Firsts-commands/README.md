# First commands

## Step-1: Verify Docker version and also login to Docker Hub
```
docker version
docker login
```

# Containers & images
## Step-2: List running containers
```
docker ps
docker ps -a 
docker container ls -a
docker ps -a -q
```


## Stop-3: List images
```
docker images ls
```

## Stop-4: Download image ubuntu
```
docker pull ubuntu
```

## Stop-5: List images. Do you see ubuntu image?
```
docker images ls
```

## Stop-6: Run container ubuntu
```
docker run ubuntu
```

## Stop-7: List container. Can you see Ubuntu cointainer?
```
docker ps -a
```

## Step-8: Connect to Container Terminal
```
docker exec -it <container-name> /bin/sh
```

```
docker start -a -i [container id]

a: standard output
i: interative mode.


write "exit" to quit
```

## step-9: drop container
```
docker rm [image id]
```

## step-9: drop image ubuntu
```
docker image rm ubuntu
```