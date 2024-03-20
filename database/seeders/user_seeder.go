package seeders

import (
	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type UserSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *UserSeeder) Signature() string {
	return "UserSeeder"
}

// Run executes the seeder logic.
func (s *UserSeeder) Run() error {
	password, _ := facades.Hash().Make("123456")
	user := models.User{
		Name:     "Admin",
		Email:    "admin@example.com",
		Password: password,
	}
	return facades.Orm().Query().UpdateOrCreate(&user, models.User{Email: user.Email}, user)
}
