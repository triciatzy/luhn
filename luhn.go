package luhn

import "math/big"

// CalculateLuhn return the check number
func CalculateLuhn(number *big.Int) *big.Int {
	checkNumber := checksum(number)

	zero := big.NewInt(0)
	if checkNumber.Cmp(zero) == 0 {
		return zero
	}

	result := big.NewInt(10)
	result.Sub(result, checkNumber)

	return result
}

// Valid check number is valid or not based on Luhn algorithm
func Valid(number *big.Int) bool {
	modResult := new(big.Int)
	divResult := new(big.Int)
	summation := new(big.Int)
	ten := big.NewInt(10)

	modResult.Mod(number, ten)
	checkDigit := checksum(divResult.Div(number, ten))
	summation.Add(modResult, checkDigit)
	summation.Mod(summation, ten)

	return summation.Cmp(big.NewInt(0)) == 0
}

func checksum(number *big.Int) *big.Int {
	result := new(big.Int)
	cur := new(big.Int)
	luhn := new(big.Int)
	ten := big.NewInt(10)
	one := big.NewInt(1)
	two := big.NewInt(2)

	for i := big.NewInt(0); number.Cmp(big.NewInt(0)) == 1; i.Add(i, one) {
		cur.Mod(number, ten)
		result.Mod(i, two)
		if result.Int64() == 0 {
			cur.Mul(cur, two)
			if cur.Int64() > 9 {
				a := new(big.Int)
				b := new(big.Int)

				a.Mod(cur, ten)
				b.Div(cur, ten)
				cur.Add(a, b)
			}
		}
		luhn.Add(luhn, cur)
		number.Div(number, ten)
	}

	luhn.Mod(luhn, ten)

	return luhn
}
