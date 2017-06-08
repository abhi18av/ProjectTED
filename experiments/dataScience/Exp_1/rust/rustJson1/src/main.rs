extern crate serialize;

use serialize::json;

fn main() {
    println!("My very first real world Rust program!");

for x in [1,2,3,4] {
println(x)
}


    println!("{}", json::encode(&vec!["to", "be", "or", "not", "to", "be"]));
    println!("{}", json::encode(&Some(true))); 

}


