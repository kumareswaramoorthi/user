package model

type Userprofile struct {
	ID          uint32 `json:"id"`
	UserName    string `json:"username"`
	Dob         string `json:"dob"`
	Age         int32  `json:"age"`
	Email       string `json:"email"`
	PhoneNumber uint32 `json:"phonenumber"`
}
