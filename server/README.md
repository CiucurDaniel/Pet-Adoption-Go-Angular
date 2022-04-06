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

# Info

Processing HTTP requests in Go is primarily about two things: `handlers` and `servemuxes`.


### Handlers 

If you think of MVC, handlers are kinda like controllers. They are responsible for carrying out your application bussiness logic
and writing response headers and bodies.

### Servemuxes 

A servmux also known as a `router` stores a mapping between the predefined URL path and the corresponding handler.
Usually you have one servemux for your application containing all your routes.

For example Go's `net/http` package ships with a default `http.ServeMux`.
And when you do for example

```go
http.HandleFunc("/hello", hello)
```

What this function does is that it actually forwards that to the servemux. The internal implementation looks like this.

```go
// HandleFunc registers the handler function for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}
```

This is also why you start your http server like this:

```go
http.ListenAndServe(":8080", nil)
```

The `nil` denotes that we do not have a custom servemux, we simply want to use the default one from `net/http` package.

However, in our case for the application we do use a custom servemux, namely the `gorilla/mux`.
