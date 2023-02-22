package models

import "fmt"

type CriteriaGroup struct {
	ID          string      `json:"id"`
	Alias       string      `json:"alias"`
	Description string      `json:"description"`
	Expression  string      `json:"expression"`
	Criteria    []Criterion `json:"criteria"`
}

func (cg *CriteriaGroup) Scan(data interface{}) error {
	err := scan(data, cg)
	if err != nil {
		return fmt.Errorf("failed to deserialize criteria group: %s", err.Error())
	}
	return nil
}
