package govaluator

import (
	"testing"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/models"
)

var g Govaluator

func executeOne(g Govaluator, n int) {
	data := map[string]interface{}{
		"parent.child_1": map[string]interface{}{
			"grand_child_1": 10,
			"grand_child_2": 20,
		},
		"parent.child_2": map[string]interface{}{
			"grand_child_1": 10,
			"grand_child_2": 20,
		},
		"parent.child_3": map[string]interface{}{
			"grand_child_1": 10,
			"grand_child_2": 20,
		},
	}

	for i := 0; i < n; i++ {
		_, _ = g.ExecuteOne(data, getRule("valid_linear_rule"))
	}

}

func executeMany(g Govaluator, n int) {
	data := map[string]interface{}{
		"parent.child_1": map[string]interface{}{
			"grand_child_1": 10,
			"grand_child_2": 20,
		},
		"parent.child_2": map[string]interface{}{
			"grand_child_1": 10,
			"grand_child_2": 20,
		},
		"parent.child_3": map[string]interface{}{
			"grand_child_1": 10,
			"grand_child_2": 20,
		},
	}
	rules := []models.Rule{
		getRule("valid_linear_rule"),
		getRule("valid_linear_waterfall_rule"),
	}

	for i := 0; i < n; i++ {
		_, _ = g.ExecuteMany(data, rules)
	}

}

func getRule(ruleType string) models.Rule {
	rules := map[string]models.Rule{
		"valid_linear_rule": {
			ID:          "123",
			Name:        "Linear Rule",
			Description: "Description - Linear Rule",
			Remark:      "",
			Engine:      constants.LinearEngine,
			Linear: &models.Linear{
				Expression: "(({{.A}} && {{.B}}) || {{.C}})",
				Criteria: []models.Criterion{
					{
						ID:         "1",
						Alias:      "A",
						Comparator: constants.StaticComparator,
						Static: &models.Static{
							FieldOne: "parent.child_1.grand_child_1",
							Operator: constants.Equals,
							Value:    10,
						},
					},
					{
						ID:         "2",
						Alias:      "B",
						Comparator: constants.DynamicComparator,
						Dynamic: &models.Dynamic{
							FieldOne: "parent.child_1.grand_child_1",
							Operator: constants.LessThan,
							FieldTwo: "parent.child_1.grand_child_2",
						},
					},
					{
						ID:         "3",
						Alias:      "C",
						Comparator: constants.StaticComparator,
						Static: &models.Static{
							FieldOne: "parent.child_1.grand_child_3",
							Operator: constants.GreaterThan,
							Value:    20,
						},
					},
				},
			},
		},
		"valid_linear_waterfall_rule": {
			ID:          "456",
			Name:        "Linear Waterfall Rule",
			Description: "Description - Linear Waterfall Rule",
			Remark:      "",
			Engine:      constants.LinearWaterfallEngine,
			LinearWaterfall: &models.LinearWaterfall{
				ExitOn: constants.FirstTrue,
				Criteria: []models.Criterion{
					{
						ID:         "1",
						Alias:      "A",
						Comparator: constants.StaticComparator,
						Static: &models.Static{
							FieldOne: "parent.child_1.grand_child_1",
							Operator: constants.Between,
							Value:    []int{0, 50},
						},
						TruthyValue: 80,
						FalseyValue: 0,
					},
					{
						ID:         "2",
						Alias:      "B",
						Comparator: constants.StaticComparator,
						Static: &models.Static{
							FieldOne: "parent.child_1.grand_child_1",
							Operator: constants.Between,
							Value:    []int{50, 70},
						},
						TruthyValue: 90,
						FalseyValue: 0,
					},
					{
						ID:         "3",
						Alias:      "C",
						Comparator: constants.StaticComparator,
						Static: &models.Static{
							FieldOne: "parent.child_1.grand_child_1",
							Operator: constants.Between,
							Value:    []int{70, 100},
						},
						TruthyValue: 100,
						FalseyValue: 0,
					},
				},
			},
		},
	}
	return rules[ruleType]
}

func BenchmarkExecuteOne10(b *testing.B) {
	g = Initialize()
	executeOne(g, 10)
}

func BenchmarkExecuteOne100(b *testing.B) {
	g = Initialize()
	executeOne(g, 100)
}

func BenchmarkExecuteOne1000(b *testing.B) {
	g = Initialize()
	executeOne(g, 1000)
}

func BenchmarkExecuteMany10(b *testing.B) {
	g = Initialize()
	executeMany(g, 10)
}

func BenchmarkExecuteMany100(b *testing.B) {
	g = Initialize()
	executeMany(g, 100)
}

func BenchmarkExecuteMany1000(b *testing.B) {
	g = Initialize()
	executeMany(g, 1000)
}
