package mysql

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
    bun.BaseModel `bun:"table:users,alias:u"`
    
    ID        string    `bun:"id,pk"`
    Name      string    `bun:"name,notnull"`
    Email     string    `bun:"email,notnull"`
    CreatedAt time.Time `bun:"created_at,notnull"`
    UpdatedAt time.Time `bun:"updated_at,notnull"`
}