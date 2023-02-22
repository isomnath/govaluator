package exitonstrategies

import (
	"github.com/isomnath/govaluator/models"
)

type ExitOnStrategy interface {
	GetResultValue(results []models.TransientResult) interface{}
}
