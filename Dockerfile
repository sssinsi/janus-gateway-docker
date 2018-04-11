FROM ubuntu:16.04

WORKDIR /root

RUN apt-get update

RUN apt-get install -y git cmake wget python libmicrohttpd-dev libjansson-dev libnice-dev \
	libssl-dev libsrtp-dev libsofia-sip-ua-dev libglib2.0-dev \
	libopus-dev libogg-dev libcurl4-openssl-dev liblua5.3-dev \
	pkg-config gengetopt libtool automake doxygen graphviz libavutil-dev libavformat-dev

RUN git clone https://github.com/warmcat/libwebsockets.git && \
      cd libwebsockets && \
      mkdir build && \
      cd build && \
      cmake -DCMAKE_INSTALL_PREFIX:PATH=/usr -DCMAKE_C_FLAGS="-fpic" .. && \
      make && \
      make install

RUN cd /root
RUN rm -rf libwebsockets

RUN wget https://github.com/cisco/libsrtp/archive/v2.0.0.tar.gz && \
    tar xfv v2.0.0.tar.gz && \
    cd libsrtp-2.0.0 && \
    ./configure --prefix=/usr --enable-openssl && \
   make && \
   make shared_library && \
   make install

RUN cd /root
RUN rm -f v2.0.0.tar.gz
RUN rm -rf libsrtp-2.0.0

RUN git clone https://github.com/meetecho/janus-gateway.git && \
    cd janus-gateway && \
    sh autogen.sh && \
    ./configure --enable-post-processing --enable-docs --prefix=/usr/local && \
    make && \
    make install && \
    make configs
RUN sed -i -e "s/admin_http = no/admin_http = yes/g" /usr/local/etc/janus/janus.transport.http.cfg
RUN sed -i -e "s/admin_https = no/admin_https = yes/g" /usr/local/etc/janus/janus.transport.http.cfg
RUN cd /root
RUN rm -rf janus-gateway

CMD /usr/local/bin/janus