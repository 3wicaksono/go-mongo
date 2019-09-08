# Stage build
FROM golang:1.11-alpine as builder

RUN apk update && apk add --no-cache \
                  bash \
                  curl \
                  openssh \
                  make \
                  git \
                  g++ \
                  pkgconf

ENV SERVICE_NAME=go-mongo

RUN wget https://github.com/Masterminds/glide/releases/download/v0.13.3/glide-v0.13.3-linux-amd64.tar.gz && \
    tar xvfz glide-v0.13.3-linux-amd64.tar.gz -C /usr/local/bin --strip-components=1 linux-amd64/glide && \
    rm glide-v0.13.3-linux-amd64.tar.gz

ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# compile go project
WORKDIR /go/src/${SERVICE_NAME}
COPY . .
RUN glide install
RUN CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -a -o go-mongo -tags static_all .

# Stage Runtime Applications
FROM alpine:latest

ENV BUILD_PATH=/go/src/go-mongo
ENV RUNTIME_PATH=/opt/go-mongo
WORKDIR ${RUNTIME_PATH}

# Download Depedencies
RUN apk update && apk add --no-cache ca-certificates bash jq curl


# Setting timezone
ENV TZ=Asia/Jakarta
RUN apk add -U tzdata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Add user wicak
RUN adduser -D wicak wicak

# Setting folder workdir
WORKDIR ${RUNTIME_PATH}
RUN mkdir configurations

# Copy Data App
COPY --from=builder ${BUILD_PATH}/go-mongo go-mongo
COPY --from=builder ${BUILD_PATH}/configurations/App.yaml configurations/App.yaml

# Setting owner file and dir
RUN chown -R wicak:wicak .

USER wicak

EXPOSE 8787
