package usermgmtuc

import (
	"context"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/go-playground/validator/v10"
)

type RegisterRequest struct {
	Username     string `validate:"required,min=3,max=15"`
	Password     string `validate:"required"`
	FirstName    string `validate:"min=1,max=30"`
	LastName     string `validate:"min=1,max=30"`
	Email        string `validate:"required,email"`
	MobileNumber string
}

type RegisterResponse struct {
	User *gocloak.User
}

type registerUseCase struct {
	identityManager identityManager
}

func NewRegisterUseCase(im identityManager) *registerUseCase {
	return &registerUseCase{
		identityManager: im,
	}
}

func (uc *registerUseCase) Register(ctx context.Context, request RegisterRequest) (*RegisterResponse, error) {

	var validate = validator.New()
	err := validate.Struct(request)
	if err != nil {
		return nil, err
	}

	var user = gocloak.User{
		Username:      gocloak.StringP(request.Username),
		FirstName:     gocloak.StringP(request.FirstName),
		LastName:      gocloak.StringP(request.LastName),
		Email:         gocloak.StringP(request.Email),
		EmailVerified: gocloak.BoolP(true),
		Enabled:       gocloak.BoolP(true),
		Attributes:    &map[string][]string{},
	}
	if strings.TrimSpace(request.MobileNumber) != "" {
		(*user.Attributes)["mobile"] = []string{request.MobileNumber}
	}

	userResponse, err := uc.identityManager.CreateUser(ctx, user, request.Password, "viewer")
	if err != nil {
		return nil, err
	}

	var response = &RegisterResponse{User: userResponse}
	return response, nil
}
