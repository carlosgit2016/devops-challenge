FROM golang:alpine3.19

# Implementation goes here

# Using go build ./... will generate a main artifact that can be executed
CMD [ "/main" ]
