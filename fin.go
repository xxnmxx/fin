package fin

import "fmt"

func Compound(p float64, r float64, d int) float64 {
	ret := p
	for i := 0; i <= d; i++ {
		ret = ret * (1 + r)
	}
	return ret
}

func CompoundEach(p float64, r float64, d int) {
	ret := p
	for i := 0; i <= d; i++ {
		ret = ret * (1 + r)
		fmt.Printf("piriod %v:\t%v\n", i, ret)
	}
}
