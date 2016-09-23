package main

import (
	"fmt"

	"github.com/abraithwaite/pogs_example"
	capnp "zombiezen.com/go/capnproto2"
	"zombiezen.com/go/capnproto2/pogs"
)

type Phone struct {
	Number   string
	Location string
}

type User struct {
	Which Phone `capnp:"which=phone"`
	Id    int64
	Name  string
}

func main() {
	_, seg, err := capnp.NewMessage(capnp.MultiSegment(nil))
	if err != nil {
		panic(err)
	}
	usr, err := users.NewRootUser(seg)
	if err != nil {
		panic(err)
	}
	usr.SetName("Alan")
	usr.SetId(12412)
	phone, err := users.NewPhone(seg)
	if err != nil {
		panic(err)
	}
	phone.SetNumber("123-456-7890")
	phone.SetLocation("Taipei")
	usr.SetPhone(phone)

	fmt.Println(usr.Id())
	ph, _ := usr.Phone()
	loc, _ := ph.Location()
	fmt.Println(loc)

	u := new(User)
	err = pogs.Extract(u, users.User_TypeID, usr.Struct)
	fmt.Println("error:", err)
	fmt.Println(u)
	fmt.Println(u.Which.Location)
}
