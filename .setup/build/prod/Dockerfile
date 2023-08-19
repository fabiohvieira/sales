FROM golang:1.18 as BUILD_IMAGE

WORKDIR /app
ENV TZ America/Sao_Paulo

COPY go.mod go.sum ./

RUN go mod download

COPY . /app

RUN CGO_ENABLED=1 GOOS=linux go build -o sales cmd/app/main.go

EXPOSE 3000

# FROM gcr.io/distroless/static

# COPY --from=BUILD_IMAGE /app/main /

CMD ["./sales"]
