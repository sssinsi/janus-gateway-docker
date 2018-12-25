FROM ubuntu:16.04

WORKDIR /root

RUN apt-get update

RUN apt-get install -y sudo git cmake wget python libmicrohttpd-dev libjansson-dev libnice-dev \
	libssl-dev libsrtp-dev libsofia-sip-ua-dev libglib2.0-dev \
	libopus-dev libogg-dev libcurl4-openssl-dev liblua5.3-dev \
	libconfig-dev pkg-config gengetopt libtool automake doxygen graphviz libavutil-dev libavformat-dev

RUN git clone https://github.com/warmcat/libwebsockets && \
    cd libwebsockets && \
    mkdir build && \
    cd build && \
    cmake -DLWS_MAX_SMP=1 -DCMAKE_INSTALL_PREFIX:PATH=/usr -DCMAKE_C_FLAGS="-fpic" .. && \
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


RUN git clone https://github.com/sctplab/usrsctp && \
	cd usrsctp && \
	./bootstrap && \
	./configure --prefix=/usr && make && make install 
RUN rm -rf usrsctp 

RUN git clone https://github.com/meetecho/janus-gateway.git && \
    cd janus-gateway && \
    sh autogen.sh && \
    ./configure --enable-post-processing --prefix=/opt/janus --libdir=/usr/lib64 && \
    make && \
    make install && \
    make configs
# RUN sed -i -e "s/admin_http = no/admin_http = yes/g" /usr/local/etc/janus/janus.transport.http.cfg
# RUN sed -i -e "s/admin_https = no/admin_https = yes/g" /usr/local/etc/janus/janus.transport.http.cfg
RUN cd /root
RUN rm -rf janus-gateway

# install Rust
RUN wget -O rustup.sh https://sh.rustup.rs && \
    sudo sh ./rustup.sh -y

RUN git clone https://github.com/mozilla/janus-plugin-sfu.git && \
    cd janus-plugin-sfu && \
    $HOME/.cargo/bin/cargo build --release && \
    cp -ip ./target/release/libjanus_plugin_sfu.so /opt/janus/plugins

CMD /opt/janus/bin/janus -P=/opt/janus/plugins
