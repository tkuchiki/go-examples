# file checksum

```console
$ cat file.txt ; echo
111111111122222222223333333333
$ cat file2.txt ; echo
111111111122222222224444444444

$ go run main.go
file.txt checksum c669d64fb5338cfb1dfc3a7fcfb028b8cccf0ff65f5cbfe11838fc1d9d90c44a
1111111111
[checksum] head 10 bytes d2d02ea74de2c9fab1d802db969c18d409a8663a9697977bb1c98ccdd9de4372
3333333333
[checksum] tail 10 bytes 36ecf9133cf3bb7963d53f30ff300c9f376c7ba12eefa341114563518bea8ec5

file2.txt checksum 20287dc82f5efa967a5afce36803927cc37ec55e1ce8b9f9a90e870de409d77a
1111111111
[checksum] head 10 bytes d2d02ea74de2c9fab1d802db969c18d409a8663a9697977bb1c98ccdd9de4372
4444444444
[checksum] tail 10 bytes f12a38838db97f7767c61d3922fa073656e407f00d8dc7337e5b5d0b009221da

file.txt head == file2.txt head
```
