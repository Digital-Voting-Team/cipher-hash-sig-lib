package ecc

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
	"math/big"
)

func GetKeyPair(curve MontgomeryCurve) (int64, *Point) {
	_ = genPrivateKey()
	// pubKey := getPublicKey(pKey, curve)
	return -1, &Point{}
}

func genPrivateKey() *big.Int {
	seed := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, seed); err != nil {
		log.Fatal(err)
	}

	return Hex2int(hex.EncodeToString(seed))
}

// func getPublicKey(d int64, curve MontgomeryCurve) *Point {
// 	return curve.G().Mul(d)
// }
