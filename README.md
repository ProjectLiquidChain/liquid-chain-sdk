# liquid-sdk
## Install liquid-sdk from source code
### Clone project
```bash
git clone --recursive https://github.com/QuoineFinancial/liquid-chain-sdk
```
### Build Tool (MacOS and Ubuntu)
build tools to support C, C++, rust language (require install golang > 1.12)
```bash
./build.sh
```
### Build with docker 
```bash
make build
```
### usage for liquid tools
```bash
liquid-cdt --help
```
```bash
NAME:
   smart contract development CLI - liquid-cdt [language option] compile [file]

USAGE:
   liquid-cdt [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
   liquid team

COMMANDS:
   compile  compile c,c++ language file
   build, r   build rust language project
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
#### Compile to WebAssembly with tool
C language
```bash
liquid-cdt compile <file .c> --export-function (-f) <functions name>
```
 C++ language
```bash
liquid-cdt compile <file .cpp> --export-function (-f) <functions name>
```
rust language
```bash
liquid-cdt build <path folder project>
```
#### Compile to WebAssembly with tool docker
config path of project
```bash
./start.sh <path of project>
```
compile project 
```bash
liquid-compile <file.c> <functions name>
```
### Create first project in rust language
create new project
```bash
liquid-cdt init --name <name_of_project>
```
example code in lib.rs :
```rust
#[no_mangle]
pub extern fn add(x: u32, y: u32) -> u32 {
    return x + y
}
```
compile the project to WebAssembly
```bash
liquid-cdt build <name_of_project>
```
### Create first project in C,C++ language
code demo in c demo.c
```c
#include <liquid.h>
Event Add(int num1, int num2);
Event Call(address x);
int sum(int num1, int num2){
   int num3 = num1+num2;
   Add(num1,num2);
   return num3;
}
```
compile demo.c
```bash
liquid-cdt compile demo.c --export-function sum
```
code demo in c++ demo.cpp
```c++
#include <liquid.h>
WASI_EXPORT Event Add(int num1, int num2);
WASI_EXPORT int sum(int num1, int num2){
   int num3 = num1+num2;
   Add(num1,num2);
   return num3;
}
```
compile demo.cpp
```bash
liquid-cdt compile demo.cpp --export-function sum
```