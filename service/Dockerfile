# ######################################################################
# Build golang backend
# ######################################################################
FROM golang:alpine as builder

# Add Maintainer Info
LABEL maintainer="Duy Nguyen <duynguyenngoc@hotmail.com"

RUN mkdir /app

# Set the Current Working Directory inside the container
WORKDIR /app

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# # Copy the source from the current directory to the Working Directory inside the container
COPY go.mod .
COPY go.sum .

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY ./app/ /app/

# Build the Go app
RUN go build -o main


# # ######################################################################
# # Deploy just keep file needed
# # ######################################################################
# FROM alpine:latest

# RUN apk --no-cache add ca-certificates

# WORKDIR $GOPATH/src/

# # Keep main and template
# COPY --from=builder $GOPATH/src/main $GOPATH/src/
# COPY --from=builder $GOPATH/src/template/ $GOPATH/src/template/

# EXPOSE 8000


go mod init duynguyenngoc/robot

go get -u github.com/gin-gonic/gin