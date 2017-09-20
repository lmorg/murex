package sexp

import (
	"errors"
	"fmt"
)

type Sexp struct {
	data interface{}
}

func New(data ...interface{}) (*Sexp, error) {
	return &Sexp{data}, nil
}

func Parse(data []byte) (*Sexp, error) {
	exp, err := Unmarshal(data)
	return cast(exp), err
}

func cast(v interface{}) *Sexp {
	if s, ok := v.(*Sexp); ok {
		return s
	}
	return &Sexp{v}
}

func (s *Sexp) Encode(canonical bool) ([]byte, error) {
	a, err := s.Array()
	if err != nil {
		return nil, err
	}
	return Marshal(a, canonical)
}

func (s *Sexp) Bytes() ([]byte, error) {
	if b, ok := (s.data).([]byte); ok {
		return b, nil
	}
	return []byte(fmt.Sprint(s.data)), nil
}

func (s *Sexp) Array() ([]interface{}, error) {
	if a, ok := (s.data).([]interface{}); ok {
		return a, nil
	}
	return nil, errors.New("type assertion to []interface{} failed")
}

func (s *Sexp) Push(data ...interface{}) error {
	a, err := s.Array()
	if err != nil {
		return err
	}
	s.data = append(a, data...)
	return nil
}

func (s *Sexp) Pop() (*Sexp, error) {
	a, err := s.Array()
	if err != nil {
		return cast(nil), err
	}
	v := a[len(a)-1]
	s.data = a[:len(a)-1]
	return cast(v), nil
}

func (s *Sexp) At(index int) *Sexp {
	if a, err := s.Array(); err == nil {
		if len(a) > index {
			return cast(a[index])
		}
	}
	return cast(nil)
}
