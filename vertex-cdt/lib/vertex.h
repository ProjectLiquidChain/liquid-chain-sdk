#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#define WASI_EXPORT extern "C"
typedef uint8_t byte_t;
#define ADDRESS_SIZE 35
typedef uint8_t address[ADDRESS_SIZE];
typedef int Event;
typedef uint8_t* pointer;
typedef pointer* lparray;
#ifdef __cplusplus
extern "C" {
#endif
extern size_t chain_storage_size_get(byte_t*, size_t);
extern byte_t* chain_storage_get(byte_t*, size_t, byte_t*);
extern void chain_storage_set(byte_t*, size_t, byte_t*, size_t);
extern void chain_print_bytes(byte_t*, size_t);
extern void chain_event_emit(byte_t*);
extern void chain_get_caller(byte_t*);
extern void chain_get_creator(byte_t*);
extern byte_t* chain_invoke(byte_t*, byte_t* params);
lparray to_lparray(char s[]) {
  lparray result = (lparray) malloc(2 * sizeof(pointer));
  result[0] = (pointer)strlen(s);
  result[1] = (pointer)s;
  return result;
}
#ifdef __cplusplus
}
#endif