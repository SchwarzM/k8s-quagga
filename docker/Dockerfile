FROM gliderlabs/alpine
MAINTAINER scorp@3dwars.net

ENV UPDATED_AT 2015-08-17

RUN apk-install quagga \
                supervisor

ADD supervisord.conf /etc/supervisord.conf
ADD zebra.conf.sample /etc/quagga/zebra.conf
ADD ospfd.conf.sample /etc/quagga/ospfd.conf

EXPOSE 179 2601 2604

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
