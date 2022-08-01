FROM alpine:latest

ARG ARCH=${ARCH}
ADD release/${ARCH}/bond-bot /usr/local/bin
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.bfsu.edu.cn/g' /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Taipei /etc/localtime \
    && echo "Asia/Taipei" > /etc/timezone \
    && mkdir -p /etc/bond-bot

CMD ["bond-bot", "-c", "/etc/bond-bot/config.json"]
