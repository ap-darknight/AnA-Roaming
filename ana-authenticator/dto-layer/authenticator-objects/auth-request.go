package authenticator_objects

type AuthRequest struct {
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	Token        string `json:"token,omitempty"`
	AuthType     string `json:"authType,omitempty"`
	IsThirdParty bool   `json:"isThirdParty,omitempty"`
}
