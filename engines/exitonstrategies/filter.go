package exitonstrategies

import (
	"github.com/isomnath/govaluator/models"
)

func filter(flag bool, results []models.TransientResult) []models.TransientResult {
	var filteredResults []models.TransientResult

	for _, result := range results {
		if result.Flag == flag {
			filteredResults = append(filteredResults, result)
		}
	}

	return filteredResults
}
