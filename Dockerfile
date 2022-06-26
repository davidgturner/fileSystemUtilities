# use a go lang based image
FROM golang:1.15.15-alpine3.14 as builder
#FROM golang:1.15-buster as builder

# ARG GITHUB_TOKEN
ENV component_name = "fsu"
ENV CHROME_BIN="/usr/bin/chromium-browser"

# RUN apk update && apk add --no-cache ca-certificates git chromium npm nodejs sudo && rm -rf /var/cache/apk/*
RUN apk update && apk add --no-cache ca-certificates && rm -rf /var/cache/apk/*

# RUN set -x \
#     && apk update \
#     && apk upgrade \
#     && apk add --no-cache \
    # udev \
    # ttf-freefont \
    # chromium 

# RUN mkdir -p /opt/fs/bin
WORKDIR /opt/fs/bin

# RUN git config --global url."https://${GITHUB_TOKEN}}:@github.com/".insteadOf "https://github.com/"

COPY . .

# RUN npm ci

# RUN whoami

# RUN addgroup -S cassandra && adduser -S cassandra -G cassandra
# RUN chown -R cassandra:cassandra /opt/fs/bin
# RUN echo 'cassandra  ALL=(ALL) /bin/su' >>  /etc/sudoers
# USER cassandra

# RUN whoami

RUN go mod download

# RUN go run github.com/mxschmitt/playwright-go/cmd/playwright install --with-deps
RUN go run github.com/mxschmitt/playwright-go/cmd/playwright install --with-deps

# RUN CGO_ENABLED=0 go build -ldflags '-w -s' -o bin/${component_name} .
RUN go build -o bin/fsu .

FROM alpine:3.14

# FROM debian:buster-20210816-slim
#FROM zenika/alpine-chrome:with-playwright
#docker pull zenika/alpine-chrome:89-with-playwright
# FROM zenika/alpine-chrome:89-with-playwright

# ENV CHROME_BIN="/usr/bin/chromium-browser"
#     PUPPETEER_SKIP_CHROMIUM_DOWNLOAD="true"
# RUN set -x \
#     && apk update \
#     && apk upgrade \
#     && apk add --no-cache \
#     udev \
#     ttf-freefont \
#     chromium \
#     nodejs \
#     npm
    #\
    # && npm install puppeteer@1.10.0

# RUN npm i playwright

# RUN mkdir -p /opt/fs/bin
WORKDIR /opt/fs
COPY --from=builder /opt/fs/bin ./

CMD ["bin/fsu"]