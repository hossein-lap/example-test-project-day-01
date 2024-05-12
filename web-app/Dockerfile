## builder
FROM golang:alpine as builder
LABEL maintainer="Hossein Esmailzadeh <hosteam01@gmail.com>"

WORKDIR /app

# COPY go.mod go.sum ./
COPY go.mod ./
# RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

## runner
FROM alpine:latest  

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 9090

CMD ["./main"] 
