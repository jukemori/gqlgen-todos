package mysql

import (
	"time"

	"github.com/uptrace/bun"
)

type Todo struct {
    bun.BaseModel `bun:"table:todos,alias:t"`
    
    ID        string    `bun:"id,pk"`
    Text      string    `bun:"text,notnull"`
    Done      bool      `bun:"done,notnull"`
    UserID    string    `bun:"user_id,notnull"`
    CreatedAt time.Time `bun:"created_at,notnull"`
    UpdatedAt time.Time `bun:"updated_at,notnull"`
}

func (*Todo) ForeignKey(query *bun.CreateTableQuery) *bun.CreateTableQuery {
    return query.ForeignKey("(`user_id`) REFERENCES `users` (`id`)")
}