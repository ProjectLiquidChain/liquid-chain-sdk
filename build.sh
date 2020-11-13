# build vetex-sdk for smart contract
build_mac () {
    if [ ! -f "/usr/local/bin/cmake" ]
    then
        echo "cmake not exits"
        brew install cmake
    fi
    if [ ! -d "/usr/local/Cellar/unzip" ]
    then
        echo "unzip not exits"
        brew install unzip
    fi
    # unzip and install wasi-sdk
    unzip wasi-sdk-7.0.zip

    if [ -d "/usr/local/opt/wasi-sdk/" ]
    then
        echo "wasi-sdk exits"
        rm -rf /usr/local/opt/wasi-sdk/
    fi

    cp -r ./wasi-sdk-7.0/opt/ /usr/local/opt/
    rm -rf ./wasi-sdk-7.0
    # install llvm, require llvm version 9
    if [ ! -d "/usr/local/opt/llvm@9/" ]
    then
        echo "llvm DOES NOT exists, install llvm: "
        brew install llvm@9
    fi
    # install c2ffi tool to support generate ABI from c, c++ langiage
    cd c2ffi/
    mkdir build
    cd build/
    LLVM_DIR=/usr/local/opt/llvm@9/ cmake ..
    make
    mv ./bin/c2ffi /usr/local/bin
    if [ ! -f "/usr/local/bin/c2ffi" ]
    then
        echo "c2ffi not exists, please check config in c2ffi"
        exit
    else
        echo "c2ffi exists."
        cd ../..
        #install rust toolchain, choose option 1 to default install rust toolchain
        if [ ! -x "$(command -v cargo)" ]
        then
            echo "rust toolchain not exit"
            curl https://sh.rustup.rs -sSf | sh
            source $HOME/.cargo/env
            rustup target add wasm32-wasi
            rustup update nightly
        fi
        # install liquid-cdt
        cd liquid-cdt && ./build.sh
        if [ ! -f "/usr/local/bin/liquid-cdt" ]
        then
            echo "build liquid-cdt fail, please check go version"
        else
            echo "Build successful ! "
        fi
    fi
}
build_linux() {
    ## need support
    if [ ! -f "/usr/local/bin/cmake" ]
    then
        echo "cmake not exits"
        apt-get install cmake
    fi
    if [ ! -d "/usr/local/Cellar/unzip" ]
    then
        echo "unzip not exits"
        apt-get install unzip
    fi
    # unzip and install wasi-sdk
    unzip wasi-sdk-7.0.zip
    echo ""
    # install llvm, require llvm version 9
    if [ ! -d "/usr/local/opt/llvm/" ]
    then
        echo "llvm DOES NOT exists, install llvm: "
        apt-get install -y llvm-9 libllvm-9-ocaml-dev libllvm9 llvm-9-dev llvm-9-doc llvm-9-examples llvm-9-runtime 
        apt-get install clang-9 clang-tools-9 clang-9-doc libclang-common-9-dev libclang-9-dev libclang1-9 clang-format-9 python-clang-9 clangd-9
        apt-get install libfuzzer-9-dev
        apt-get install lldb-9
        apt-get install lld-9
        apt-get install libc++-9-dev libc++abi-9-dev
        apt-get install libomp-9-dev
    fi
    # install c2ffi tool to support generate ABI from c, c++ langiage
    cd c2ffi/
    mkdir build
    cd build/
    cmake ..
    make
    mv ./bin/c2ffi /usr/local/bin
    if [ ! -f "/usr/local/bin/c2ffi" ]
    then
        echo "c2ffi not exists, please check config in c2ffi"
        exit
    else
        echo "c2ffi exists."
        cd ../..
        #install rust toolchain, choose option 1 to default install rust toolchain
        if [ ! -x "$(command -v cargo)" ]
        then
            echo "rust toolchain not exit"
            curl https://sh.rustup.rs -sSf | sh
            source $HOME/.cargo/env
            rustup target add wasm32-wasi
            rustup update nightly
        fi
        # install liquid-cdt
        cd liquid-cdt && ./build.sh
        if [ ! -f "/usr/local/bin/liquid-cdt" ]
        then
            echo "build liquid-cdt fail, please check go version"
        else
            echo "Build successful ! "
        fi
    fi
}
if [[ "$OSTYPE" == "darwin"* ]]; then
    # Mac OSX
    build_mac
elif [[ "$OSTYPE" == "linux"* ]]; then
    # Ubuntu OS,...
    build_linux
    echo "build in linux"
else
    echo "OS not support !"
fi