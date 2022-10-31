package ecc

type ECDSA struct {
	GenPoint *Point
	Curve    *MontgomeryCurve
}

func NewECDSA() *ECDSA {
	curve := NewCurve25519()
	return &ECDSA{
		GenPoint: curve.G(),
		Curve:    curve,
	}
}

func (ec *ECDSA) Sign(privateKey int64, message string) (int64, int64) {
	// rand.Seed(time.Now().UnixNano())
	// r, s, randK := int64(0), int64(0), int64(0)
	// for s == 0 {
	// 	for r == 0 {
	// 		// 1. Select a random or pseudorandom integer k, 1 ≤ k ≤ n - 1
	// 		min, max := int64(1), *ec.Curve.N-1
	// 		randK = int64(rand.Intn(int(max-min+1))) + min
	// 		ed25519.PrivateKeySize
	// 		// 2. Compute kG = (x1, y1) and convert x1 to an integer x1
	// 		kG, err := ec.Curve.MulPoint(randK, ec.GenPoint)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		// 3. Compute r = x1 mod n. If r = 0 then go to step 1.
	// 		r = *kG.X % *ec.Curve.N
	// 	}
	//
	// 	// 4. Compute k-1 mod n.
	// 	invK, err := Modinv(randK, *ec.Curve.N)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	// 5. Compute SHA-1(m) and convert this bit string to an integer ec.
	// 	h := sha1.New()
	// 	_, _ = io.WriteString(h, message)
	// 	e := Hex2int(string(h.Sum(nil)))
	//
	// 	// 6. Compute 5 = k-1(ec + dr) mod n. If s = 0 then go to step 1.
	// 	s = invK * (e + privateKey*r) % *ec.Curve.N
	// }
	// // 7. A's signature for the message m is (r, s).
	// return r, s
	return -1, -1
}

func (ec *ECDSA) Verify(publicKey Point, message string, r, s int64) bool {
	// // 1. Verify that r and s are integers in the interval [1, n - 1].
	// if inSlice(r, makeRange(1, *ec.Curve.N-1)) || inSlice(s, makeRange(1, *ec.Curve.N-1)) {
	// 	return false
	// }
	//
	// // 2. Compute SHA-1(m) and convert this bit string to an integer e
	// h := sha1.New()
	// _, _ = io.WriteString(h, message)
	// e := Hex2int(string(h.Sum(nil)))
	//
	// // 3. Compute w = s^-1 mod n.
	//
	// w, err := Modinv(s, *ec.Curve.N)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// // 4. Compute u1 = ew mod n and u2 = rw mod n.
	// u1 := (e * w) % *ec.Curve.N
	// u2 := (r * w) % *ec.Curve.N
	//
	// // 5. Compute X = u1G + u2Q.
	// u1G, err := ec.Curve.MulPoint(u1, ec.GenPoint)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// u2G, err := ec.Curve.MulPoint(u2, &publicKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pointX, err := ec.Curve.AddPoint(u1G, u2G)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// // 6. If X = 0, then reject the signature.
	// // Otherwise, convert the x-coordinate x1 of X to an integer x1, and compute v = x1 mod n.
	// if !ec.Curve.IsOnCurve(pointX) {
	// 	return false
	// }
	// v := *pointX.X % *ec.Curve.N
	//
	// // 7. Accept the signature if and only if u = r.
	// return v == r
	return true
}

func inSlice(a int64, list []int64) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
