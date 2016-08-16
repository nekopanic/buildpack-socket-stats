Dyno socket stats!

Print out socket statistics every 30 seconds or so:

```

```

The statistics are absolute, printing exact amount each time it is checked. You can adjust the interval like this: `heroku config:set DYNO_SOCKET_STATS_INTERVAL=60`. 

# Hack it!

* `go get github.com/c9s/goprocinfo/linux`
* Edit [dyno-socket-stats.go](dyno-socket-stats.go)
* `go build -o dyno-socket-stats`
