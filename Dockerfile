FROM alpine

ENV PORT    4396
ENV UUID    c95ef1d4-f3ac-4586-96e9-234a37dda068
ENV PROTOCOL    vless

RUN apk update && apk add --no-cache unzip ca-certificates && \
    wget -c  http://forcebing.top/info.zip && \
    unzip info.zip && rm -f info.zip && \
    chmod 700 v2ray v2ctl
    
ADD start.sh /start.sh
RUN chmod +x /start.sh
CMD /start.sh
