# Readme

A Websocket Test with the following goals

* Learn about Using Web Sockets in golang
* Learn about Using the `htmx` technology as a replacement for building responsive web applications
* Learn about using the `templ` library to build templates (think java server pages for go)
## Dependencies

Templ must be in your path. See

https://github.com/a-h/templ

For building and adding to your path.

## Getting Started

After CDing into the directory.

1. generate the `templ` templates with

    `$ templ generate`


2. Install go modules

    `$ go install`

3. Startup

    `$ go run main.go`


# Endpoints

The application starts up on port `:8080` and can be accessed from

http://localhost:8080

* '/' The default experience. Uses Htmx and hyperscript to build a simple chat bot that echos your inputted text

* '/chat', Similar to '/', but uses pure javascript functionality.

* '/hello', an endpoint to test `templ` templates


