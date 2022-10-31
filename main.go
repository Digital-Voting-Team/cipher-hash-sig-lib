package main

import (
	"fmt"

	"github.com/Digital-Voting-Team/cipher-hash-sig-lib/ecc"
)

func main() {
	sign := ecc.NewECDSA()
	curve := ecc.NewCurve25519()
	pk1, pbk1 := ecc.GetKeyPair(*curve)
	pk2, pbk2 := ecc.GetKeyPair(*curve)
	fmt.Println(pk1, "\n -- \n", pbk1, "\n -- \n", pk2, "\n -- \n", pbk2)

	msg := "String ...."
	r, s := sign.Sign(pk1, msg)
	fmt.Println(sign.Verify(*pbk1, msg, r, s))
	fmt.Println(sign.Verify(*pbk2, msg, r, s))
}
