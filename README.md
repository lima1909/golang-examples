# GoLang
Examples for GoLang exrcises

## Examples:
* [consumer (http-client) use json](./json-consumer)
* [rest-server (can use from json-consumer, with -url command line parameter)](./rest-server)
* [command line interface (cli) with cobra](./cli)
* [ping-pong to illustrate concurrent acces (with channel)](./concurrent)
* [templates](./template)

## How can I test it?

### start the server

    # login bei [play-with-docker](https://labs.play-with-docker.com/)
    # clone the source from github
    $ git clone https://github.com/lima1909/golang-examples
    
    # start a docker container
    $ docker run -it -v /root/golang-examples:/go/src/golang-examples -p 8080:8080 golang:alpine3.7 sh

    # start the server (IN THE DOCKER CONTAINER)
    $ cd src/golang-examples
    $ go run rest-server/rest-server.go

### start the client/consumer

    # add new instance
    # clone the source from github
    $ git clone https://github.com/lima1909/golang-examples
    
    # start a docker container
    $ docker run -it -v /root/golang-examples:/go/src/golang-examples golang:alpine3.7 sh

    # start the json-consumer (IN THE DOCKER CONTAINER)
    $ cd src/golang-examples
    $ go run json-consumer/json-consumer.go -url http://[ip-from-server-container]:8080/user/1