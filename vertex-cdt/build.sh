go build github.com/QuoineFinancial/vertex-sdk/vertex-cdt
mv ./vertex-cdt /usr/local/bin
cp ./lib/vertex.h /usr/local/opt/wasi-sdk/share/wasi-sysroot/include