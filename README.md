# spectrum-relay

## UPATE

Spectrum is no longer in active development.  If you have any questions or concerns please contact us at https://revsys.com

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

### Making a release

Making a release requires having correctly set your `$GOPATH` variable and directory structure according to https://golang.org/doc/code.html. It assumes the current directory is at `$GOPATH/src/github.com/SpectrumApp/spectrum-relay` or symlinked there.

It requires `goxc` and `bumpversion`

1. Run `bumpversion <version part>`
2. Run `make build`
