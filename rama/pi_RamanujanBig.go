package main

import (
	"fmt"
	"math/big"
)

func factorial(n *big.Int) *big.Int {
	if n.Cmp(big.NewInt(1)) <= 0 {
		return big.NewInt(1.0)
	}
	return new(big.Int).Mul(n, factorial(new(big.Int).Sub(n, big.NewInt(1))))
}

func pow(x, y *big.Int) *big.Int {
	if y.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(1)
	}
	if y.Cmp(big.NewInt(1)) <= 0 {
		return x
	}
	return new(big.Int).Mul(x, pow(x, new(big.Int).Sub(y, big.NewInt(1))))
}

func main() {
	var (
		qPi, ram    *big.Float
		n, den, num *big.Int
	)
	const prec = 500
	qPi = big.NewFloat(0)
	qPi.SetPrec(prec)
	ram = new(big.Float)
	ram.SetPrec(prec)
	den = new(big.Int)
	num = new(big.Int)
	for n = big.NewInt(0); ; n = new(big.Int).Add(n, big.NewInt(1)) {
		den.Mul(pow(big.NewInt(-1), n), new(big.Int).Mul(factorial(new(big.Int).Mul(big.NewInt(4), n)), new(big.Int).Add(big.NewInt(1123), new(big.Int).Mul(big.NewInt(21460), n))))
		num.Mul(pow(big.NewInt(882), new(big.Int).Add(new(big.Int).Mul(big.NewInt(2), n), big.NewInt(1))), pow(new(big.Int).Mul(pow(big.NewInt(4), n), factorial(n)), big.NewInt(4)))
		denF := new(big.Float).SetInt(den)
		numF := new(big.Float).SetInt(num)
		ram.Quo(denF, numF)
		qPi.Add(qPi, ram)
		pi := new(big.Float).Quo(big.NewFloat(4), qPi)
		fmt.Printf("%v:\t%.300g\tprec=%v\tacc=%v\n", n, pi, pi.Prec(), pi.Acc())
	}
}
