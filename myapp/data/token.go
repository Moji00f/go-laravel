package data

import "time"

type Token struct {
	ID        int       `db:"id" json:"id"`
	User_ID   int       `db:"user_id" json:"user_id"`
	FirstName string    `db:"first_name" json:"first_name"`
	Email     string    `db:"email" json:"email"`
	PlainText string    `db:"-" json:"token"`
	Hash      []byte    `db:"token_hash" json:"-"`
	CreateAt  time.Time `db:"created_at" json:"created_at"`
	UpdateAt  time.Time `db:"updated_at" son:"updated_at"`
	Expires   time.Time `db:"expires" json:"expires"`
}

func (t *Token) Table() string {
	return "tokens"
}
