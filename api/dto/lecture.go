package dto

type CreateLectureDto struct {
	ID                  string `json:"id" validate:"required"`
	Name                string `json:"name" validate:"required"`
	LectureStartTime    string `json:"lecture_start_time" validate:"required,lecturetime"`
	BuildingID          int    `json:"building_id" validate:"required"`
	AttendanceValidTime int    `json:"attendance_valid_time" validate:"required"`
}

type CreateLectureStudentDto struct {
	LectureID  string `json:"lecture_id" validate:"required"`
	StudentIDs []int  `json:"student_id_list" validate:"required"`
}

type DeleteLectureStudent struct {
	LectureID string `json:"lecture_id" validate:"required"`
	StudentID int    `json:"student_id" validate:"required"`
}
