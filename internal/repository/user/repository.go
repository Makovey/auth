package user

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Makovey/microservice_auth/internal/adapter"
	"github.com/Makovey/microservice_auth/internal/client/db"
	"github.com/Makovey/microservice_auth/internal/model"
	"github.com/Makovey/microservice_auth/internal/repository"
	modelRepo "github.com/Makovey/microservice_auth/internal/repository/model"
	sq "github.com/Masterminds/squirrel"
)

const (
	tableName = "users"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	roleColumn      = "role"
	passwordColumn  = "password"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, user *model.User) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn, createdAtColumn).
		Values(user.Name, user.Email, user.Password, user.Role, time.Now()).
		Suffix(fmt.Sprintf("RETURNING %s", idColumn))

	query, args, err := builder.ToSql()
	if err != nil {
		log.Fatalf("failed to generate query: %v", err)
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		log.Fatalf("failed to execute query: %v", err)
	}

	fmt.Printf("Added new user with id -> %d", id)

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var user modelRepo.RepoUser
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		log.Fatalf("failed to select user: %v", err)
	}

	return adapter.ToUserFromRepo(&user), nil
}
