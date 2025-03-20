FROM golang:1.24.1

# set working directory
WORKDIR /go/src/app

# copy the source code
COPY . .

# expose the port
EXPOSE 8000

# build the source code
RUN go build -o main cmd/api/main.go

# run the application
CMD ["./main"]