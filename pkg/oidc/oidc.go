package oidc

import (
	"context"
	"wae/config"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var Oauth2config *oauth2.Config
var Verifier *oidc.IDTokenVerifier

func InitOIDC(ctx context.Context) error {
	provider, err := oidc.NewProvider(ctx, config.Config.OIDCConfig.Provider)
	if err != nil {
		return err
	}
	Oauth2config = &oauth2.Config{
		ClientID:     config.Config.OIDCConfig.ClientID,
		ClientSecret: config.Config.OIDCConfig.ClientSecret,
		RedirectURL:  config.Config.OIDCConfig.RedirectURI,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	Verifier = provider.Verifier(&oidc.Config{
		ClientID: config.Config.OIDCConfig.ClientID,
	})
	return nil
}

type Claims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

const ClaimsKey = "claims"

func NewContext(ctx *gin.Context, claims *Claims) *gin.Context {
	ctx.Set(ClaimsKey, claims)
	return ctx
}

func FromContext(ctx *gin.Context) (*Claims, bool) {
	claims, ok := ctx.Get(ClaimsKey)
	return claims.(*Claims), ok
}
