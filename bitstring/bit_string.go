package bitstring

import "fmt"

type BitString struct {
	BitString []byte //May want to make immutable ?!?
	bitCount  uint
}

func New() *BitString {
	bitString := make([]byte, 4, 8)
	return &BitString{BitString: bitString, bitCount: 0}
}

func NewWithLenAndCap(length, capacity uint) *BitString {
	bitString := make([]byte, length, capacity)
	return &BitString{BitString: bitString, bitCount: 0}
}

func (b *BitString) BitCount() uint {
	return b.bitCount
}

func T() {
	fmt.Println("Okay")
}
