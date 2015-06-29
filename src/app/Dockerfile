FROM busybox

RUN mkdir -p /etc/duplo
COPY bin/duplo_linux_amd64 /usr/bin/duplo

COPY .docker/.dockercfg /root/.dockercfg

VOLUME ["/etc/duplo"]
EXPOSE 5732

CMD ["/usr/bin/duplo", "server"]
