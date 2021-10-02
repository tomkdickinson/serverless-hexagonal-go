# Severless Hexagonal Go Template

This is an opinionated [Serverless Framework](http://www.serverless.com) template that can be used to bootstrap a Go 
project to use a hexagonal architecture.

It's designed primarily for building microservices around lambdas. 

## Features

* Implements a 
[hexagonal layout](https://medium.com/@tomkdickinson/hexagonal-architecture-with-go-and-google-wire-e4344dd24b94), 
abstracting lambdas/storage from domain logic
* [Zerolog](https://github.com/rs/zerolog) included for contextual JSON logging
* [Wire](https://github.com/google/wire) included to manage dependency injection

## To Use

You can setup a new serverless project with this template with:

```shell
sls create --path {SERVICE_NAME} --template-url https://github.com/tomkdickinson/serverless-hexagonal-go/tree/main
```

Then run `make gomodgen MODULE={MODULE_NAME}` (or `make gomodgen` to auto generate module based on parent and folder name) to 
generate the go.mod and go.sum files. 

This template also uses [Wire](https://github.com/google/wire) for the examples, although you can easily be taken out if
not needed.

## Example Folder Structure

The example provided with this template is a simple Blog endpoint, where we want to find either a list of posts, or a
specific post by its slug. 

The `cmd` folder contains main applications for each lambda we declare in [serverless.yml](serverless.yml), including
wiring of layers using [Wire](https://github.com/google/wire).

[internal/lambda](internal/lambda) contains the actual lambda implementations for each http request.

[internal/blog](internal/blog) contains the domain logic for interacting with the blog.

[internal/storage](internal/storage) contains a simple `memory` implementation of the blogs repository, to load posts.

