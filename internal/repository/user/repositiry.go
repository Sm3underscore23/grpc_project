package user

import (
	"context"
	"project/internal/model"
	"project/internal/repository"
	"project/internal/repository/user/converter"
	repoModel "project/internal/repository/user/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	tableName = "users"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	passwordColumn  = "password"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.UserInfoPrivate) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(
			nameColumn,
			emailColumn,
			passwordColumn,
			roleColumn,
			createdAtColumn,
			updatedAtColumn,
		).Values(
		info.Name,
		info.Email,
		info.Password,
		info.Role,
		timestamppb.Now().AsTime(),
		timestamppb.Now().AsTime(),
	).Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var id int64

	err = r.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.UserInfoPublic, error) {
	builder := sq.Select(
		idColumn,
		nameColumn,
		emailColumn,
		roleColumn,
		createdAtColumn,
		updatedAtColumn,
	).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var user repoModel.UserInfoPublic
	err = r.db.QueryRow(ctx, query, args...).
		Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}

	return converter.ToServiceFromRepoUIPb(&user), nil
}

func (r *repo) Update(ctx context.Context, id int64, info *model.UpdateUserInfo) error {
	builder := sq.Update(tableName).
		SetMap(map[string]interface{}{
			nameColumn:  info.Name,
			emailColumn: info.Email,
		}).Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Query(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Query(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
