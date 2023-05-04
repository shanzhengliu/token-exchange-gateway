FROM golang:alpine as builder

RUN mkdir /application/
WORKDIR /application/
COPY . .
RUN go build -o main ./main.go

FROM scratch
COPY --from=builder /application/main /application/main
EXPOSE 8080
ENTRYPOINT [ "/application/main" ]