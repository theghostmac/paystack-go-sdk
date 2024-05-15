package paystack_transactions

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// CustomInt is a custom type to handle JSON fields that might be returned as strings or integers.
type CustomInt int

// UnmarshalJSON implements the json.Unmarshaler interface for CustomInt.
func (ci *CustomInt) UnmarshalJSON(data []byte) error {
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	switch v := value.(type) {
	case float64:
		*ci = CustomInt(int(v))
	case string:
		intValue, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*ci = CustomInt(intValue)
	default:
		return fmt.Errorf("unexpected type for CustomInt: %T", v)
	}
	return nil
}
