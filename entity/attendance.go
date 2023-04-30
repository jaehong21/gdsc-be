package entity

import "time"

type Attendance struct {
	ID              int       `db:"id" json:"id"`
	StudentID       int       `db:"student_id" json:"student_id"`
	LectureID       string    `db:"lecture_id" json:"lecture_id"`
	Status          string    `db:"status" json:"status"`
	RequestIP       string    `db:"request_ip" json:"request_ip"`
	RequestLocation string    `db:"request_location" json:"request_location"`
	Validator       string    `db:"validator" json:"validator"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
}

const (
	STATUS_OK      = "OK"
	STATUS_LATE    = "LATE"
	STATUS_ABSENT  = "ABSENT"
	STATUS_EXCUSED = "EXCUSED"
	STATUS_UNKNOWN = "UNKNOWN"
	STATUS_ETC     = "ETC"
)
