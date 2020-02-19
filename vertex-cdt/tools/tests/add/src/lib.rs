extern crate vertex_lib;
use vertex_lib::*;
// declare event
extern {
   fn Add(x: i32, y: i32) -> Event;
}
// declare public function
/*
    fn add(x: i32, y:i32)
    return i32
 */
#[no_mangle]
pub fn add(x: i32, y: i32) -> i32
{
    unsafe {
        Add(x,y) ;
    }
    return x + y
}