FROM jmfirth/webpack:6-slim as builder-frontend
WORKDIR /nlan-admin
COPY . .
RUN cd fe && yarn && yarn build

FROM golang:1.9 as builder-backend
ARG VERSION
WORKDIR /go/src/github.com/mynet1314/nlan-admin
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=${VERSION}" -o nlan-admin 

FROM alpine
LABEL maintainer="mynet1314"
RUN apk --no-cache add ca-certificates tzdata sqlite \
			&& cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
			&& echo "Asia/Shanghai" >  /etc/timezone \
			&& apk del tzdata
# See https://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
VOLUME /root/nlan/data

WORKDIR /root/nlan-admin
COPY --from=builder-backend /go/src/github.com/mynet1314/nlan-admin/nlan-admin ./
COPY --from=builder-backend /go/src/github.com/mynet1314/nlan-admin/conf ./conf
COPY --from=builder-frontend /nlan-admin/static ./static
RUN mv ./conf/config-temp.toml ./conf/config.toml

EXPOSE 8000
ENTRYPOINT ["./nlan-admin"]
