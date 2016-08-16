Dyno socket stats!

Print out socket statistics every 30 seconds or so:

```
2016-08-16T10:18:20.706249+00:00 app[web.1]: ps=web.1 established=2 listen=1
```

The statistics are absolute, printing exact amount each time it is checked. You can adjust the interval like this: `heroku config:set DYNO_SOCKET_STATS_INTERVAL=60`. 

# Hack it!

* `go get github.com/c9s/goprocinfo/linux`
* Edit [dyno-socket-stats.go](dyno-socket-stats.go)
* `go build -o dyno-socket-stats`
