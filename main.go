package main

import (
	"fmt"

	"github.com/Digital-Voting-Team/cipher-hash-sig-lib/ecc"
)

func main() {
	sign := ecc.NewECDSA()
	curve := ecc.NewCurve25519()
	pk1, pbk1 := ecc.GetKeyPair(*curve)
	_, pbk2 := ecc.GetKeyPair(*curve)

	msg := "String ...."
	r, s := sign.Sign(pk1, msg)
	fmt.Println(sign.Verify(*pbk1, msg, r, s))
	fmt.Println(sign.Verify(*pbk2, msg, r, s))
}
