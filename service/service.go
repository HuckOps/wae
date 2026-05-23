package service

import (
	"fmt"
	"wae/model"
	"wae/pkg/oidc"
	"wae/repo"
	"wae/schema"

	"github.com/gin-gonic/gin"
)

func GetService(ctx *gin.Context, params schema.GetServiceRequestParams) (int64, []model.Service, error) {
	claims, ok := oidc.FromContext(ctx)
	if !ok {
		return 0, make([]model.Service, 0), fmt.Errorf("not authenticated")
	}

	opts := []repo.Option{}
	if params.Keyword != "" {
		opts = append(opts, repo.WithWhere("name LIKE ?", "%"+params.Keyword+"%"))
	}
	opts = append(opts, repo.WithWhere("creator = ?", claims.Name))
	return repo.GetByPagination[model.Service](ctx, int(params.Page), int(params.PageSize), opts...)

}

func CreateService(ctx *gin.Context, req schema.CreateServiceRequest) error {
	claims, ok := oidc.FromContext(ctx)
	if !ok {
		return fmt.Errorf("not authenticated")
	}

	service := model.Service{
		Name:        req.Name,
		Repo:        req.Repo,
		Domain:      req.Domain,
		Cluster:     req.Cluster,
		Description: req.Description,
		Creator:     claims.Name,
		Admins:      []byte("[]"),
		Status:      "pending",
		Ref:         "",
	}

	return repo.Create(ctx, service)
}
