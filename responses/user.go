package responses

type LoginResp struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
