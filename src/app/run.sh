docker rm -f duplo && \
docker run \
  --name=duplo \
  --publish=5732:5732 \
  --volume=$(which docker):/usr/bin/docker \
  --volume=/var/run/docker.sock:/var/run/docker.sock \
  --volume=/lib64/libdevmapper.so.1.02:/usr/lib/libdevmapper.so.1.02 \
  --volume=/lib64/libudev.so.0:/usr/lib/libudev.so.0 \
  --volume=$(pwd)/apps:/etc/duplo \
  --detach \
  --privileged \
  docker.owm.io/marvell/duplo:0.0.6

./duplo_linux_amd64 -dir=./apps server >/home/marvell/duplo.log 2>&1 &
