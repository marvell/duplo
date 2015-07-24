#!/bin/sh

set -x

GOOS=linux gb build
docker build -t marvell/duplo:latest .
