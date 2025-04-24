package user

import (
	"context"
	"project/iternal/model"
)

func (s *serv) Update(ctx context.Context, id int64, info *model.UpdateUserInfo) error {
	err := s.userRepository.Update(ctx, id, info)
	if err != nil {
		return err
	}

	return nil
}
