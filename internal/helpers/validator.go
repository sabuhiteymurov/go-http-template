package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type SqidsID uint64

func (s *SqidsID) UnmarshalJSON(data []byte) error {
	var encoded string
	if err := json.Unmarshal(data, &encoded); err != nil {
		return fmt.Errorf("invalid ID format")
	}

	if encoded == "" {
		*s = 0
		return nil
	}

	nums, err := Decode(encoded)
	if err != nil || len(nums) != 1 {
		return fmt.Errorf("invalid encoded ID")
	}

	*s = SqidsID(nums[0])
	return nil
}

func (s SqidsID) MarshalJSON() ([]byte, error) {
	if s == 0 {
		return []byte("null"), nil
	}

	encoded, err := Encode(uint64(s))
	if err != nil {
		return nil, err
	}

	return json.Marshal(encoded)
}

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	v := validator.New()
	v.RegisterValidation("sqids", validateSqidsID)
	return &CustomValidator{validator: v}
}

func validateSqidsID(fl validator.FieldLevel) bool {
	id, ok := fl.Field().Interface().(SqidsID)
	if !ok {
		return false
	}
	return id != 0
}
