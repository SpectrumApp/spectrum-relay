# spectrum-relay

Usage:

```
$ myprocess | relay [--level <LEVEL>]  [--sublevel <sublevel>]  [--endpoint <endpoint>]

```

options:

* `--level` Default: `INFO`
* `--sublevel` Default: `user`
* `--endpoint` Default: `http://127.0.0.1:9000`

### Quick Testing

1. Make a virtualenv and install the requirements.
2. run `go build relay.go && ./echo.py | ./relay`