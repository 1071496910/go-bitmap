package go_bitmap

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidIndex = errors.New("invalid index")
)

type Bitmap interface {
	Len() int
	Set(index uint) error
	Unset(index uint) error
	Check(index uint) (bool, error)
}

type innerByte struct {
	i byte
}

func (i *innerByte) checkIndexValid(index uint) bool {
	return index >= 0 && index < 8
}

func (i *innerByte) Set(index uint) error {
	if !i.checkIndexValid(index) {
		return ErrInvalidIndex
	}
	i.i = i.i | (byte(1) << index)
	return nil
}

func (i *innerByte) Unset(index uint) error {
	if !i.checkIndexValid(index) {
		return ErrInvalidIndex
	}
	i.i = i.i & ^(byte(1) << index)
	return nil
}

func (i *innerByte) Check(index uint) (bool, error) {
	if !i.checkIndexValid(index) {
		return false, ErrInvalidIndex
	}
	return 0 != (i.i & (byte(1) << index)), nil
}

type bitmap struct {
	storage []innerByte
	len     int
}

func NewBitmap(n int) *bitmap {
	return &bitmap{
		storage: make([]innerByte, n, n),
		len:     n,
	}
}

func (b *bitmap) checkIndexValid(index uint) bool {
	return index >= 0 && index < uint(b.Len())
}

func (b *bitmap) Len() int {
	return b.len
}

func (b *bitmap) Set(index uint) error {
	if !b.checkIndexValid(index) {
		return fmt.Errorf("Valid index")
	}
	return b.storage[index/8].Set(index % 8)
}

func (b *bitmap) Unset(index uint) error {
	if !b.checkIndexValid(index) {
		return fmt.Errorf("Valid index")
	}
	return b.storage[index/8].Unset(index % 8)
}

func (b *bitmap) Check(index uint) (bool, error) {
	if !b.checkIndexValid(index) {
		return false, fmt.Errorf("Valid index")
	}
	return b.storage[index/8].Check(index % 8)
}
