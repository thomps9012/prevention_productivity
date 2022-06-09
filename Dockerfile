FROM golang:1.18-alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/prevention_productivity /app/
WORKDIR /app
CMD ["./prevention_productivity"]