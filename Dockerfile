FROM golang:alpine3.18 AS builder

WORKDIR /app

COPY . .
RUN go build
RUN apk update && apk add curl
RUN curl https://metlo-releases.s3.us-west-2.amazonaws.com/metlo_agent_linux_amd64_latest > metlo-agent
RUN chmod +x metlo-agent

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/test-gin .
COPY --from=builder /app/metlo-agent /usr/local/bin

EXPOSE 3031

CMD ["/app/test-gin"]
