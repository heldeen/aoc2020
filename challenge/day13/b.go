package day13

import (
	"fmt"
	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
	"math/big"
	"strings"
)

//Answer: 1058443396696792
func B(challenge *challenge.Input) int {

	//burn one
	<-challenge.Lines()

	var a, n []*big.Int
	for i, e := range strings.Split(<-challenge.Lines(), ",") {
		if e == "x" {
			continue
		}
		a = append(a, big.NewInt(int64(-1*i)))
		n = append(n, big.NewInt(int64(util.MustAtoI(e))))
	}

	b, err := crt(a, n)
	if err != nil {
		panic(err)
	}
	return int(b.Int64())
}

var one = big.NewInt(1)

func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}
