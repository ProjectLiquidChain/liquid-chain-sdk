#include <stdint.h>
#include <stdlib.h>
#define WASI_EXPORT extern "C"
typedef uint8_t byte_t;
typedef uint8_t* address;
typedef int Event;
WASI_EXPORT size_t chain_storage_size_get(byte_t*, size_t);
WASI_EXPORT byte_t* chain_storage_get(byte_t*, size_t, byte_t*);
WASI_EXPORT void chain_storage_set(byte_t*, size_t, byte_t*, size_t);
WASI_EXPORT void chain_print_bytes(byte_t*, size_t);
WASI_EXPORT void chain_event_emit(byte_t*);
WASI_EXPORT void chain_get_caller(byte_t*);
WASI_EXPORT void chain_get_creator(byte_t*);
WASI_EXPORT byte_t* chain_invoke(byte_t*, byte_t* params);