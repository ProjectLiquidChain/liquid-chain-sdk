# build vetex-sdk for smart contract
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

if [ ! -d "/usr/local/opt/wasi-sdk/" ]
then
    rm -rf /usr/local/opt/wasi-sdk/
fi

cp -r ./wasi-sdk-7.0/opt/ /usr/local/opt/
rm -rf ./wasi-sdk-7.0
# install llvm, require llvm version 9
if [ ! -d "/usr/local/opt/llvm/" ]
then
    echo "llvm DOES NOT exists, install llvm: "
    brew install llvm@9
fi

if [ ! -d "/usr/local/Cellar/llvm/9.0.0_1/lib/" ]
then
    echo "Directory /usr/local/Cellar/llvm/9.0.0_1/lib/ DOES NOT exists, please update llvm version"
else
    # install c2ffi tool to support generate ABI from c, c++ langiage
    cd c2ffi/
    mkdir build
    cd build/
    LLVM_DIR=/usr/local/opt/llvm/ cmake ..
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
            rustup target add wasm32-wasi
        fi
        # install vertex-cdt
        cd vertex-cdt && ./build.sh
        if [ ! -f "/usr/local/bin/vertex-cdt" ]
        then
            echo "build vertex-cdt fail, please check go version"
        else
            echo "Build successful ! "
        fi
    fi
fi