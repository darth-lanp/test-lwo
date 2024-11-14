ARG GO_IMAGE_VERSION
FROM --platform=linux/amd64 golang:$GO_IMAGE_VERSION

ENV GOFLAGS=-mod=vendor
ENV TRIVY_INSECURE=true
ENV SSL_CERT_DIR=/etc/ssl/certs

ARG cert_location=/usr/local/share/ca-certificates
ARG TRIVY_VERSION

WORKDIR /app

ADD https://github.com/aquasecurity/trivy/releases/download/v${TRIVY_VERSION}/trivy_${TRIVY_VERSION}_Linux-64bit.tar.gz trivy_${TRIVY_VERSION}_Linux-64bit.tar.gz

RUN mkdir bin && \
    mkdir /opt/reports && \
    tar -xzf trivy_${TRIVY_VERSION}_Linux-64bit.tar.gz -C bin && \
    apk update --no-check-certificate && \
    apk add --no-check-certificate ca-certificates openssl && \
    apk cache clean && \
    mkdir -p ${cert_location} && \
    openssl s_client -showcerts -connect github.com:443       </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt && \
    openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/proxy.golang.crt && \
    update-ca-certificates

COPY source .

RUN go mod download && \
    go build -o strgt ./cmd/trivy/

CMD ["sh", "/app/cmd/trivy/start.sh"]