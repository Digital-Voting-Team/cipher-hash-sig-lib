package ecc

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func Egdc(a, b int64) (int64, int64, int64) {
	if a == 0 {
		return b, 0, 1
	}
	g, y, x := Egdc(b%a, a)
	return g, x - (b/a)*y, y
}

func Modinv(a, m int64) (int64, error) {
	a = a % m
	g, x, _ := Egdc(a, m)
	if g != 1 {
		return -1, errors.New("modular inverse does not exist")
	}
	res := x % m
	if res < 0 {
		res += m
	}
	return res, nil
}

// Modsqrt
// Find a quadratic residue (mod p) of 'a'. p must be an odd prime.
// Solve the congruence of the form: x^2 = a (mod p)
// And returns x. Note that p - x is also a root.
// 0 is returned is no square root exists for these a and p.
// The Tonelli-Shanks algorithm is used (except for some simple cases in which the solution
// is known from an identity). This algorithm runs in polynomial time (unless the
// generalized Riemann hypothesis is false).
func Modsqrt(a, p int64) int64 {
	// Simple cases
	if legendreSymbol(a, p) != 1 {
		return 0
	} else if a == 0 {
		return 0
	} else if p == 2 {
		return p
	} else if p%4 == 3 {
		return int64(math.Pow(float64(a), float64((p+1)/4))) % p
	}

	// Partition p-1 to s * 2^e for an odd s (i.e.
	// reduce all the powers of 2 from p-1)
	s, e := p-1, 0
	for s%2 == 0 {
		s /= 2
		e += 1
	}

	// Find some 'n' with a legendre symbol n|p = -1.
	// Shouldn't take long.
	n := 2
	for legendreSymbol(int64(n), p) != -1 {
		n += 1
	}

	// Here be dragons!
	// Read the paper "Square roots from 1; 24, 51,
	// 10 to Dan Shanks" by Ezra Brown for more information

	// x is a guess of the square root that gets better with each iteration.
	// b is the "fudge factor" - by how much we're off with the guess. The invariant x^2 = ab (mod p)
	// is maint64ained throughout the loop.
	// g is used for successive powers of n to update both a and b
	// r is the exponent - decreases with each update
	x := int64(math.Pow(float64(a), float64((s+1)/2))) % p
	b := int64(math.Pow(float64(a), float64(s))) % p
	g := int64(math.Pow(float64(n), float64(s))) % p
	r := e

	for {
		t, m := b, 0
		for _, m = range makeRange(0, int64(r)) {
			if t == 1 {
				break
			}
			t = int64(math.Pow(float64(t), 2)) % p
		}
		if m == 0 {
			return x
		}

		gs := int64(math.Pow(float64(g), math.Pow(2, float64(r-m-1)))) % p
		g = (gs * gs) % p
		x = (x * gs) % p
		b = (b * g) % p
		r = m
	}
}

func makeRange(min, max int64) []int64 {
	a := make([]int64, max-min+1)
	for i := range a {
		a[i] = min + int64(i)
	}
	return a
}

// Compute the Legendre symbol a|p using
// Euler's criterion. p is a prime, a is relatively prime to p (if p divides a, then a|p = 0)
// Returns 1 if a has a square root modulo p, -1 otherwise.
func legendreSymbol(a, p int64) int64 {
	ls := int64(math.Pow(float64(a), float64((p-1)/2))) % p
	if ls == p-1 {
		return -1
	}
	return ls
}

func IntLenInByte(n int64) int64 {
	length := 0
	for n != 0 {
		n >>= 8
		length += 1
	}
	return int64(length)
}

func Hex2int(hexStr string) int64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return int64(result)
}
