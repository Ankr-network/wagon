package gas

import "math/big"

type GasMetric interface {
  SpendGas(gas *big.Int) bool
}
