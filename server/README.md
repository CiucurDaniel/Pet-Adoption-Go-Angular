# Pet-adoption-server

This is the server-side of the pet adoption project.

# How to run

The project already provides a Dockerfile in order to easily set-up and rune the code.

Build the image:
```
docker build -t pet-adoption-server:latest .
```

Rune the image in a container:
```
docker run -it -p 8080:8080 pet-adoption-server
```

Go to:
```
http://localhost:8080/
```

# MySQL Database set-up

To get up as fast as possible use docker to spin-up a database

```go
docker run --name dev-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-password -d mysql:latest
```

You also have to go on the DB server and first create the `petadotiondb` schema.
This will be updated later.