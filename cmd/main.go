package main

import (
	"coffeebeans-people-backend/config"
	"context"
	"fmt"
)

func main()  {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	appConfig := config.GetAppConfiguration()

	fmt.Println("Application loaded successfully on port : ", appConfig.PORT)
}
