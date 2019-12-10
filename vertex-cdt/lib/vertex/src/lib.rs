pub type address =  [u8;35];
pub type Event = i32;
extern {
   pub fn chain_storage_get(x: &[u8]) ->  &[u8] ;
   pub fn chain_get_caller() -> address ;
   pub fn chain_get_creator() -> address;
   pub fn chain_storage_set(x: &[u8], y: &[u8]) ;
}
