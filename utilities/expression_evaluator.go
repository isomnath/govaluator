package utilities

import (
	"fmt"
	"log"

	"github.com/Knetic/govaluate"

	"github.com/isomnath/govaluator/constants"
)

func (utils *Utilities) Evaluate(tmpResult map[string]interface{}, expression string) (interface{}, error) {
	if expression == "" {
		log.Printf("expression string is empty")
		return nil, fmt.Errorf(constants.ErrorEmptyExpression)
	}
	parsedExp, err := utils.parseExp(tmpResult, expression)
	if err != nil {
		log.Printf("failed to parse expression: %s with error: %v", expression, err)
		return nil, fmt.Errorf(constants.ErrorInvalidExpression)
	}
	expr, err := govaluate.NewEvaluableExpression(parsedExp)
	if err != nil {
		log.Printf("failed to initialize expression for evaluation with error: %v", err)
		return nil, fmt.Errorf(constants.ErrorInitializingExpressionParser)
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		log.Printf("failed to evaluate expression with error: %v", err)
		return nil, fmt.Errorf(constants.ErrorEvaluatingExpression)
	}
	return result, nil
}
