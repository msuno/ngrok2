FROM alpine:3.5

LABEL maintainer The Ngrok Project <msuno@msuno.cn>

EXPOSE 80 443 4443 8000

ADD ngrokd /bin/ngrokd
RUN chmod 775 /bin/ngrokd

RUN mkdir /lib64
RUN ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

CMD ["/bin/ngrokd","-log=/log/ngrok.log", "-url=user:password@tcp(127.0.0.1:3306)/db?parseTime=true&loc=Local"]
