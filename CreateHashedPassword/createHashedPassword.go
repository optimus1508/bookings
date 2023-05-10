package createhashedpassword

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hashedPassword() {
	password := "password" //Put your password

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

	fmt.Println(string(hashedPassword))
}
