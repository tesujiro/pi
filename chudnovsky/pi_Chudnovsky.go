package main

import (
	"flag"
	"fmt"
	"math/big"

	"github.com/pkg/profile"
)

var (
	int_minus1                  *big.Int   = big.NewInt(-1)
	int_0                       *big.Int   = big.NewInt(0)
	int_1                       *big.Int   = big.NewInt(1)
	int_3                       *big.Int   = big.NewInt(3)
	int_6                       *big.Int   = big.NewInt(6)
	int_12                      *big.Int   = big.NewInt(12)
	int_16                      *big.Int   = big.NewInt(16)
	int_13591409                *big.Int   = big.NewInt(13591409)
	int_545140134               *big.Int   = big.NewInt(545140134)
	int_minus262537412640768000 *big.Int   = big.NewInt(-262537412640768000)
	float_0                     *big.Float = big.NewFloat(0)
	float_426880sq10005         *big.Float = new(big.Float).Mul(big.NewFloat(426880), new(big.Float).Sqrt(big.NewFloat(10005)))
)

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
		//pi(*prec)
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

	// https://en.wikipedia.org/wiki/Chudnovsky_algorithm
	pi := big.NewFloat(0)
	pi.SetPrec(prec)
	prev_pi := new(big.Float)
	prev_pi.SetPrec(prec)

	K := big.NewInt(6)
	L := big.NewInt(13591409)
	M := new(big.Rat).SetInt(int_1)
	X := big.NewInt(1)
	S := new(big.Rat).SetInt(int_13591409)

	for k := int_1; ; k = new(big.Int).Add(k, int_1) {
		//M = new(big.Rat).Mul(M, new(big.Rat).Quo(new(big.Rat).SetInt(new(big.Int).Sub(pow(K, int_3), new(big.Int).Mul(int_16, K))), new(big.Rat).SetInt(pow(k, int_3))))
		M = new(big.Rat).Quo(new(big.Rat).Mul(new(big.Rat).SetInt((new(big.Int).Sub(pow(K, int_3), new(big.Int).Mul(int_16, K)))), M), new(big.Rat).SetInt(pow(k, int_3)))
		L = new(big.Int).Add(L, int_545140134)
		X = new(big.Int).Mul(X, int_minus262537412640768000)
		S = new(big.Rat).Add(S, new(big.Rat).Quo(new(big.Rat).Mul(M, new(big.Rat).SetInt(L)), new(big.Rat).SetInt(X)))
		K = new(big.Int).Add(K, int_12)
		pi = pi.Quo(float_426880sq10005, new(big.Float).SetRat(S))
		//pr(prec, k, pi)
		if pi.Cmp(prev_pi) == 0 {
			return prec, k, pi
		} else {
			prev_pi.Copy(pi)
		}
	}
}
