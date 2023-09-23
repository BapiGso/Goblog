package customw

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

type Validator struct {
	once     sync.Once
	validate *validator.Validate
}

func (c *Validator) Validate(i any) error {
	c.once.Do(func() {
		c.validate = validator.New()
	})
	return c.validate.Struct(i)
}
