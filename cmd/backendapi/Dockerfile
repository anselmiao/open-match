FROM open-match-base-build as builder

WORKDIR /go/src/github.com/GoogleCloudPlatform/open-match/cmd/backendapi/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM gcr.io/distroless/static
COPY --from=builder /go/src/github.com/GoogleCloudPlatform/open-match/cmd/backendapi/backendapi .

ENTRYPOINT ["/backendapi"]
