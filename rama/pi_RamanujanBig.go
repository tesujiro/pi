package main

import (
	"flag"
	"fmt"
	"math/big"

	"github.com/pkg/profile"
)

//var no_cache = true
var no_cache = false

var (
	c_factorial map[string]*big.Int
	int_minus1  *big.Int   = big.NewInt(-1)
	int_0       *big.Int   = big.NewInt(0)
	int_1       *big.Int   = big.NewInt(1)
	int_2       *big.Int   = big.NewInt(2)
	int_4       *big.Int   = big.NewInt(4)
	int_882     *big.Int   = big.NewInt(882)
	int_1123    *big.Int   = big.NewInt(1123)
	int_21460   *big.Int   = big.NewInt(21460)
	float_0     *big.Float = big.NewFloat(0)
	float_4     *big.Float = big.NewFloat(4)
)

func init() {
	c_factorial = make(map[string]*big.Int)
}

func factorial(n *big.Int) *big.Int {
	if n.Cmp(int_1) <= 0 {
		return int_1
	}
	return new(big.Int).Mul(n, factorial_cached(new(big.Int).Sub(n, int_1)))
}

func factorial_cached(n *big.Int) *big.Int {
	if no_cache {
		return factorial(n)
	}
	if v, ok := c_factorial[n.String()]; ok {
		return v
	}
	v := factorial(n)
	c_factorial[n.String()] = v
	return (v)
}

func pow(x, y *big.Int) *big.Int {
	return new(big.Int).Exp(x, y, nil)
}

func main() {
	defer profile.Start(profile.ProfilePath(".")).Stop()
	var (
		prec = flag.Uint("p", 0, "precision")
	)

	flag.Parse()
	if *prec != 0 {
		pr(pi(*prec))
		return
	}
	var i uint
	for i = 1; ; i++ {
		pr(pi(i))
	}
}

func pr(prec uint, n *big.Int, pi *big.Float) {
	fmt.Printf("%v:\t%."+fmt.Sprintf("%v", prec)+"g\tprec=%v\tacc=%v\n", n, pi, pi.Prec(), pi.Acc())
}

func pi(prec uint) (uint, *big.Int, *big.Float) {
	var (
		pi, prev_pi, qPi, ram *big.Float
		n, den, num           *big.Int
	)
	qPi = big.NewFloat(0)
	qPi.SetPrec(prec)
	prev_pi = new(big.Float)
	prev_pi.SetPrec(prec)
	ram = new(big.Float)
	ram.SetPrec(prec)
	den = new(big.Int)
	num = new(big.Int)
	for n = int_0; ; n = new(big.Int).Add(n, int_1) {
		den.Mul(pow(int_minus1, n), new(big.Int).Mul(factorial_cached(new(big.Int).Mul(int_4, n)), new(big.Int).Add(int_1123, new(big.Int).Mul(int_21460, n))))
		num.Mul(pow(int_882, new(big.Int).Add(new(big.Int).Mul(int_2, n), int_1)), pow(new(big.Int).Mul(pow(int_4, n), factorial_cached(n)), int_4))
		denF := new(big.Float).SetInt(den)
		numF := new(big.Float).SetInt(num)
		ram.Quo(denF, numF)
		qPi.Add(qPi, ram)
		pi = new(big.Float).Quo(float_4, qPi)
		if pi.Cmp(prev_pi) == 0 {
			return prec, n, pi
		} else {
			prev_pi = pi
		}
	}
}
