package govaluator

import (
	"github.com/isomnath/govaluator/dependencies"
	"github.com/isomnath/govaluator/models"
)

type Govaluator interface {
	GetAllDataTypes() []string
	GetOperatorsByDataType(dataType string) []string
	GetComparators() []string
	GetEngines() []string
	ExecuteOne(data map[string]interface{}, rule models.Rule) (models.Result, error)
	ExecuteMany(data map[string]interface{}, rules []models.Rule) ([]models.Result, error)
}

type govaluator struct {
	utils       utilities
	operators   operationManager
	comparators comparisonManager
	engines     engineManager
}

func Initialize() Govaluator {
	dep := dependencies.Initialize()
	return &govaluator{
		utils:       dep.Utilities,
		operators:   dep.OperationManager,
		comparators: dep.ComparisonManager,
		engines:     dep.EngineManager,
	}
}

func (e *govaluator) GetAllDataTypes() []string {
	return e.operators.GetAllDataTypes()
}

func (e *govaluator) GetOperatorsByDataType(dataType string) []string {
	return e.operators.GetOperators(dataType)
}

func (e *govaluator) GetComparators() []string {
	return e.comparators.GetComparators()
}

func (e *govaluator) GetEngines() []string {
	return e.engines.GetEngines()
}

func (e *govaluator) ExecuteOne(data map[string]interface{}, rule models.Rule) (models.Result, error) {
	flattenedData := e.utils.Flatten(data)
	engine := e.engines.GetEngine(rule.Engine)
	return engine.Execute(flattenedData, rule), nil
}

func (e *govaluator) ExecuteMany(data map[string]interface{}, rules []models.Rule) ([]models.Result, error) {
	var results []models.Result
	flattenedData := e.utils.Flatten(data)
	for _, rule := range rules {
		engine := e.engines.GetEngine(rule.Engine)
		result := engine.Execute(flattenedData, rule)
		results = append(results, result)
	}

	return results, nil
}
