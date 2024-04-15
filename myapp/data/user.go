package data

import (
	"time"

	up "github.com/upper/db/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `db:"id,omitempty"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	Active    int       `db:"user_active"`
	Password  string    `db:"password"`
	CreateAt  time.Time `db:"created_at"`
	UpdateAt  time.Time `db:"updated_at"`
	Token     Token     `db:"-"`
}

// Table returns the table name associated with this model in the database
func (u *User) Table() string {
	return "users"
}

// GetAll returns a slice of all users
func (u *User) GetAll(condition up.Cond) ([]*User, error) {
	collection := upper.Collection(u.Table())

	var all []*User

	res := collection.Find(condition)
	err := res.All(&all)
	if err != nil {
		return nil, err
	}

	return all, nil
}

// GetByEmail gets one user, by email
func (u *User) GetByEmail(email string) (*User, error) {
	var theUser User

	collection := upper.Collection(u.Table())
	res := collection.Find(up.Cond{"email=": email})
	err := res.One(&theUser)
	if err != nil {
		return nil, err
	}

	var token Token
	collection = upper.Collection(token.Table())
	res = collection.Find(up.Cond{"user_id=": theUser.ID, "expiry <": time.Now()}).OrderBy("created_at desc")
	err = res.One(&token)
	if err != nil {
		if err != up.ErrNilRecord && err != up.ErrNoMoreRows {
			return nil, err
		}
	}

	theUser.Token = token

	return &theUser, nil
}

// Get gets one user by id
func (u *User) Get(id int) (*User, error) {
	var theUser User
	collection := upper.Collection(u.Table())
	res := collection.Find(up.Cond{"id =": id})

	err := res.One(&theUser)
	if err != nil {
		return nil, err
	}

	var token Token
	collection = upper.Collection(token.Table())
	res = collection.Find(up.Cond{"user_id =": theUser.ID, "expiry <": time.Now()}).OrderBy("created_at desc")
	err = res.One(&token)
	if err != nil {
		if err != up.ErrNilRecord && err != up.ErrNoMoreRows {
			return nil, err
		}
	}

	theUser.Token = token
	return &theUser, nil
}

// Update updates a user record in the database
func (u *User) Update(theUser User) error {
	theUser.UpdateAt = time.Now()
	collection := upper.Collection(u.Table())
	res := collection.Find(theUser.ID)
	err := res.Update(&theUser)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a user by id
func (u *User) Delete(id int) error {
	collection := upper.Collection(u.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a new user, and returns the newly inserted id
func (u *User) Insert(theUser User) (int, error) {
	newHash, err := bcrypt.GenerateFromPassword([]byte(theUser.Password), 12)
	if err != nil {
		return 0, err
	}

	theUser.CreateAt = time.Now()
	theUser.UpdateAt = time.Now()
	theUser.Password = string(newHash)

	collection := upper.Collection(u.Table())
	res, err := collection.Insert(theUser)
	if err != nil {
		return 0, err
	}

	id := getInsertID(res.ID)

	return id, nil

}

// ResetPassword resets a users's password, by id, using supplied password
func (u *User) ResetPassword(id int, password string) error{
	newHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil{
		return err
	}

	theUser , err:= u.Get(id)
	if err != nil{
		return err
	}

	u.Password = string(newHash)

	
	err = theUser.Update(*u)
	if err != nil{
		return err
	}

	return nil

}
