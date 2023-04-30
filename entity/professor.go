package entity

import "golang.org/x/crypto/bcrypt"

type Professor struct {
	ID       int    `db:"id" json:"id"`
	Password string `db:"password" json:"password,omitempty"`
	Name     string `db:"name" json:"name"`
	Major    string `db:"major" json:"major"`
	Phone    string `db:"phone" json:"phone"`
	Email    string `db:"email" json:"email"`
}

func (p *Professor) CompareHashAndPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password)); err != nil {
		return false
	} else {
		return true
	}
}
