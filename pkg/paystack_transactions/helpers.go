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

// Metadata is a custom type to handle JSON fields that might be returned as strings or structs.
type Metadata struct {
	CustomFields []struct {
		DisplayName  string `json:"display_name"`
		VariableName string `json:"variable_name"`
		Value        string `json:"value"`
	} `json:"custom_fields"`
	Referrer string `json:"referrer"`
}

// UnmarshalJSON implements the json.Unmarshaler interface for Metadata.
func (m *Metadata) UnmarshalJSON(data []byte) error {
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	switch v := value.(type) {
	case string:
		// Treat as empty Metadata if it is a string
		*m = Metadata{}
	case map[string]interface{}:
		// Treat as struct
		var metadataStruct struct {
			CustomFields []struct {
				DisplayName  string `json:"display_name"`
				VariableName string `json:"variable_name"`
				Value        string `json:"value"`
			} `json:"custom_fields"`
			Referrer string `json:"referrer"`
		}
		if err := json.Unmarshal(data, &metadataStruct); err != nil {
			return err
		}
		*m = Metadata{
			CustomFields: metadataStruct.CustomFields,
			Referrer:     metadataStruct.Referrer,
		}
	default:
		return fmt.Errorf("unexpected type for Metadata: %T", v)
	}
	return nil
}
