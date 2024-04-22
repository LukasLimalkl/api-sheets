package sheets

import (
	"fmt"
	"time"
)

func formatInterface(value interface{}) string {
	if value == nil {
		return ""
	}
	switch v := value.(type) {
	case string:
		return v
	case time.Time:
		return v.Format("02/01/2006 15:04:05")
	case map[string]interface{}:
		if stringValue, ok := v["string_value"].(string); ok {
			return stringValue
		}
	default:
		return fmt.Sprintf("%v", v)
	}
	return ""
}
