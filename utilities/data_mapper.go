package utilities

import (
	"github.com/nqd/flat"
)

func (utils *Utilities) Flatten(data map[string]interface{}) map[string]interface{} {
	result, _ := flat.Flatten(data, nil)
	return result
}

func (utils *Utilities) UnFlatten(data map[string]interface{}) map[string]interface{} {
	result, _ := flat.Unflatten(data, nil)
	return result
}
