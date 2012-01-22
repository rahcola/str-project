package main

type BitArray struct {
	size int
	arr []uint32
}

func NewBitArray(size int) BitArray {
	s := size / 32
	if s * 32 < size {
		s++
	}
	return BitArray{size, make([]uint32, s)}
}

func (a BitArray) Len() int {
	return a.size
}

func (a BitArray) Set(i int) {
	if i < a.Len() {
		a.arr[i / 32] = a.arr[i / 32] | (1 << uint(i % 32))
	}
}

func (a BitArray) ForEach(do func(int)) {
	index := 0
	for i := 0; i < len(a.arr); i++ {
		word := a.arr[i]
		for k := 0; k < 32; k++ {
			if word & 1 != 0 && index < a.Len() {
				do(index)
			}
			word = word >> 1
			index++
		}
	}
}

func (a BitArray) Union(b BitArray) BitArray {
	r := NewBitArray(a.Len())
	for i := 0; i < len(r.arr); i++ {
		r.arr[i] = a.arr[i] | b.arr[i]
	}
	return r
}

func (a BitArray) Intersection(b BitArray) BitArray {
	r := NewBitArray(a.Len())
	for i := 0; i < len(r.arr); i++ {
		r.arr[i] = a.arr[i] & b.arr[i]
	}
	return r
}