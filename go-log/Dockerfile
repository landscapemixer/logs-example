From golang:1.14 as bd
RUN adduser --disabled-login appuser
WORKDIR /go-log/
COPY . /go-log/
# In zh change to the GOPROXY="https://goproxy.io,direct"
RUN GOPROXY="https://goproxy.io,direct" GOSUMDB=off go build -ldflags="-w -s" -a -o .

FROM alpine
RUN apk --update add ca-certificates
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=bd /go-log /app/
COPY --from=bd /etc/passwd /etc/passwd
USER appuser
WORKDIR /app
CMD ./go-log