package main

import (
	"fmt"
	"math/big"
)

func factorial(n *big.Float) *big.Float {
	if n.Cmp(big.NewFloat(1)) <= 0 {
		return big.NewFloat(1.0)
	}
	return new(big.Float).Mul(n, factorial(new(big.Float).Sub(n, big.NewFloat(1))))
}

func pow(x, y *big.Float) *big.Float {
	if y.Cmp(big.NewFloat(0)) == 0 {
		return big.NewFloat(1)
	}
	if y.Cmp(big.NewFloat(1)) <= 0 {
		return x
	}
	return new(big.Float).Mul(x, pow(x, new(big.Float).Sub(y, big.NewFloat(1))))
}

func main() {
	var n, qPi, den, num, ram *big.Float
	qPi = big.NewFloat(0)
	qPi.SetPrec(300)
	den = new(big.Float)
	den.SetPrec(300)
	num = new(big.Float)
	num.SetPrec(300)
	ram = new(big.Float)
	ram.SetPrec(300)
	for n = big.NewFloat(0); ; n = new(big.Float).Add(n, big.NewFloat(1)) {
		den.Mul(pow(big.NewFloat(-1), n), new(big.Float).Mul(factorial(new(big.Float).Mul(big.NewFloat(4), n)), new(big.Float).Add(big.NewFloat(1123), new(big.Float).Mul(big.NewFloat(21460), n))))
		num.Mul(pow(big.NewFloat(882), new(big.Float).Add(new(big.Float).Mul(big.NewFloat(2), n), big.NewFloat(1))), pow(new(big.Float).Mul(pow(big.NewFloat(4), n), factorial(n)), big.NewFloat(4)))
		ram.Quo(den, num)
		qPi.Add(qPi, ram)
		//fmt.Printf("%v:\t%v / %v -> %v\t%v\n", n, den, num, ram, 4/qPi)
		fmt.Printf("%v:\t%.50g\n", n, new(big.Float).Quo(big.NewFloat(4), qPi))
	}
}
