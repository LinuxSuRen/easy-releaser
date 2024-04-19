ARG REGISTRY=172.11.0.6:30002/
FROM ${REGISTRY}docker.io/library/golang:1.20 as builder
WORKDIR /app
COPY . .
RUN go build -o bin/releaser .

FROM ${REGISTRY}al-cloud/alpine:3.18.4
COPY --from=builder /app/bin/releaser /usr/local/bin/releaser

ENTRYPOINT ["releaser"]
