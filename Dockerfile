FROM golang:onbuild
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD [ "/app/main" ]
EXPOSE 80 8080 300
