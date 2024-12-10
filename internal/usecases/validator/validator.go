package validator

import (
	"fmt"
	"reflect"
)

type Validator struct {
	steps []validatorStep
}

func New() *Validator {
	return &Validator{}
}

type validatorStep func() error

var (
	ErrRequiredValidator = fmt.Errorf("required value")
	ErrNotBetweenValue   = fmt.Errorf("not between value")
)

func (v *Validator) Required(t any) *Validator {
	return &Validator{
		steps: append(v.steps, func() error {
			emptyVal := reflect.Zero(reflect.TypeOf(t)).Interface()

			if emptyVal == t {
				return ErrRequiredValidator
			}
			return nil
		}),
	}
}

func (v *Validator) Between(value, minV, maxV int) *Validator {
	return &Validator{
		steps: append(v.steps, func() error {
			if value < minV {
				return ErrNotBetweenValue
			}
			if value > maxV {
				return ErrNotBetweenValue
			}
			return nil
		}),
	}
}

func (v *Validator) Validate() error {
	for _, step := range v.steps {
		if err := step(); err != nil {
			return err
		}
	}
	return nil
}
