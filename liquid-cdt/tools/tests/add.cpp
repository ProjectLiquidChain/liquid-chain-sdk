#include <liquid.h>
WASI_EXPORT Event Add(int a, int b);
WASI_EXPORT int add (int a, int b) {
    Add(a,b);
    return a+b ;
}