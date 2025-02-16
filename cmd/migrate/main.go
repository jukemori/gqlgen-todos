package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jukemori/gqlgen-todos/infra/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {
    ctx := context.Background()
    dsn := "todo_user:todo_password@tcp(localhost:3306)/todo_db?parseTime=true"
    
    sqldb, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer sqldb.Close()

    db := bun.NewDB(sqldb, mysqldialect.New())
    defer db.Close()

    // Enable query logging
    db.AddQueryHook(bundebug.NewQueryHook(
        bundebug.WithVerbose(true),
        bundebug.FromEnv("BUNDEBUG"),
    ))

    fmt.Println("Creating tables...")

    // Create tables
    models := []interface{}{
        (*mysql.User)(nil),
        (*mysql.Todo)(nil),
    }

    for _, model := range models {
        _, err := db.NewCreateTable().
            Model(model).
            IfNotExists().
            Exec(ctx)
        
        if err != nil {
            log.Fatal(err)
        }
    }

    fmt.Println("Migration completed successfully")
}