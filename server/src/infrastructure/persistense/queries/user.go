package queries

import (
	"adaptivetesting/src/domain/aggregates/user"
	"adaptivetesting/src/domain/identificators"
	"adaptivetesting/src/infrastructure/persistense/models"
	"context"
	"errors"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserReader struct {
	pool *pgxpool.Pool
}

func (this *UserReader) GetByID(ctx context.Context, userID identificators.UserID) (*user.User, error) {
	var rows models.UserRows
	err := pgxscan.Get(ctx, this.pool, &rows, `
		select u.id, u.user_name, u.password_hash, u.registered_at, array_agg(r.role_name) as roles from users u
		inner join user_roles ur on ur.user_id = u.id
		inner join roles r on ur.role_id = r.role_name
		group by u.id
		where u.id = $1
		limit 1;
	`, userID.String())

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("User with id (%s) not found", userID.String())
		}
		return nil, err
	}
	return user.Fabric.Recover(rows.ID, rows.UserName, rows.RegisteredAt, rows.PasswordHash, rows.Roles)
}

func (this *UserReader) GetByUserName(ctx context.Context, username string) (*user.User, error) {
	var rows models.UserRows
	err := pgxscan.Get(ctx, this.pool, &rows, `
		select u.id, u.user_name, u.password_hash, u.registered_at, array_agg(r.role_name) as roles from users u
		inner join user_roles ur on ur.user_id = u.id
		inner join roles r on ur.role_id = r.role_name
		where u.user_name = $1
		group by u.id
		limit 1;
	`, username)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("User with username (%s) not found", username)
		}
		return nil, err
	}
	return user.Fabric.Recover(rows.ID, rows.UserName, rows.RegisteredAt, rows.PasswordHash, rows.Roles)
}

func (this *UserReader) IsUserNameExists(ctx context.Context, username string) bool {
	var exists bool
	err := this.pool.QueryRow(ctx, `select exists(select 1 from users where user_name = $1 limit 1);`, username).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func NewUserReader(pool *pgxpool.Pool) *UserReader {
	return &UserReader{pool: pool}
}
