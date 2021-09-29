FROM golang:1.14 as build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -o hello-gitops cmd/main.go

FROM alpine:3.14.2 
EXPOSE 8080
WORKDIR /app
COPY --from=build /build/hello-gitops .
CMD ["./hello-gitops"]

#3.12 alpine image had following issues - CVE-2021-36159,CVE-2021-3711,CVE-2021-3711,