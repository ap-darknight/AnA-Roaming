package user_objects

// UserRequest is a struct that contains the request parameters for the User APIs.
type UserRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
}
