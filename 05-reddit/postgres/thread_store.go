package postgres

import "github.com/jmoiron/sqlx"

type ThreadStore struct {
	*sqlx.DB
}
