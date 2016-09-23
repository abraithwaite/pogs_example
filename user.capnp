@0xe31a02a1aad04e56;
using Go = import "go.capnp";

$Go.package("pogs_example");
$Go.import("pogs_example");


struct Phone {
    number @0 :Text;
    location @1 :Text;
}


struct User {
    id @0 :Int64;
    name @1 :Text;
    phone @2 :Phone;
}
