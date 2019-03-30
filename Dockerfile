FROM ubuntu:14.04
COPY bin/linux/amd64/process_exporter /usr/sbin/process_exporter

EXPOSE 9020

CMD /usr/sbin/process_exporter -port=9020
