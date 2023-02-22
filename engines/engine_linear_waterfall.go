package engines

import (
	"github.com/isomnath/govaluator/models"
)

type linearWaterfall struct {
	utilities         utils
	strategyManger    strategyManger
	comparisonManager comparisonManager
}

func (engine *linearWaterfall) Execute(data map[string]interface{}, rule models.Rule) models.Result {
	var tempResults []models.TransientResult

	for _, criterion := range rule.LinearWaterfall.Criteria {
		comparator := engine.comparisonManager.GetComparator(criterion.Comparator)
		flag, res := comparator.ExecuteCriterion(data, criterion)
		tempResults = append(tempResults, models.TransientResult{Flag: flag, Result: res})
	}

	strategy := engine.strategyManger.GetStrategy(rule.LinearWaterfall.ExitOn)
	finalResult := strategy.GetResultValue(tempResults)
	if finalResult == nil {
		finalResult = rule.LinearWaterfall.DefaultReturnValue
	}

	remark, err := engine.utilities.ProcessRemark(data, rule, finalResult)
	if err != nil {
		return mapErrorResult(rule, []error{err})
	}

	return mapSuccessResult(rule, finalResult, remark)
}

func initializeLinearWaterfallEngine(utilities utils, strategyManger strategyManger, comparisonManager comparisonManager) Engines {
	return &linearWaterfall{
		utilities:         utilities,
		strategyManger:    strategyManger,
		comparisonManager: comparisonManager,
	}
}
