package main

import (
	"adaptivetesting/src/application/dto/requests"
	user_usercases "adaptivetesting/src/application/usecases/user"
	"adaptivetesting/src/infrastructure/persistense/database"
	"adaptivetesting/src/infrastructure/persistense/readers"
	"adaptivetesting/src/infrastructure/persistense/writers"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	pool, err := database.GetConnectionPoolFromEnv(ctx)
	if err != nil {
		panic(err)
	}

	_ = writers.NewWriter(pool)
	reader := readers.NewUserReader(pool)

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
