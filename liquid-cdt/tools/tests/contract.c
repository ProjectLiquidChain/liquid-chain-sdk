#include <string.h>
#include <liquid.h>
extern Event Add(int a, int b);
int add (int a, int b) {
     Add(a,b);
    return a+b ;
}