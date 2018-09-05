FROM golang:latest

# for install go packages
RUN go get github.com/go-sql-driver/mysql

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/kudou-reira/myproject

RUN go get github.com/tockins/realize
RUN go get github.com/pilu/fresh
RUN go get -u github.com/gorilla/mux

# build executable
# don't build, i guess
# RUN go install github.com/kudou-reira/myproject


# execute 
# ENTRYPOINT /go/bin/myproject

WORKDIR /go/src/github.com/kudou-reira/myproject

CMD ["realize", "start", "--no-config", "--run"]


# Document that the service listens on port 8080.
EXPOSE 8080
