package search

import (
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const DefaultBoost = float64(0)

type WithBoost interface {
	Boost() float64
	SetBoost(v float64)
}

type BoostParam struct {
	BoostValue *float64 `bson:"boost,omitempty" json:"boost,omitempty"`
}

func (b BoostParam) Boost() float64 {
	if b.BoostValue == nil {
		return DefaultBoost
	}
	return *b.BoostValue
}

func (b *BoostParam) SetBoost(v float64) {
	if b.Boost() != v {
		b.BoostValue = &v
	}
}

func marshalBoost(data []byte, source interface{}) ([]byte, error) {
	if b, ok := source.(WithBoost); ok {
		if b.Boost() != DefaultBoost {
			return sjson.SetBytes(data, "boost", b.Boost())
		}
	}
	return data, nil
}
func unmarshalBoostParam(value gjson.Result, target interface{}) error {
	if r, ok := target.(WithBoost); ok {
		r.SetBoost(value.Num)
	}
	return nil
}
