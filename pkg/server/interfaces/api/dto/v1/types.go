package v1

type LoginInput struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}
