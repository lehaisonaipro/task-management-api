package models

type User struct {
	ID          string `json:"id" bson:"_id"`
	Username    string `json:"username" bson:"username"`
	Role        string `json:"role" bson:"role"`                             // "Employer" or "Employee"
	Password    string `json:"password,omitempty" bson:"password,omitempty"` // Store hashed password
	DisplayName string `json:"display_name" bson:"display_name"`
	Email       string `json:"email" bson:"email"`
	Dob         string `json:"dob" bson:"dob"`
}
