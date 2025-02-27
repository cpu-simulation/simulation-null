package bus

type AddressBus struct {
	Address []int
}

func NewAddressBus(size int) *AddressBus {
	return &AddressBus{
		Address: make([]int, size),
	}
}

func (b *AddressBus) Read() int {
	return b.Address[0]
}

func (b *AddressBus) Write(address int) {
	b.Address[0] = address
}
