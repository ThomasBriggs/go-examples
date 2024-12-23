package BitArray

import (
	"fmt"
	"math/bits"
	"strings"
)

const _BLOCK_SIZE = bits.UintSize

type BitArray struct {
	length uint
	data   []uint
}

func findBlock(position uint) uint {
	return position / _BLOCK_SIZE
}

func shiftAmount(position uint) uint {
	return position - (findBlock(position) * _BLOCK_SIZE)
}

func NewBitArray(size uint) BitArray {
	length := size
	return BitArray{length, make([]uint, findBlock(size)+1)}
}

func (ba BitArray) Length() uint {
	return ba.length
}

func (ba BitArray) setBit(pos uint) {
	ba.data[findBlock(pos)] |= (1 << shiftAmount(pos))
}

func (ba BitArray) unsetBit(pos uint) {
	ba.data[findBlock(pos)] &= ^(1 << shiftAmount(pos))
}

func (ba BitArray) Set(position uint, value bool) {
	if value {
		ba.setBit(position)
	} else {
		ba.unsetBit(position)
	}
}

func (ba BitArray) Get(position uint) bool {
	return (ba.data[findBlock(position)] & (1 << shiftAmount(position))) > 0
}

func (ba BitArray) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(ba.data); i++ {
		sb.WriteString(fmt.Sprintf("%0*b\n", _BLOCK_SIZE, bits.Reverse(ba.data[i])))
	}

	output := sb.String()
	return output[:len(output)-1]
}
