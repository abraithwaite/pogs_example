@0xe31a02a1aad04e56;
using Go = import "go.capnp";

$Go.package("pogs_example");
$Go.import("pogs_example");


struct Phone {
    number @0 :Text;
    location @1 :Text;
}

struct Address {
    number @0 :Text;
    street @1 :Text;
    zipCode @2 :Int32;
}

struct User {
    id @0 :Int64;
    name @1 :Text;
    union {
        phone @2 :Phone;
        address @3 :Address;
    }
}
