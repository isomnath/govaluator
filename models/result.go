package models

type Result struct {
	Success bool         `json:"success"`
	Rule    RuleMetadata `json:"rule"`
	Data    *ResultData  `json:"data"`
	Errors  []error      `json:"errors"`
}

type RuleMetadata struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Engine      string `json:"engine"`
}

type ResultData struct {
	Remark string      `json:"remark"`
	Value  interface{} `json:"value"`
}
