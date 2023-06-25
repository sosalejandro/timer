FROM golang:1.20.5

RUN useradd -ms /bin/bash mockgen
USER mockgen
WORKDIR /home/mockgen/src

RUN go install github.com/golang/mock/mockgen@v1.6.0

ENTRYPOINT ["go", "generate", "-v", "./..."]
