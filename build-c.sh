brew install unzip
unzip wasi-sdk-7.0.zip
rm -rf /usr/local/opt/wasi-sdk/
cp -r ./wasi-sdk-7.0/opt/ /usr/local/opt/
rm -rf ./wasi-sdk-7.0
brew install llvm@9
if [ ! -d "/usr/local/Cellar/llvm/9.0.0_1/lib/" ]
then
    echo "Directory /usr/local/Cellar/llvm/9.0.0_1/lib/ DOES NOT exists."
else
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
        cd vertex-cdt && ./build.sh
        if [ ! -f "/usr/local/bin/vertex-cdt" ]
        then
            echo "build vertex-cdt fail, please check go version"
        else
            echo "Build successful !"
        fi
    fi
fi