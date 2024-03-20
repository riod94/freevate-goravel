package seeders

import (
	"encoding/json"
	"fmt"
	"goravel/app/models"
	"io"
	"net/http"

	"github.com/goravel/framework/facades"
)

type BankSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *BankSeeder) Signature() string {
	return "BankSeeder"
}

// Run executes the seeder logic.
func (s *BankSeeder) Run() error {
	resp, err := http.Get("https://raw.githubusercontent.com/riod94/list-bank/main/bank.json")

	if err != nil {
		fmt.Println("Error Fetching: ", err)
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error read body: ", err)
		return err
	}

	banks := []models.Bank{}

	if err := json.Unmarshal(body, &banks); err != nil {
		fmt.Println("Error Unmarshal: ", err)
		return err
	}

	for _, bank := range banks {
		bank := models.Bank{
			Name:     bank.Name,
			Code:     bank.Code,
			Status:   "active",
			Category: bank.Category,
			Link:     bank.Link,
			Contact:  bank.Contact,
		}
		facades.Orm().Query().UpdateOrCreate(&bank, models.Bank{Code: bank.Code}, bank)
	}

	return nil
}
