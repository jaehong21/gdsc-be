package entity

type Lecture struct {
	ID          string `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	ProfessorID int    `db:"professor_id" json:"professor_id"`
	// ex. lecture_start_time "09:30"
	LectureStartTime    string `db:"lecture_start_time" json:"lecture_start_time"`
	BuildingID          int    `db:"building_id" json:"building_id"`
	AttendanceValidTime int    `db:"attendance_valid_time" json:"attendance_valid_time"`
}

type LectureStudent struct {
	StudentID int    `db:"student_id" json:"student_id"`
	LectureID string `db:"lecture_id" json:"lecture_id"`
}
