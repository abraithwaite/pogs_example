package users

// AUTO GENERATED - DO NOT EDIT

import (
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

type User struct{ capnp.Struct }

// User_TypeID is the unique identifier for the type User.
const User_TypeID = 0xe66193ff15e52403

func NewUser(s *capnp.Segment) (User, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return User{st}, err
}

func NewRootUser(s *capnp.Segment) (User, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
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
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewPhone sets the phone field to a newly
// allocated Phone struct, preferring placement in s's segment.
func (s User) NewPhone() (Phone, error) {
	ss, err := NewPhone(s.Struct.Segment())
	if err != nil {
		return Phone{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// User_List is a list of User.
type User_List struct{ capnp.List }

// NewUser creates a new list of User.
func NewUser_List(s *capnp.Segment, sz int32) (User_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
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

const schema_e31a02a1aad04e56 = "x\xda\\\xd0\xb1K\xc3`\x10\x05\xf0\xf7\xeekTh" +
	"\xc5~\xb4\xa2tq\xe9\xa2XQ'\xe9b\x07\x05\x15" +
	"\x95\x1cZ\xc11\xb6\x01\x0b6\x09\x89\xc5\xddEpu" +
	"truq\xf7\xaf\xf0\x7f\x10\xdd\x1d\x1c\\\xfc$J" +
	"\xab8\xdd\xf18\xf8q\xaf|\xdb\x92\x15oF\x00\x9d" +
	"\xf5\xc6\\v\xe0&\xaf>\xbb\x8f\xb0E\xba\xa3\xfd\xa7" +
	"\xfb;\xa9=\xc3\x93q\xc0\xbe]\xdb\x8f|\xbe_\x80" +
	"\xce\xd4_\xa6\xddM\xf0\x0a-\xf2\xffe\xa5\xcd\xcb\xca" +
	"1\x7f\xb6\x074\xdc \x0b\xd3\xa5N\x900J\x9a\xfe" +
	"i\x1c\x85\x80O\xea\x84)\x00\x05\x02v\xbe\x09h\xdd" +
	"P\x97\x85\x96\xac2\x0f\x1b;\x80.\x1a\xea\x9ap=" +
	"\x1a\xf4O\xc2\x94%\x08K\xa0;\x8b;\xc1y/\x8e" +
	"\x00\x8c\xb2\xbfJ;\x0b\xd3o\xa3426k\x80\xb6" +
	"\x0cuW8$\xb6\x17\x00\xdd0T_h\x85U\x0a" +
	"`\xf7V\x01\xdd2\xd4C\xa1\xe9u\xe9A\xe8\x81S" +
	"Q\xd0\x0f\x87\xd8\\\x92\xbf\xc1\xf2oa \xcb\xe0W" +
	"\x00\x00\x00\xff\xffwPF\xef"

func init() {
	schemas.Register(schema_e31a02a1aad04e56,
		0xba64fe870dff5373,
		0xe66193ff15e52403)
}
