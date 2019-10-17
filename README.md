# Vertex-sdk
## Install vertex-sdk from source code
### Clone project
```bash
git clone https://github.com/QuoineFinancial/vertex-sdk
```
clone project wasi-sdk and projects config, llvm-project, wasi-libc in src (checkout to versions)
### Build Tool
build tools to support C, C++ language
```bash
brew install ninja && ./build-c.sh
```
build tool to support rust langguage
```bash
./build-rust.sh
```
build vertex tool
```bash
cd vertex-cdt && ./build.sh
```
usage for vertex tools
```bash
vertex-cdt --help
```
```bash
NAME:
   smart contract development CLI - vertex-cdt [language option] compile [file]

USAGE:
   vertex-cdt [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
   vertex team

COMMANDS:
   c++, c++  compile c++ language file
   c, c      compile c language file
   rust, r   compile rust language file
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
### Compile to WebAssembly
C language
```bash
vertex-cdt c <file .c>
```
 C++ language
```bash
vertex-cdt c++ <file .c>
```
rust language
```bash
vertex-cdt rust <path folder project>
```
### Create first project in rust language
create new project
```bash
cargo new --lib <my_project> && cd <my_project>
```
Add the following configurations to Cargo.toml
```toml
[lib]
crate-type =["cdylib"]

[profile.release]
lto = true #link time optimization to shrink wasm size
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
vertex-cdt rust <my_project>
```
### Create first project in C,C++ language
code demo in c++ demo.cpp
```c++
extern "C" int sum(int num1, int num2){
   int num3 = num1+num2;
   return num3;
}
int main(){
   int x = sum(1,99);
   return 0;
}
```
compile demo.cpp
```bash
vertex-cdt c++ demo.cpp
```