package utilities

import (
	"fmt"

	"github.com/isomnath/govaluator/models"
)

func (utils *Utilities) ProcessRemark(data map[string]interface{}, rule models.Rule, resultValue interface{}) (string, error) {
	if rule.Remark == "" {
		return "", nil
	}

	templateMap := map[string]interface{}{
		"rule": map[string]interface{}{
			"id":          rule.ID,
			"name":        rule.Name,
			"description": rule.Description,
			"engine":      rule.Engine,
		},
		"result_value": resultValue,
		"data":         utils.UnFlatten(data),
	}

	for k, v := range utils.UnFlatten(data) {
		templateMap[k] = v
	}

	parsedExp, err := utils.parseExp(templateMap, rule.Remark)
	if err != nil {
		return "", fmt.Errorf("failed to parse expression: %s with error: %v", rule.Remark, err)
	}
	return parsedExp, nil
}
