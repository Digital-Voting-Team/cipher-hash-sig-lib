package ecc

import (
	"fmt"
	"math/big"
)

type Point struct {
	X     *big.Int
	Y     *big.Int
	Curve Curve
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
	return p.Curve.NegPoint(*p)
}

type Curve struct {
	Name       string
	A, B, P, N *big.Int
	GX, GY     *big.Int
}

func (c *Curve) String() string {
	return c.Name
}

func (c *Curve) NegPoint(P Point) Point {
	return Point{}
}

func (c *Curve) G() Point {
	return Point{X: c.GX, Y: c.GY, Curve: *c}
}
