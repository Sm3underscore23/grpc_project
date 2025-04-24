package converter

import (
	"project/internal/model"
	desc "project/pkg/user_v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserInfoPublicFromService(user *model.UserInfoPublic) *desc.UserInfoPublic {
	return &desc.UserInfoPublic{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Enum:      toProtoRole(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func ToServiceFromUserInfoPrivate(user *desc.UserInfoPrivate) *model.UserInfoPrivate {
	return &model.UserInfoPrivate{
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Role:     fromProtoRole(user.GetEnum()),
		Password: user.GetPassword(),
	}
}

func ToServiceFromUpdateUserInfo(user *desc.UpdateUserInfo) *model.UpdateUserInfo {
	updateInfo := model.UpdateUserInfo{}

	if user.Name != nil {
		updateInfo.Name = *user.Name
	}
	if user.Email != nil {
		updateInfo.Email = *user.Email
	}

	return &updateInfo
}

func toProtoRole(role string) desc.Role {
	switch role {
	case "ROLE_ADMIN":
		return desc.Role_ROLE_ADMIN
	case "ROLE_USER":
		return desc.Role_ROLE_USER
	default:
		return desc.Role_ROLE_UNSPECIFIED
	}
}

func fromProtoRole(role desc.Role) string {
	switch role {
	case desc.Role_ROLE_ADMIN:
		return "ROLE_ADMIN"
	case desc.Role_ROLE_USER:
		return "ROLE_USER"
	default:
		return "ROLE_UNSPECIFIED"
	}
}
