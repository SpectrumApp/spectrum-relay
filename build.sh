#!/usr/bin/env bash
set -e

cd $GOPATH/src/github.com/SpectrumApp/spectrum-relay

goxc -c=.goxc.json -d=$GOPATH/src/github.com/SpectrumApp/spectrum-relay/dist

cd -