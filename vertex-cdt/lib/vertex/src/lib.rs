use std::mem;
use byteorder::{ByteOrder, LittleEndian};
extern crate libc;
pub type address =  *const u8;
pub type Event = i32;
type pointer =  *mut u8;
pub type lparray =  *mut pointer;
extern {
   fn chain_storage_size_get(key: *const u8, length: u8 ) -> usize;
   fn chain_get_caller(caller: address) ;
   fn chain_get_creator(creator: address);
   fn chain_storage_get(key: *const u8, length: u8, value: *const u8) -> *const u8 ;
   fn chain_storage_set(key: *const u8, key_size: u8, value: *const u8, value_size: u8) ;
}

unsafe fn set(data: *mut pointer, offset: usize, value: *mut u8) {
    let ptr = data.offset(offset as isize) as *mut pointer;
    *ptr = value;
}

pub fn to_lparray (s: &[u8]) -> lparray {
    unsafe {
        let result = libc::malloc(2 * mem::size_of::<*mut pointer>()) as *mut pointer;
        let len = s.len();
        set(result,0, len as *mut u8);
        set(result,1, s.as_ptr() as *mut u8);
        return result
    }
}

pub fn get_caller() -> address {
    let caller: address = unsafe {
        libc::malloc((mem::size_of::<u8>() * 35) as libc::size_t ) as address
    };
    unsafe {
        chain_get_caller(caller) ;
    }
    return caller;
}
pub fn get_creator() -> address {
    let creator: address = unsafe {
        libc::malloc((mem::size_of::<u8>() * 35) as libc::size_t ) as address
    };
    unsafe {
        chain_get_creator(creator) ;
    }
    return creator;
}
pub fn storage_get(key: *const u8, key_size: u8) -> *const u8 {
    let size : usize = unsafe { chain_storage_size_get(key, key_size ) };
    if size == 0 {
        return 0 as *const u8;
    }
    unsafe {
        let ret : *const u8 = libc::malloc((mem::size_of::<u8>() * size) as libc::size_t ) as *const u8;
        chain_storage_get(key, key_size, ret);
        return ret ;
    };
}
pub fn storage_set(key: *const u8, key_size: u8 , value: *const u8, value_size: u8 ) {
    unsafe {
        chain_storage_set(key, key_size, value, value_size ) ;
    }
}
pub fn from_byte(x: *const u8, length: u8) -> u64 {
    let key = unsafe { std::slice::from_raw_parts(x, length as usize) };

    return LittleEndian::read_uint(key, length as usize).to_be();
}
