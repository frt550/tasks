package models

type Backup struct {
	Id        uint64 `postgres:"id"`
	Data      string `postgres:"data"`
	CreatedAt string `postgres:"created_at"`
}
