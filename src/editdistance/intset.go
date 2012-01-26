package editdistance

type BitArray []uint32

func NewBitArray(size int) BitArray {
	s := size / 32
	if s * 32 < size {
		s++
	}
	return make([]uint32, 0, s)
}

func (a BitArray) Set(i int) BitArray {
	for i >= len(a) * 32 {
		a = append(a, 0)
	}
	a[i / 32] = a[i / 32] | (1 << uint(i % 32))
	return a
}

func (a BitArray) ForEach(do func(int)) {
	index := 0
	for i := 0; i < len(a); i++ {
		word := a[i]
		for k := 0; k < 32; k++ {
			if word & 1 != 0 {
				do(index)
			}
			word = word >> 1
			index++
		}
	}
}

func (a BitArray) Union(b BitArray) BitArray {
	l := len(a)
	if len(b) > l {
		l = len(b)
	}
	r := make([]uint32, l)
	for i := 0 ; i < len(r); i++ {
		if i >= len(a) {
			r[i] = b[i]
		} else if i >= len(b) {
			r[i] = a[i]
		} else {
			r[i] = a[i] | b[i]
		}
	}
	return r
}

func (a BitArray) Intersection(b BitArray) BitArray {
	l := len(a)
	if len(b) < l {
		l = len(b)
	}
	r := make([]uint32, l)
	for i := 0; i < len(r); i++ {
		r[i] = a[i] & b[i]
	}
	return r
}