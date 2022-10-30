package ecc

import (
	"errors"
	"fmt"
	"log"
)

type ICurve interface {
	String() string
	G() Point
	INF() Point
	IsOnCurve(P Point) bool
	AddPoint(P, Q Point) (Point, error)
	MulPoint(d int64, P Point) (Point, error)
	NegPoint(P Point) (Point, error)
	ComputeY(x int64) int64
}

type Point struct {
	X     *int64
	Y     *int64
	Curve ICurve
}

func (p *Point) IsAtInfinity() bool {
	return p.X == nil && p.Y == nil
}

func (p *Point) String() string {
	return fmt.Sprintf("X = %v, Y = %v", p.X, p.Y)
}

func (p *Point) Eq(other Point) bool {
	return p.X == other.X && p.X == other.Y
}

func (p *Point) Neg() Point {
	res, err := p.Curve.NegPoint(*p)
	log.Fatal(err)
	return res
}

type Curve struct {
	Name       string
	A, B, P, N *int64
	GX, GY     *int64
}

func (c *Curve) String() string {
	return c.Name
}

func (c *Curve) G() Point {
	return Point{X: c.GX, Y: c.GY, Curve: c}
}

func (c *Curve) INF() Point {
	return Point{nil, nil, c}
}

func (c *Curve) IsOnCurve(P Point) bool {
	if P.Curve != c {
		return false
	}
	return P.IsAtInfinity() || c.isOnCurve()
}

func (c *Curve) isOnCurve() bool {
	return false
}

func (c *Curve) AddPoint(P, Q Point) (Point, error) {
	if (!c.IsOnCurve(P)) || (!c.IsOnCurve(Q)) {
		return Point{}, errors.New("yhe points are not on the curve")
	}
	if P.IsAtInfinity() {
		return Q, nil
	}
	if Q.IsAtInfinity() {
		return P, nil
	}

	if P == Q.Neg() {
		return c.INF(), nil
	}
	if P == Q {
		return c.doublePoint(P), nil
	}

	return c.addPoint(P, Q), nil
}

func (c *Curve) addPoint(P, Q Point) Point {
	return Point{}
}

func (c *Curve) doublePoint(P Point) Point {
	return Point{}
}

func (c *Curve) MulPoint(d int64, P Point) (Point, error) {
	if !c.IsOnCurve(P) {
		return Point{}, errors.New("the point is not on the curve")
	}
	if P.IsAtInfinity() || d == 0 {
		return c.INF(), nil
	}

	var err error
	res := c.INF()
	isNegScalar := d < 0
	if isNegScalar {
		d = -d
	}
	tmp := P
	for d != 0 {
		if d&0x1 == 1 {
			res, err = c.AddPoint(res, tmp)
			if err != nil {
				return Point{}, err
			}
		}
		tmp, err = c.AddPoint(tmp, tmp)
		if err != nil {
			return Point{}, err
		}
		d >>= 1
	}
	if isNegScalar {
		return res.Neg(), nil
	}
	return res, nil
}

func (c *Curve) NegPoint(P Point) (Point, error) {
	if !c.IsOnCurve(P) {
		return Point{}, errors.New("the point is not on the curve")
	}
	if P.IsAtInfinity() {
		return c.INF(), nil
	}

	return c.negPoint(P), nil
}

func (c *Curve) negPoint(P Point) Point {
	return Point{}
}

func (c *Curve) ComputeY(x int64) int64 {
	return -1
}

// TODO: encode and decode functions

// MontgomeryCurve
// by^2 = x^3 + ax^2 + x
// https://en.wikipedia.org/wiki/Montgomery_curve
type MontgomeryCurve struct {
	Curve
}

func (mc *MontgomeryCurve) isOnCurve(P Point) bool {
	left := *mc.B * *P.Y * *P.Y
	right := (*P.X * *P.X * *P.X) + (*mc.A * *P.X * *P.X) + *P.X
	return (left-right)%*mc.P == 0
}

func (mc *MontgomeryCurve) addPoint(P, Q Point) Point {
	// s = (yP - yQ) / (xP - xQ)
	// xR = b * s^2 - a - xP - xQ
	// yR = yP + s * (xR - xP)
	deltaX := *P.X - *Q.X
	deltaY := *P.Y - *Q.Y
	modInv, err := Modinv(deltaX, *mc.P)
	if err != nil {
		log.Fatal(err)
	}
	s := deltaY * modInv
	resX := (*mc.B*s*s - *mc.A - *P.X - *Q.X) % *mc.P
	resY := (*P.Y + s*(resX-*P.X)) % *mc.P
	return (&Point{&resX, &resY, mc}).Neg()
}

func (mc *MontgomeryCurve) doublePoint(P Point) Point {
	// s = (3 * xP^2 + 2 * a * xP + 1) / (2 * b * yP)
	// xR = b * s^2 - a - 2 * xP
	// yR = yP + s * (xR - xP)
	up := 3**P.X**P.X + 2**mc.A**P.X + 1
	down := 2 * *mc.B * *P.Y
	modInv, err := Modinv(down, *mc.P)
	if err != nil {
		log.Fatal(err)
	}
	s := up * modInv
	resX := (*mc.B*s*s - *mc.A - 2**P.X) % *mc.P
	resY := (*P.Y + s*(resX-*P.X)) % *mc.P
	return (&Point{&resX, &resY, mc}).Neg()
}

func (mc *MontgomeryCurve) negPoint(P Point) Point {
	py := -(*P.Y) % *mc.P
	return Point{P.X, &py, mc}
}

func (mc *MontgomeryCurve) ComputeY(x int64) int64 {
	right := (x*x*x + *mc.A*x*x + x) % *mc.P
	invB, err := Modinv(*mc.B, *mc.P)
	if err != nil {
		log.Fatal(err)
	}
	right = (right * invB) % *mc.P
	y := Modsqrt(right, *mc.P)
	return y
}

func NewCurve25519() ICurve {
	a := int64(486662)
	b := int64(1)
	p := Hex2int("0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed")
	n := Hex2int("0x1000000000000000000000000000000014def9dea2f79cd65812631a5cf5d3ed")
	gx := Hex2int("0x9")
	gy := Hex2int("0x20ae19a1b8a086b4e01edd2c7748d14c923d4d7e6d7c61b229e9c5a27eced3d9")
	return &MontgomeryCurve{
		Curve{
			Name: "Curve25519",
			A:    &a,
			B:    &b,
			P:    &p,
			N:    &n,
			GX:   &gx,
			GY:   &gy,
		},
	}
}
