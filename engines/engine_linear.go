package engines

import (
	"github.com/isomnath/govaluator/models"
)

type linear struct {
	utilities         utils
	comparisonManager comparisonManager
}

func (engine *linear) Execute(data map[string]interface{}, rule models.Rule) models.Result {
	tmpResult := make(map[string]interface{})
	for _, criterion := range rule.Linear.Criteria {
		comparator := engine.comparisonManager.GetComparator(criterion.Comparator)
		_, result := comparator.ExecuteCriterion(data, criterion)
		tmpResult[criterion.Alias] = result
	}

	finalResult, err := engine.utilities.Evaluate(tmpResult, rule.Linear.Expression)
	if err != nil {
		return mapErrorResult(rule, []error{err})
	}

	remark, err := engine.utilities.ProcessRemark(data, rule, finalResult)
	if err != nil {
		return mapErrorResult(rule, []error{err})
	}

	return mapSuccessResult(rule, finalResult, remark)
}

func initializeLinearEngine(utilities utils, comparisonManager comparisonManager) Engines {
	return &linear{
		utilities:         utilities,
		comparisonManager: comparisonManager,
	}
}
