# after building, run with docker 

FROM ubuntu:groovy
MAINTAINER Quan Nguyen
ENV GO12MODULE=on

RUN apt-get -y update && apt-get -y upgrade

#Utilities install llvm and clang
RUN apt-get update -y \
    && apt-get install -y llvm-9 libllvm-9-ocaml-dev libllvm9 llvm-9-dev llvm-9-doc llvm-9-examples llvm-9-runtime \
    clang-9 clang-tools-9 clang-9-doc libclang-common-9-dev libclang-9-dev libclang1-9 clang-format-9 python-clang-9 clangd-9 \
    libfuzzer-9-dev \
    lldb-9 \
    lld-9 \
    libc++-9-dev libc++abi-9-dev \
    libomp-9-dev \
    cmake \
    build-essential \
    libncurses5 \
    git \ 
    unzip \
    curl 

RUN mkdir /usr/liquid-chain-sdk
COPY .  /usr/liquid-chain-sdk/
WORKDIR /usr/liquid-chain-sdk
RUN unzip wasi-sdk-ubuntu-9.0.zip
RUN cp -r ./wasi-sdk-ubuntu-9.0/opt/ /usr/local/opt/ && rm -rf ./wasi-sdk-ubuntu-9.0
# install rust 
# RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh && source $HOME/.cargo/env 
# RUN  rustup -y target add wasm32-wasi && rustup -y update nightly
# install c2ffi 
RUN cd c2ffi \
 && mkdir build \
 && cd build \
 && cmake .. \
 && make
# install go lang 
RUN apt-get install -y golang
RUN mv c2ffi/build/bin/c2ffi /usr/local/bin/
RUN cd liquid-cdt && ./build.sh
RUN cp -a  /usr/local/opt/wasi-sdk/share/wasi-sysroot/include/. /usr/lib/clang/9.0.1/include/
