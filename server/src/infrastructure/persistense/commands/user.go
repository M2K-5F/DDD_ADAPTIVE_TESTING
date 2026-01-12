package commands

import (
	"adaptivetesting/src/domain/aggregates/user"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func SaveUser(ctx context.Context, txn pgx.Tx, user *user.User) error {
	model := user.ToPersistense()
	_, err := txn.Exec(ctx, `
			insert into 
			users  (id, user_name, registered_at, password_hash)
			values ($1, $2, $3, $4)
			on conflict (id)
			do update set 
				user_name = excluded.user_name,
				registered_at = excluded.registered_at,
				password_hash = excluded.password_hash;
		`,
		model.ID,
		model.UserName,
		model.RegisteredAt,
		model.PasswordHash,
	)
	if err != nil {
		return err
	}

	_, err = txn.Exec(ctx, `
        DELETE FROM user_roles 
        WHERE user_id = $1
    `, model.ID)
	if err != nil {
		return err
	}

	for _, role := range model.Roles {
		_, err = txn.Exec(ctx, `
            INSERT INTO roles (role_name)
            VALUES ($1)
            ON CONFLICT (role_name) DO NOTHING
        `, role)
		if err != nil {
			return err
		}

		_, err = txn.Exec(ctx, `
            INSERT INTO user_roles (user_id, role_id, created_at)
            VALUES ($1, $2, $3)
            ON CONFLICT (user_id, role_id) DO UPDATE SET
                created_at = excluded.created_at
        `,
			model.ID,
			role,
			time.Now(),
		)
		if err != nil {
			return err
		}
	}

	return nil
}
