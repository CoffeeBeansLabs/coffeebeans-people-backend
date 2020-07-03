package coffeebeans

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	EmployeeId        int64              `json:"employee_id,omitempty" bson:"employee_id,omitempty"`
	Email             int64              `json:"number,omitempty" bson:"number,omitempty"`
	Gender            string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Role              string             `json:"role,omitempty" bson:"role,omitempty"`
	Designation       string             `json:"designation,omitempty" bson:"designation,omitempty"`
	DOJ               string             `json:"doj,omitempty" bson:"doj,omitempty"`
	YearsOfExp        string             `json:"experience,omitempty" bson:"experience,omitempty"`
	Id                primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	AddtitionalDetail AdditionalDetail   `json:"additiona_detail,omitempty" bson:"additional_detail,omitempty"`
}

type AdditionalDetail struct {
	Graduation      EducationalDetails `json:"educational_details,omitempty" bson:"educational_details,omitempty"`
	PostGradutation EducationalDetails `json:"educational_details,omitempty" bson:"educational_details,omitempty"`
}

type EducationalDetails struct {
	CollegeName string `json:"college_name,omitempty" bson:"college_name,omitempty"`
	CGPA        int    `json:"cgpa,omitempty" bson:"cgpa,omitempty"`
	PassoutYear int    `json:"passout_year,omitempty" bson:"passout_year,omitempty"`
}
