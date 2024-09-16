
package handler

import (
	"fmt"
	"github.com/appremon/kayveedb-server/core"
)

func ProcessCommand(command string) string {
	switch {
	case command == "read":
		value, err := core.ReadKey("some_key")
		if err != nil {
			return fmt.Sprintf("Error reading key: %v", err)
		}
		return value

	case command == "insert":
		err := core.InsertKey("some_key", "some_value")
		if err != nil {
			return fmt.Sprintf("Error inserting key: %v", err)
		}
		return "Insert successful"

	default:
		return "Unknown command"
	}
}
