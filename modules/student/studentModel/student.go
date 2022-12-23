package studentModel

type Student struct {
	Id       string `json:"id" bson:"id"`
	Gmail    string `json:"gmail" bson:"gmail" json:"gmail,omitempty"`
	ClassId  string `json:"class_id" bson:"class_id"`
	CourseId string `json:"course_id" bson:"course_id"`
	FullName string `json:"fullName" bson:"fullName"`
	Mark     Result `json:"mark" bson:"mark"`
}