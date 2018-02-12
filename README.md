# artillery-iptables-cleanup

Short cleanup program, that could be put in *cron* to periodically clean up IPTABLES *DROP* rules 
from Artillery (https://github.com/BinaryDefense/artillery).

## How to compile

If you have your environment setup (https://golang.org/doc/code.html), do:

```
go build artillery-iptables-cleanup.go
```

## How to compile for different architecture (eg. amd64)

```
env GOOS=linux GOARCH=amd64 go build artillery-iptables-cleanup.go -o artillery-iptables-cleanup_amd64
```

## How to run

```
# ./artillery-iptables-cleanup
1    DROP       all  --  1.2.3.4         0.0.0.0/0
2    DROP       all  --  1.2.3.4         0.0.0.0/0
3    DROP       all  --  123.123.123.123         0.0.0.0/0
4    DROP       all  --  123.123.123.123         0.0.0.0/0
```
