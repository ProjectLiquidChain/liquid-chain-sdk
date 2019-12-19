extern crate vertex_lib;
use vertex_lib::*;
extern {
   fn Add(x: i32, y: i32) -> Event;
}
#[no_mangle]
pub fn add(x: i32, y: i32) -> i32
{
    unsafe {
        Add(x,y) ;
    }
    return x + y
}