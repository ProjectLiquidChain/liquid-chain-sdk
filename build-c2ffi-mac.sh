brew install llvm
cd c2ffi/
mkdir build
cd build/
LLVM_DIR=/usr/local/opt/llvm/ cmake ..
make
mv ./bin/c2ffi /usr/local/bin