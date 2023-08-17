package repository

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/plinkplenk/simple_shortner/internal/domain"
)

type userRepository struct {
	db         *pgxpool.Pool
	collection string
}

func NewUserRepository(db *pgxpool.Pool, collection string) domain.UserRepository {
	return &userRepository{
		db:         db,
		collection: collection,
	}
}

func (ur *userRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	queryBuilder := sq.Insert(ur.collection).
		Columns("email", "username", "password").
		Values(user.Email, user.Username, user.Password).
		Suffix("RETURNING *")
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := ur.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[domain.User])
}

func (ur *userRepository) getByField(ctx context.Context, field, value string) (*domain.User, error) {
	queryBuilder := sq.Select("*").From(ur.collection).Where(fmt.Sprintf("%v = ?", field), value)
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := ur.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[domain.User])
}

func (ur *userRepository) GetByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*domain.User, error) {
	queryBuilder := sq.Select("*").
		From(ur.collection).
		Where("email = ? OR username = ? ", usernameOrEmail, usernameOrEmail)
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := ur.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[domain.User])
}

func (ur *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return ur.getByField(ctx, "id", id.String())
}

func (ur *userRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	return ur.getByField(ctx, "username", username)
}

func (ur *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return ur.getByField(ctx, "email", email)
}

func (ur *userRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	panic("")
}
