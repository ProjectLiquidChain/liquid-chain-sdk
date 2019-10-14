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
build tool to support check rule vertex vm
```bash
cd checker && ./build.sh
```
### Compile to WebAssembly
C language
```bash
clang -Wl,--allow-undefined,--no-entry,--export=<export function> -O3 -s -o <file .wasm> <file .c>
```
 C++ language
```bash
clang++ -Wl,--allow-undefined,--no-entry,--export=<export function> -O3 -s -o <file .wasm> <file .c>
```
rust language
```bash
cargo build --target wasm32-wasi
```
### Create first project in C,C++ language
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
compile the project to WebAssembly
```bash
cargo build --target wasm32-wasi release
```
