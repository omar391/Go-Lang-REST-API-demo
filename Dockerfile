FROM alpine:latest
WORKDIR /build
COPY . .
RUN apk add --no-cache  ca-certificates go
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go env -w GO111MODULE=on
RUN go build -v -o /clinics-apis .



#doing multi-stage build
FROM alpine:latest
WORKDIR /
ENV GIN_MODE=release
COPY --from=0 /clinics-apis .
COPY /conf.d ./conf.d
COPY /data ./data
ENTRYPOINT ["./clinics-apis"]