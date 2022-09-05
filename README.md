# Luhn Algorithm

Generating check number & validating Luhn numbers in GO

# Usage

```go
import "github.com/triciatzy/luhn"

func main() {
	// Checking if a string is a valid luhn
	n1 := big.NewInt(1111)
	n2 := big.NewInt(79927398713)
	luhn.Valid(n1) //= false
	luhn.Valid(n2) //= true

	n3 := big.NewInt(7992739871)
  luhn.CalculateLuhn(n3) //= 3
}
```
