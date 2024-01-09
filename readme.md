# Issue

Different behavior in the standard go compiler and in tinygo

## Go
```
$ go build -o test_go .
$ ./test_go
```

Output:
```
Hello, world
```

## Tinygo
```
$ tinygo build -o test_tinygo .
$ ./test_tinygo
```

Output:
```
panic: runtime error at 0x0000000000211788: type assert failed
Aborted (core dumped)
```

# Environment
```
$ go version
go version go1.21.1 linux/amd64
```

```
$ tinygo version
tinygo version 0.30.0 linux/amd64 (using go version go1.21.1 and LLVM version 16.0.1)
```
