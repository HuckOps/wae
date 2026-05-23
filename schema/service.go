package schema

type GetServiceRequestParams struct {
	Page     uint   `form:"page,default=1"`
	PageSize uint   `form:"page_size,default=10"`
	Keyword  string `form:"keyword"`
}

type CreateServiceRequest struct {
	Name        string `json:"name" binding:"required"`
	Repo        string `json:"repo" binding:"required,url"`
	Domain      string `json:"domain" binding:"required"`
	Cluster     string `json:"cluster" binding:"required"`
	Description string `json:"description"`
}

type ClusterInfo struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}
