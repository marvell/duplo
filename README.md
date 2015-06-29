# Duplo [![Circle CI](https://circleci.com/gh/marvell/duplo/tree/master.svg?style=svg)](https://circleci.com/gh/marvell/duplo/tree/master)


## Build

```
go get github.com/constabulary/gb/...
gb vendor update -all
gb build all
```

## Usage

```
./bin/app -dir=./tests -bind=:5732 server
```
