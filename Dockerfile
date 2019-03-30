FROM ubuntu:14.04
COPY bin/linux/amd64/process_collector /usr/sbin/process_collector

EXPOSE 9020

CMD /usr/sbin/process_collector -port=9020
