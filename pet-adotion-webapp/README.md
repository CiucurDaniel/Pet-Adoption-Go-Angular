# PetAdotionWebapp

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 12.2.10.

# How to run

The project already provides a Dockerfile in order to easily set-up and rune the code.

Build the image: 

```
docker build -t pet-adoption-webapp:latest .
```

Run the image in a container:

```
docker run --name pet-adoption-webapp -it -p 8079:80 pet-adoption-webapp
```

Go to:

```
http://localhost:8079/
```