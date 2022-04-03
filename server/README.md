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