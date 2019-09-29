# gzipped http response

```console
# server
$ go run main.go
```

```console
# client
$ curl -v 127.0.0.1:8080
* Rebuilt URL to: 127.0.0.1:8080/
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET / HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Sun, 29 Sep 2019 08:04:11 GMT
< Content-Length: 14
<
* Connection #0 to host 127.0.0.1 left intact
{"status":200}

$ curl -v -H "Accept-Encoding: gzip" 127.0.0.1:8080
* Rebuilt URL to: 127.0.0.1:8080/
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET / HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.54.0
> Accept: */*
> Accept-Encoding: gzip
>
< HTTP/1.1 200 OK
< Content-Encoding: gzip
< Content-Length: 38
< Content-Type: application/json
< Date: Sun, 29 Sep 2019 08:05:23 GMT
<
* Connection #0 to host 127.0.0.1 left intact
V*.I,)-V220?>

$ curl -v -H "Accept-Encoding: gzip" 127.0.0.1:8080 | gzip -d
* Rebuilt URL to: 127.0.0.1:8080/
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET / HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.54.0
> Accept: */*
> Accept-Encoding: gzip
>
< HTTP/1.1 200 OK
< Content-Encoding: gzip
< Content-Length: 38
< Content-Type: application/json
< Date: Sun, 29 Sep 2019 08:06:20 GMT
<
{ [38 bytes data]
100    38  100    38    0     0   7077      0 --:--:-- --:--:-- --:--:--  7600
* Connection #0 to host 127.0.0.1 left intact
{"status":200}
```
