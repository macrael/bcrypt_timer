# Bcrypt Timer

run it locally with:
```
$ dep ensure
$ go run main.go
```

run it in docker with
```
$ docker build .
$ docker run <container_id>
```

push it to docker with
```
docker login
docker tag <container_id> macrael/bcrypt-timer:version
docker push macrael/bcrypt-timer:version
```
