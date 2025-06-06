package converter

import (
	"project/internal/model"
	repoModel "project/internal/repository/user/model"
)

func ToServiceFromRepoUIPb(user *repoModel.UserInfoPublic) *model.UserInfoPublic {
	return &model.UserInfoPublic{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
