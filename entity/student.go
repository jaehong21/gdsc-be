package entity

import "golang.org/x/crypto/bcrypt"

type Student struct {
	ID       int    `db:"id" json:"id"`
	Password string `db:"password" json:"password,omitempty"`
	Name     string `db:"name" json:"name"`
	Major    string `db:"major" json:"major"`
	Phone    string `db:"phone" json:"phone"`
	Email    string `db:"email" json:"email"`
	DeviceID string `db:"device_id" json:"device_id"`
}

func (s *Student) CompareHashAndPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(s.Password), []byte(password)); err != nil {
		return false
	} else {
		return true
	}
}
