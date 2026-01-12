package main

import (
	"adaptivetesting/src/application/dto/requests"
	user_usercases "adaptivetesting/src/application/usecases/user"
	"adaptivetesting/src/infrastructure/persistense/queries"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func getPool(ctx context.Context) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, "postgres://postgres:postgres@localhost:5433/testdb?sslmode=disable&pool_max_conns=10")
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func main() {
	ctx := context.Background()
	pool, err := getPool(ctx)
	if err != nil {
		panic(err)
	}
	// writer := database.NewWriter(pool)
	reader := queries.NewUserReader(pool)
	pool.Exec(ctx, "select * from users")
	start := time.Now()
	// _, err = user_usercases.FabricRegistration(writer, reader).Execute(ctx, requests.RegisterUserDTO{UserName: "user1", Password: "12345", Role: requests.Teacher})
	// if err != nil {
	// 	panic(err)
	// }
	res, err := user_usercases.FabricAuthorization(reader).Execute(ctx, &requests.AuthUserDTO{UserName: "user1", Password: "12345"})
	if err != nil {
		panic(err)
	}
	end := time.Since(start)
	jsn, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsn), end)
}
