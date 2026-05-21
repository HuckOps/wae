package schema

type GetServiceRequestParams struct {
	Page     uint   `json:"page"`
	PageSize uint   `json:"page_size"`
	Keyword  string `json:"keyword"`
}
