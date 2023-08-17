package repository

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/plinkplenk/simple_shortner/internal/domain"
)

type urlRepository struct {
	db         *pgxpool.Pool
	collection string
}

func NewURLRepository(db *pgxpool.Pool, collection string) domain.URLRepository {
	return &urlRepository{
		db:         db,
		collection: collection,
	}
}

func (ur *urlRepository) Create(ctx context.Context, url *domain.URL) (*domain.URL, error) {
	urlQuery := sq.Insert("urls").Columns("id", "original", "expire").Values(url.ID, url.Original, url.Expire)
	if url.UserID != nil {
		urlQuery = urlQuery.Columns("user_id").Values(*url.UserID)
	}
	urlQuery = urlQuery.Suffix("RETURNING *")
	query, args, err := urlQuery.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := ur.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[domain.URL])
}

func (ur *urlRepository) GetByID(ctx context.Context, id string) (*domain.URL, error) {
	urlQuery := sq.Select("id", "original", "user_id").From(ur.collection).Where("id = ?", id)
	query, args, err := urlQuery.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := ur.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[domain.URL])
}

func (ur *urlRepository) GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.URL, error) {
	urlQuery := sq.Select("id", "original", "user_id").From(ur.collection).Where("user_id = ?", userID)
	query, args, err := urlQuery.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := ur.db.Query(ctx, query, args...)
	return pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[domain.URL])
}
