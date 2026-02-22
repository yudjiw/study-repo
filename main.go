package main

import (
	"Study/feature_postgres/simple_connection"
	"Study/feature_postgres/simple_sql"
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	if err := simple_sql.InsertRow(
		ctx,
		conn,
		"lunch",
		"eat",
		false,
		time.Now(),
	); err != nil {
		panic(err)
	}

	if err := simple_sql.UpdateRow(ctx, conn); err != nil {
		panic(err)
	}

	fmt.Println("succeed")
}
