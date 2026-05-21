package service

import (
	"context"
	"wae/model"
	"wae/pkg/oidc"
	"wae/repo"
	"wae/schema"
)

func GetService(ctx context.Context, params schema.GetServiceRequestParams) (int64, []model.Service, error) {
	_, ok := oidc.FromContext(ctx)
	if !ok {
		return 0, nil, nil
	}

	opts := []repo.Option{}
	if params.Keyword != "" {
		opts = append(opts, repo.WithWhere("name LIKE ?", "%"+params.Keyword+"%"))
	}
	total, items, err := repo.GetByPagination[model.Service](ctx, int(params.Page), int(params.PageSize), opts...)
	if err != nil {
		return 0, nil, err
	}
	return total, items, nil
}
