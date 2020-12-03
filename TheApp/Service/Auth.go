package Service

type AuthInformation struct {
	Email string
	Permission []string
	Role string
}

func StaticAuthService() []AuthInformation{
	var authlist = []AuthInformation{
		AuthInformation{
			Role:      "admin",
			Permission: []string{"GET", "POST","PUT","DELETE"},
		},
		AuthInformation{
			Role:      "user",
			Permission: []string{"GET"},
		},
	}
	return authlist


}