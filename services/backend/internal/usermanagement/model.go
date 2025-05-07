package usermanagement

import (
	"time"
)

type User struct {
	LastName              string    `json:"last_name"`
	FirstName             string    `json:"first_name"`
	DateOfBirth           time.Time `json:"date_of_birth"`
	PlaceOfBirth          string    `json:"place_of_birth"`
	MaritalStatus         string    `json:"marital_status"`
	Profession            string    `json:"profession"`
	Citizenship           string    `json:"citizenship"`
	Degree                string    `json:"degree"`
	SalaryPerMonth        float32   `json:"salary_per_month"`
	Assets                string    `json:"assets"`
	SelfEmployed          bool      `json:"self_employed"`
	TradingBudgetPerMonth float32   `json:"trading_budget_per_month"`
	HasTradingExperience  bool      `json:"has_trading_experience"`
	TradingExperience     string    `json:"trading_experience"`
	KnowledgeInTrading    string    `json:"knowledge_in_trading"`
	Motivation            string    `json:"motivation"`
	ExcludeCompanies      string    `json:"exclude_companies"`
	RiskTaking            int       `json:"risk_taking"`
	Preferences           string    `json:"preferences"`
}
