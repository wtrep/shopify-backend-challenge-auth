package auth

type UserRequest struct {
	// username of the user
	Name string `json:"name,omitempty"`
	// password of the user
	Password string `json:"password,omitempty"`
}

type UserResponse struct {
	// JWT to authenticate user
	Token string `json:"token,omitempty"`
}
