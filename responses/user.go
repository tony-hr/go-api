package responses

type LoginResp struct {
	Username    string `json:"username"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Token       string `json:"token"`
}
