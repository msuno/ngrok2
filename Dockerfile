FROM alpine:3.5

MAINTAINER The Ngrok Project <msuno@msuno.cn>

EXPOSE 80 443 4443 8000

ADD ngrokd /bin/ngrokd
RUN chmod 775 /bin/ngrokd

RUN mkdir /lib64
RUN ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

CMD ["/bin/ngrokd","-domain=ngrok.msuno.cn","-log-level=INFO","-redisAddr=47.102.201.87:6379","-redisPwd=msuno","-log=/log/ngrok.log"]
