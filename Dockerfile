FROM alpine:latest

COPY ./bin/app /bin/duplo
RUN mkdir -p /etc/duplo

VOLUME [ "/etc/duplo" ]

CMD [ "/bin/duplo" ]
