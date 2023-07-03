package identity

import (
	"context"
	"fmt"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type identityManager struct {
	baseUrl             string
	realm               string
	restApiClientId     string
	restApiClientSecret string
}

func NewIdentityManager() *identityManager {
	return &identityManager{
		baseUrl:             viper.GetString("KeyCloak.BaseUrl"),
		realm:               viper.GetString("KeyCloak.Realm"),
		restApiClientId:     viper.GetString("KeyCloak.RestApi.ClientId"),
		restApiClientSecret: viper.GetString("KeyCloak.RestApi.ClientSecret"),
	}
}

func (im *identityManager) loginRestApiClient(ctx context.Context) (*gocloak.JWT, error) {
	client := gocloak.NewClient(im.baseUrl)

	token, err := client.LoginClient(ctx, im.restApiClientId, im.restApiClientSecret, im.realm)
	if err != nil {
		return nil, errors.Wrap(err, "unable to login the rest client")
	}
	return token, nil
}

func (im *identityManager) CreateUser(ctx context.Context, user gocloak.User, password string, role string) (*gocloak.User, error) {

	token, err := im.loginRestApiClient(ctx)
	if err != nil {
		return nil, err
	}

	client := gocloak.NewClient(im.baseUrl)

	userId, err := client.CreateUser(ctx, token.AccessToken, im.realm, user)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create the user")
	}

	err = client.SetPassword(ctx, token.AccessToken, userId, im.realm, password, false)
	if err != nil {
		return nil, errors.Wrap(err, "unable to set the password for the user")
	}

	var roleNameLowerCase = strings.ToLower(role)
	roleKeycloak, err := client.GetRealmRole(ctx, token.AccessToken, im.realm, roleNameLowerCase)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to get role by name: '%v'", roleNameLowerCase))
	}
	err = client.AddRealmRoleToUser(ctx, token.AccessToken, im.realm, userId, []gocloak.Role{
		*roleKeycloak,
	})
	if err != nil {
		return nil, errors.Wrap(err, "unable to add a realm role to user")
	}

	userKeycloak, err := client.GetUserByID(ctx, token.AccessToken, im.realm, userId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get recently created user")
	}

	return userKeycloak, nil
}

func (im *identityManager) RetrospectToken(ctx context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, error) {

	client := gocloak.NewClient(im.baseUrl)

	rptResult, err := client.RetrospectToken(ctx, accessToken, im.restApiClientId, im.restApiClientSecret, im.realm)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrospect token")
	}
	return rptResult, nil
}
