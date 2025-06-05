package controller 

// Fungsi login user
func LoginUser(c *fiber.Ctx) error {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

}