package internal

import (
	"github.com/shopspring/decimal"
	"math"
	"math/rand"
)

func Pointer[T any](v T) *T {
	return &v
}

type Range struct {
	Left  int
	Right int
}

func NewRange(left, right int) *Range {
	return &Range{
		Left:  left,
		Right: right,
	}
}

func RandValue(r *Range, distribution string) int {
	if r == nil {
		return 0
	}
	switch distribution {
	case NormalDistribution:
		prop := decimal.NewFromFloat(math.MaxFloat64).Add(decimal.NewFromFloat(math.MaxFloat64)).Div(
			decimal.NewFromInt(int64(r.Right - r.Left)))
		randNorm := decimal.NewFromFloat(rand.NormFloat64()).Add(decimal.NewFromFloat(math.MaxFloat64)).Div(prop)
		return int(randNorm.Add(decimal.NewFromInt(int64(r.Left))).IntPart())
	case UniformDistribution:
		return rand.Intn(r.Right-r.Left) + r.Left
	}
	return r.Left
}
