# http server listen unix domain socket

```console
# server
$ go run main.go
```

```console
# client
$ curl --unix-socket /tmp/app.sock 127.0.0.1
{"status":200"}
```

```console
# server
## Ctrl-C
2019/09/29 15:53:15 Caught signal interrupt: shutting down.
$ test -f /tmp/app.sock
$ echo $?
1
```
