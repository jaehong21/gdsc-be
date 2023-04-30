package dto

type CreateAttendanceDto struct {
	LectureID       string `json:"lecture_id" validate:"required"`
	RequestIP       string `json:"request_ip" validate:"required,ip"`
	RequestLocation string `json:"request_location" validate:"required"`
}

type UpdateAttendanceDto struct {
	Status string `json:"status" validate:"required,attendancestatus"`
}
