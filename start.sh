docker stop liquid && docker rm liquid
docker run -d --name liquid  -v $1:/usr/liquid-chain-sdk/  -it liquid-cdt
cp ./compile.sh /usr/local/bin
alias liquid-compile='compile.sh'