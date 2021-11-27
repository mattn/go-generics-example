# binary-size

## Binary sizes are same with `go build`

```
$ make
go tool compile -S main1.go > main1.s
go tool compile -S main2.go > main2.s

$ go build -o main1.exe main1.go

$ go build -o main2.exe main2.go

$ ls -la main1.exe main2.exe
-rwxr-xr-x 1 mattn mattn 1163264 Nov 28 00:43 main1.exe
-rwxr-xr-x 1 mattn mattn 1163264 Nov 28 00:43 main2.exe
```
