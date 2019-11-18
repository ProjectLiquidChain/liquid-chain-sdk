go get github.com/wasmerio/go-ext-wasm/wasmer
go get github.com/urfave/cli
go build vertex-cdt.go
mv ./vertex-cdt /usr/local/bin
cp ./lib/vertex.h /usr/local/opt/wasi-sdk/share/wasi-sysroot/include