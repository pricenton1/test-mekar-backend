package model

type Job struct {
	JobID    string `json:"jobID"`
	JobLabel string `json:"jobLabel"`
}

type Education struct {
	EducationID    string `json:"educationID"`
	EducationLabel string `json:"educationLabel"`
}

type User struct {
	UserID        string    `json:"userID"`
	UserNIK       string    `json:"userNIK"`
	UserName      string    `json:"userName"`
	UserBirth     string    `json:"userBirth"`
	UserJob       Job       `json:"userJob"`
	UserEducation Education `json:"userEducation"`
	UserStatus    string    `json:"userStatus"`
	UserCreate    string    `json:"userCreated"`
}
