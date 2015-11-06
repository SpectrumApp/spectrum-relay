# spectrum-relay

Usage:

```
$ myprocess | relay [--level <LEVEL>]  [--sublevel <sublevel>]  [--endpoint <endpoint>] [--version]

```

options:

* `--level` Default: `INFO`
* `--sublevel` Default: `user`
* `--endpoint` Default: `http://127.0.0.1:9000`
* `--version` Prints out the current version and exits

### Quick Testing

1. Make a virtualenv and install the requirements.
2. run `go build relay.go && ./echo.py | ./relay`