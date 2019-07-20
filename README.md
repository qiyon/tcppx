# tcppx

A sample pure TCP proxy, Golang, with go.mod file.

## build and run

```
go build
./tcppx -l 0.0.0.0:18888 -r 127.0.0.1:18889
```

## pm2

pm2.json

```
{
  "apps" : [{
    "name"        : "tcppx",
    "script"      : "./tcppx",
    "args"        : ["-l", "0.0.0.0:18888", "-r", "127.0.0.1:18889"]
  }]
}
```

Run: `pm2 start pm2.json`

Runing Info: `pm2 show tcppx`

## reference

https://github.com/jpillora/go-tcp-proxy
