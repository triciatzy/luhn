package luhn

import "math/big"

// CalculateLuhn return the check number
func CalculateLuhn(number *big.Int) int {
	checkNumber := checksum(number)

	if checkNumber == 0 {
		return 0
	}

	return 10 - checkNumber
}

// Valid check number is valid or not based on Luhn algorithm
func Valid(number *big.Int) bool {
	modResult := new(big.Int)
	divResult := new(big.Int)
	ten := big.NewInt(10)

	modResult.Mod(number, ten)
	checkDigit := checksum(divResult.Div(number, ten))
	summation := modResult.Int64() + int64(checkDigit)

	return summation%10 == 0

}

func checksum(number *big.Int) int {
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

	return int(luhn.Int64())
}
