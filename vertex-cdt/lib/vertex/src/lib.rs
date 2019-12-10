type address =  [u8;35];
type Event = i32;
extern {
   fn chain_storage_get(x: &[u8]) ->  &[u8] ;
   fn chain_get_caller() -> address ;
   fn chain_get_creator() -> address;
   fn chain_storage_set(x: &[u8], y: &[u8]) ;
}
