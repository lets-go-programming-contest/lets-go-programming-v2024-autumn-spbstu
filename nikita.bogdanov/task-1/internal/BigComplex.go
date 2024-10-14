package internal

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

type BigComplex struct {
	Real *big.Float
	Imag *big.Float
}

func isValidComplex(input string) bool {
	complexRegex := regexp.MustCompile(`^([-+]?\d*\.?\d+([eE][-+]?\d+)?)([-+]\d*\.?\d+([eE][-+]?\d+)?)i$`)
	return complexRegex.MatchString(input)
}

func ParseComplex(input string) (*BigComplex, bool) {
	if !isValidComplex(input) {
		realPart, _, err := big.ParseFloat(input, 10, 256, big.ToNearestEven)
		if err != nil {
			return nil, false
		}
		return &BigComplex{Real: realPart, Imag: big.NewFloat(0.0)}, true
	}
	input = strings.ReplaceAll(input, " ", "")
	re := regexp.MustCompile(`([-+]?\d*\.?\d+(?:[eE][-+]?\d+)?)([-+]\d*\.?\d+(?:[eE][-+]?\d+)?)i`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 3 {
		return nil, false
	}

	realPart, _, err := big.ParseFloat(matches[1], 10, 256, big.ToNearestEven)
	if err != nil {
		return nil, false
	}

	imagPart, _, err := big.ParseFloat(matches[2], 10, 256, big.ToNearestEven)
	if err != nil {
		return nil, false
	}

	return &BigComplex{Real: realPart, Imag: imagPart}, true
}

func (bc *BigComplex) String() string {
	real_ := bc.Real.Text('f', -1)
	img_ := bc.Imag.Text('f', -1)
	if img_[0] == '-' {
		return fmt.Sprintf("%s%si", real_, img_)
	}
	return fmt.Sprintf("%s+%si", real_, img_)
}

func (bc *BigComplex) Add(other *BigComplex) *BigComplex {
	result := &BigComplex{
		Real: new(big.Float).Add(bc.Real, other.Real),
		Imag: new(big.Float).Add(bc.Imag, other.Imag),
	}
	return result
}

func (bc *BigComplex) Sub(other *BigComplex) *BigComplex {
	return &BigComplex{
		Real: new(big.Float).Sub(bc.Real, other.Real),
		Imag: new(big.Float).Sub(bc.Imag, other.Imag),
	}
}

func (bc *BigComplex) Mul(other *BigComplex) *BigComplex {
	realPart := new(big.Float).Sub(
		new(big.Float).Mul(bc.Real, other.Real),
		new(big.Float).Mul(bc.Imag, other.Imag),
	)

	imagPart := new(big.Float).Add(
		new(big.Float).Mul(bc.Real, other.Imag),
		new(big.Float).Mul(bc.Imag, other.Real),
	)

	return &BigComplex{
		Real: realPart,
		Imag: imagPart,
	}
}

func (bc *BigComplex) Div(other *BigComplex) *BigComplex {
	denom := new(big.Float).Add(
		new(big.Float).Mul(other.Real, other.Real),
		new(big.Float).Mul(other.Imag, other.Imag),
	)

	if denom.Cmp(big.NewFloat(0)) == 0 {
		panic("Ошибка: деление на ноль невозможно")
	}

	realPart := new(big.Float).Quo(
		new(big.Float).Add(
			new(big.Float).Mul(bc.Real, other.Real),
			new(big.Float).Mul(bc.Imag, other.Imag),
		), denom)

	imagPart := new(big.Float).Quo(
		new(big.Float).Sub(
			new(big.Float).Mul(bc.Imag, other.Real),
			new(big.Float).Mul(bc.Real, other.Imag),
		), denom)

	return &BigComplex{
		Real: realPart,
		Imag: imagPart,
	}
}
