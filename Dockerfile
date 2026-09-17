FROM golang:1.21 AS builder
WORKDIR /workspace
COPY go.mod go.mod
# COPY go.sum go.sum
RUN go mod download
COPY cmd/ cmd/
COPY api/ api/
COPY pkg/ pkg/
RUN CGO_ENABLED=0 GOOS=linux go build -a -o manager cmd/manager/main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/manager .
USER 65532:65532
ENTRYPOINT ["/manager"]
