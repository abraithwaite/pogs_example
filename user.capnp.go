package pogs_example

// AUTO GENERATED - DO NOT EDIT

import (
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Phone struct{ capnp.Struct }

// Phone_TypeID is the unique identifier for the type Phone.
const Phone_TypeID = 0xba64fe870dff5373

func NewPhone(s *capnp.Segment) (Phone, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Phone{st}, err
}

func NewRootPhone(s *capnp.Segment) (Phone, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Phone{st}, err
}

func ReadRootPhone(msg *capnp.Message) (Phone, error) {
	root, err := msg.RootPtr()
	return Phone{root.Struct()}, err
}

func (s Phone) String() string {
	str, _ := text.Marshal(0xba64fe870dff5373, s.Struct)
	return str
}

func (s Phone) Number() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Phone) HasNumber() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Phone) NumberBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Phone) SetNumber(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Phone) Location() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Phone) HasLocation() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Phone) LocationBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Phone) SetLocation(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// Phone_List is a list of Phone.
type Phone_List struct{ capnp.List }

// NewPhone creates a new list of Phone.
func NewPhone_List(s *capnp.Segment, sz int32) (Phone_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Phone_List{l}, err
}

func (s Phone_List) At(i int) Phone { return Phone{s.List.Struct(i)} }

func (s Phone_List) Set(i int, v Phone) error { return s.List.SetStruct(i, v.Struct) }

// Phone_Promise is a wrapper for a Phone promised by a client call.
type Phone_Promise struct{ *capnp.Pipeline }

func (p Phone_Promise) Struct() (Phone, error) {
	s, err := p.Pipeline.Struct()
	return Phone{s}, err
}

type Address struct{ capnp.Struct }

// Address_TypeID is the unique identifier for the type Address.
const Address_TypeID = 0xc4fb68e5d0f3a192

func NewAddress(s *capnp.Segment) (Address, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Address{st}, err
}

func NewRootAddress(s *capnp.Segment) (Address, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Address{st}, err
}

func ReadRootAddress(msg *capnp.Message) (Address, error) {
	root, err := msg.RootPtr()
	return Address{root.Struct()}, err
}

func (s Address) String() string {
	str, _ := text.Marshal(0xc4fb68e5d0f3a192, s.Struct)
	return str
}

func (s Address) Number() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Address) HasNumber() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Address) NumberBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Address) SetNumber(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Address) Street() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Address) HasStreet() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Address) StreetBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Address) SetStreet(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s Address) ZipCode() int32 {
	return int32(s.Struct.Uint32(0))
}

func (s Address) SetZipCode(v int32) {
	s.Struct.SetUint32(0, uint32(v))
}

// Address_List is a list of Address.
type Address_List struct{ capnp.List }

// NewAddress creates a new list of Address.
func NewAddress_List(s *capnp.Segment, sz int32) (Address_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Address_List{l}, err
}

func (s Address_List) At(i int) Address { return Address{s.List.Struct(i)} }

func (s Address_List) Set(i int, v Address) error { return s.List.SetStruct(i, v.Struct) }

// Address_Promise is a wrapper for a Address promised by a client call.
type Address_Promise struct{ *capnp.Pipeline }

func (p Address_Promise) Struct() (Address, error) {
	s, err := p.Pipeline.Struct()
	return Address{s}, err
}

type User struct{ capnp.Struct }
type User_Which uint16

const (
	User_Which_phone   User_Which = 0
	User_Which_address User_Which = 1
)

func (w User_Which) String() string {
	const s = "phoneaddress"
	switch w {
	case User_Which_phone:
		return s[0:5]
	case User_Which_address:
		return s[5:12]

	}
	return "User_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// User_TypeID is the unique identifier for the type User.
const User_TypeID = 0xe66193ff15e52403

func NewUser(s *capnp.Segment) (User, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return User{st}, err
}

func NewRootUser(s *capnp.Segment) (User, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return User{st}, err
}

func ReadRootUser(msg *capnp.Message) (User, error) {
	root, err := msg.RootPtr()
	return User{root.Struct()}, err
}

func (s User) String() string {
	str, _ := text.Marshal(0xe66193ff15e52403, s.Struct)
	return str
}

func (s User) Which() User_Which {
	return User_Which(s.Struct.Uint16(8))
}
func (s User) Id() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s User) SetId(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s User) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s User) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s User) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s User) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s User) Phone() (Phone, error) {
	p, err := s.Struct.Ptr(1)
	return Phone{Struct: p.Struct()}, err
}

func (s User) HasPhone() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s User) SetPhone(v Phone) error {
	s.Struct.SetUint16(8, 0)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewPhone sets the phone field to a newly
// allocated Phone struct, preferring placement in s's segment.
func (s User) NewPhone() (Phone, error) {
	s.Struct.SetUint16(8, 0)
	ss, err := NewPhone(s.Struct.Segment())
	if err != nil {
		return Phone{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s User) Address() (Address, error) {
	p, err := s.Struct.Ptr(1)
	return Address{Struct: p.Struct()}, err
}

func (s User) HasAddress() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s User) SetAddress(v Address) error {
	s.Struct.SetUint16(8, 1)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewAddress sets the address field to a newly
// allocated Address struct, preferring placement in s's segment.
func (s User) NewAddress() (Address, error) {
	s.Struct.SetUint16(8, 1)
	ss, err := NewAddress(s.Struct.Segment())
	if err != nil {
		return Address{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// User_List is a list of User.
type User_List struct{ capnp.List }

// NewUser creates a new list of User.
func NewUser_List(s *capnp.Segment, sz int32) (User_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2}, sz)
	return User_List{l}, err
}

func (s User_List) At(i int) User { return User{s.List.Struct(i)} }

func (s User_List) Set(i int, v User) error { return s.List.SetStruct(i, v.Struct) }

// User_Promise is a wrapper for a User promised by a client call.
type User_Promise struct{ *capnp.Pipeline }

func (p User_Promise) Struct() (User, error) {
	s, err := p.Pipeline.Struct()
	return User{s}, err
}

func (p User_Promise) Phone() Phone_Promise {
	return Phone_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p User_Promise) Address() Address_Promise {
	return Address_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

const schema_e31a02a1aad04e56 = "x\xdad\x92?k\x14A\x18\xc6\x9f\xe7\x9d]#\xe4" +
	"\xc4]66\xc1C\x02\xb1P\x8c\xa8\x95\xa4I\xfc\x13" +
	"PQ\xb9\x17\xffT\x16\xae\xd9\x81[\xc8\xed.;\x1b" +
	"\x04\x9b\xa0 \xa2\xa5vvim\xec\xad\xfd\x00\xf9\x0e" +
	"\x12\xf3\x01,\x85\x8cL\x8e\xbb\x0b\xa6\x9by\xe6\x81\xdf" +
	"o\xde\x99\xe4`]\xae\xc7\xb5\x00z>>\xe5\xdd\x13" +
	"\x7f\xe6\xc3a\xf1\x03\xe9<\xfd\xf3\xc7{\xdfve\xf1" +
	"\x17b\x99\x03\xb2>?e\x17\x19VK|\x0d\xfa\xcf" +
	"\xbb\x7f\xf6\xf6\x87\x7f\x7fB\xe7y\xa2\xfc\x96_\xb3\x8f" +
	"G\xe5\xf7\xfc\x0ez\xb3\xbc\x7f\xce\x7f\xc9\x7f\x87\xb2\xcc" +
	"\xca\x1b2'\x8c\xb2\xbe\xbc\xcb\x96\xc6\x109\xc0\x8a\xdf" +
	"v\xb6\xbd\xba\x997\xac\x9a\xd5\xc1\xb0\xae,0 \xf5" +
	"\xb4\x89\x80\x88@zi\x15\xd0eC\xbd&L\xc9\x05" +
	"\x86p\xe5\x01\xa0W\x0c\xf5\xa6p\xad\xda\x1e\xbd\xb2-" +
	"{\x10\xf6@\xbfUo\xe6]YW\x00\xa6\xd9q\xca" +
	"\xad\xa2h\xadq.`zS\xccF\xc0\xac\x1b\xea\xc3" +
	"c\x98\xfb!\xbck\xa8\x03!e\x81\x02\xa4\x8fn\x03" +
	"z\xcfP\x9f\x9e@\xaf\xb9\xae\xb5\xb6\x9blw\xde\x94" +
	"\xcd\x9d\xba\xb0\x8c \x8c\xfe\xb3x\xe6l{t\xd3d" +
	"\xaa\x90/\x02\xfa\xc2P\x87\xc2\x89\x81\xbd\x0c\xe8KC" +
	"\xdd\x12\xf6\xc5{\x8e%\xca\x1b\x80\x16\x86\xda\x08\xfb\xe6" +
	"0\xc4\x06HG\xc1mh\xa8\x9d\xd0\x94\x05c\x08c" +
	"\xf0l\x95\x8f\xec\xc4\xeaB\x13\xa6\xccd\xf6\x01@&" +
	"\xe0N\x1e\xe6\xe2\x1c\x93\xd9s\x8fO\xfe\x05\x00\x00\xff" +
	"\xff\xc66~\x01"

func init() {
	schemas.Register(schema_e31a02a1aad04e56,
		0xba64fe870dff5373,
		0xc4fb68e5d0f3a192,
		0xe66193ff15e52403)
}
