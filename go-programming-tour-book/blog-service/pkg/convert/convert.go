package convert

import "strconv"

type SrcTo string

func (s SrcTo) String() string {
	return string(s)
}

func (s SrcTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

func (s SrcTo) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s SrcTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

func (s SrcTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}
