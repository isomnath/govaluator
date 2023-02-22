package engines

import (
	"github.com/isomnath/govaluator/models"
)

func mapSuccessResult(rule models.Rule, result interface{}, remark string) models.Result {
	return models.Result{
		Success: true,
		Rule:    mapMetadata(rule),
		Data: &models.ResultData{
			Remark: remark,
			Value:  result,
		},
	}
}

func mapErrorResult(rule models.Rule, errs []error) models.Result {
	return models.Result{
		Success: false,
		Rule:    mapMetadata(rule),
		Errors:  errs,
	}
}

func mapMetadata(rule models.Rule) models.RuleMetadata {
	return models.RuleMetadata{
		ID:          rule.ID,
		Name:        rule.Name,
		Description: rule.Description,
		Engine:      rule.Engine,
	}
}
