package models

import "time"

type StudentAuth struct {
	InstitutionalKey 	string `json:"institutional_key"`
	StudentID     	 	string `json:"student_id"`
	InstitutionID 	 	string `json:"Institutional_id"`
}
type StudentToken struct {
	Token         string
	StudentID     string
	InstitutionID string
	ExpiresAt     time.Time
}

type TeacherAuth struct {
	TeacherID     string `json:"teacher_id"`
	InstitutionID string `json:"institutional_id"`
}

type TeacherToken struct {
	Token         string
	TeacherID     string
	InstitutionID string
	ExpiresAt     time.Time
}
