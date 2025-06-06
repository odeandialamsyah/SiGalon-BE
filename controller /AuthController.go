package controller 

// Fungsi login user
func LoginUser(c *fiber.Ctx) error {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// Cek apakah username ada
	user, err := repositories.GetUserByUsername(loginData.Username)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Verifikasi password
	if !utils.CheckPasswordHash(loginData.Password, user.Password) {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID.Hex())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate JWT"})
	}

	return c.JSON(fiber.Map{
		"token": token,
		"user": fiber.Map{
			"username": user.Username,
			"email": user.Email,
			"role": user.Role.Name,
		},
	})
}