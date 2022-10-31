package ecc

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/big"
)

func GetKeyPair(curve MontgomeryCurve) (big.Int, *Point) {
	pKey := genPrivateKey()
	fmt.Println(pKey)
	pubKey := getPublicKey(new(big.Int).Set(&pKey), curve)
	fmt.Println(pKey)
	return pKey, pubKey
}

func genPrivateKey() big.Int {
	seed := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, seed); err != nil {
		log.Fatal(err)
	}

	return *Hex2int(hex.EncodeToString(seed))
}

func getPublicKey(d *big.Int, curve MontgomeryCurve) *Point {
	return curve.G().Mul(d)
}
