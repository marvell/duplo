#!/bin/sh

docker run \
  --name=duplo \
  --detach \
  --restart=always \
  --volume=/usr/bin/docker:/usr/bin/docker \
  --volume=/var/run/docker.sock:/var/run/docker.sock \
  --publish=5732:5732 \
  marvell/duplo:latest
