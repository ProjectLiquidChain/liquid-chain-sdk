extern crate vertex_lib;
use vertex_lib::*;

// declare event
extern {
   fn Add(x: i32, y: i32) -> Event;
   fn Test(str: lparray) -> Event;
}
#[no_mangle]
pub fn add(x: i32, y: i32) -> i32
{
    const TEST: &[u8] = b"TESTEVENT";
    unsafe {
        Add(x,y) ;
        Test(to_lparray(TEST));
    }
    return x + y
}