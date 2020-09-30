FROM centos:7
MAINTAINER mhanzaipeng 493096871@qq.com


RUN yum install -y gcc gcc-c++ glibc make autoconf openssl openssl-devel wget
RUN yum update -y

RUN yum install -y libxslt-devel -y gd gd-devel  pcre pcre-devel
RUN wget http://nginx.org/download/nginx-1.14.0.tar.gz && tar -zxvf  nginx-1.14.0.tar.gz \
    && cp -r nginx-1.14.0 /usr/local/
RUN useradd -M -s /sbin/nologin nginx
WORKDIR /usr/local/nginx-1.14.0
RUN ./configure --user=nginx --group=nginx --prefix=/usr/local/nginx --with-file-aio --with-http_ssl_module --with-http_realip_module --with-http_addition_module --with-http_xslt_module --with-http_image_filter_module --with-http_sub_module --with-http_dav_module --with-http_flv_module --with-http_mp4_module --with-http_gunzip_module --with-http_gzip_static_module --with-http_auth_request_module --with-http_random_index_module --with-http_secure_link_module --with-http_degradation_module --with-http_stub_status_module && make && make install

RUN ls /usr/local
EXPOSE 80

CMD /bin/sh -c '/usr/local/nginx/sbin/nginx  -g "daemon off;"'
