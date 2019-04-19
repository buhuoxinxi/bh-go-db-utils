# bh-go-db-utils

捕获科技 db 工具

## config

see ./config/config.go

## dev environment

go version

> go version go1.12 darwin/amd64 & go module enable

mysql version

> 5.7

## test 

run test

```bash

go test -v .
# === RUN   TestNewDBConn
# --- PASS: TestNewDBConn (0.04s)
#     database_test.go:20: test mysql ... 
#     database_test.go:31: &{<nil> <nil> 0 0xc0000e03c0 false 2 {0xc00011be00} <nil> {{0 0} {<nil>} map[] 0} 0xc0000aa370 0x16a99c0 0xc00017a500 false} 
#     database_test.go:34: test postgres ... 
#     database_test.go:47: &{<nil> <nil> 0 0xc0000e0540 false 2 {0xc00011be00} <nil> {{0 0} {<nil>} map[] 0} 0xc0000aa630 0x16a99c0 0xc00017a660 false} 
# PASS
# ok      github.com/buhuoxinxi/bh-go-db-utils    (cached)

```