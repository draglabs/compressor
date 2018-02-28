FROM golang:1.9.0




RUN mkdir -p /go/src/compressor
WORKDIR /go/src/compressor

COPY . /go/src/compressor

RUN go-wrapper download
RUN go-wrapper install


expose 8081

CMD ["go-wrapper", "run"]
