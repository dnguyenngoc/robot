# ######################################################################
# Build golang backend
# ######################################################################
FROM golang:alpine as builder

LABEL maintainer="Duy Nguyen <duynguyenngoc@hotmail.com"

WORKDIR /go/src/app

# Get Reflex for live reload in dev
ENV GO111MODULE=on

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

COPY go.mod .
COPY go.sum .

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./build .


# ######################################################################
# Deploy just keep file needed
# ######################################################################
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/


#Copy executable from builder
COPY --from=builder /go/src/app/build .

EXPOSE 8080
CMD ["./build"]
