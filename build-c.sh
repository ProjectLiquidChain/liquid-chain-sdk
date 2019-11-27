brew install unzip
unzip wasi-sdk-7.0.zip
rm -rf /usr/local/opt/wasi-sdk/
cp -r ./wasi-sdk-7.0/opt/ /usr/local/opt/
brew install llvm@9
cd c2ffi/
mkdir build
cd build/
LLVM_DIR=/usr/local/opt/llvm/ cmake ..
make
mv ./bin/c2ffi /usr/local/bin
cd ../..
cd vertex-cdt && ./build.sh
rm -rf ../wasi-sdk-7.0
