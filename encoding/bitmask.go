package encoding

import (
	"fmt"
	"math/big"
)

func Generate(bits []int64) string {
	bitmask := big.NewInt(0)

	for _, bit := range bits {
		bigIntBit := big.NewInt(bit)
		bigIntTwo := big.NewInt(2)
		twoToPowerOfBit := bigIntTwo.Exp(bigIntTwo, bigIntBit, nil)
		bitmask = bitmask.Or(bitmask, twoToPowerOfBit)
	}

	return bitmask.String()
}

func ExtractBitMask(bitmask string) []int64 {
	activeRuleID := []int64{}
	a := big.NewInt(0)
	a.SetString(bitmask, 10)
	bit := fmt.Sprintf("%b", a)
	//start from right to the left, because binary always from right.
	//if we reverse the string, maybe we can do from left to right (from 1 to len(bit) )
	//we start from 2^1 not 2^0
	for i := len(bit) - 2; i >= 0; i-- {
		if bit[i] == byte('1') {
			//because in binary, last array = first.
			activeRuleID = append(activeRuleID, int64(len(bit)-i)-1)
		}
	}
	return activeRuleID
}

//cannot handle when bitmask is below 0, so please dont send bitmask string -1 or below
func Has(bitmask string, bit int64) bool {

	if bit < 0 {
		return false
	}

	var bigIntBitmask big.Int
	bigIntBitmask.SetString(bitmask, 10)

	bigIntbit := big.NewInt(bit)
	bigIntTwo := big.NewInt(2)

	twoToPowerOfBit := bigIntTwo.Exp(bigIntTwo, bigIntbit, nil)
	isHas := bigIntBitmask.And(&bigIntBitmask, twoToPowerOfBit)

	return isHas.Cmp(twoToPowerOfBit) == 0
}

func Toggle(bitmask string, bits []int64) string {

	if len(bits) <= 0 {
		return bitmask
	}

	bigIntBitmask := big.NewInt(0)
	bigIntBitmask.SetString(bitmask, 10)

	for _, bit := range bits {
		bigIntBit := big.NewInt(bit)
		bigIntTwo := big.NewInt(2)
		twoToPowerOfBit := bigIntTwo.Exp(bigIntTwo, bigIntBit, nil)
		bigIntBitmask = bigIntBitmask.Xor(bigIntBitmask, twoToPowerOfBit)
	}

	return bigIntBitmask.String()
}

func Clear(bitmask string, bits []int64) string {
	if len(bits) <= 0 {
		return bitmask
	}

	bigIntBitmask := big.NewInt(0)
	bigIntBitmask.SetString(bitmask, 10)

	for _, bit := range bits {
		bigIntBit := big.NewInt(bit)
		bigIntTwo := big.NewInt(2)
		twoToPowerOfBit := bigIntTwo.Exp(bigIntTwo, bigIntBit, nil)
		bigIntBitmask = bigIntBitmask.AndNot(bigIntBitmask, twoToPowerOfBit)
	}

	return bigIntBitmask.String()

}
