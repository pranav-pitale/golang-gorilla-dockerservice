FROM golang

# Create the directory where the application will reside
RUN mkdir /go/src/velociraptorgo


# ADD the application files (needed for production)
ADD . /go/src/velociraptorgo

# Specify working directory
WORKDIR /go/src/velociraptorgo

# Import dependencies needed
RUN go get github.com/gorilla/handlers
RUN go get github.com/gorilla/mux

CMD ["go", "run", "main.go"]

EXPOSE 3000