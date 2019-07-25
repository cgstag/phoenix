ARG GO_VERSION=1.12
# Step 1 : Build
FROM golang:${GO_VERSION}-alpine AS builder
# All these steps will be cached
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group
RUN apk add --no-cache ca-certificates git
WORKDIR /src
# COPY go.mod and go.sum files to the workspace
COPY go.mod .
COPY go.sum .
RUN go mod download
# COPY the source code as the last step
COPY . .

# Build the binary
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app .
# <- Second step to build minimal image
FROM scratch AS final
# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/
# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Import the compiled executable from the first stage.
COPY --from=builder /app /app
EXPOSE 8080
# Perform any further action as an unprivileged user.
USER nobody:nobody
ENTRYPOINT ["/app"]