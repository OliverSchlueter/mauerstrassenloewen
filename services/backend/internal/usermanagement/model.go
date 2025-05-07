package usermanagement

import (
	"time"
)

type User struct {
	ID                    string    `json:"id" bson:"user_id"`
	Email                 string    `json:"e_mail" bson:"e_mail"`
	Password              string    `json:"password" bson:"password"`
	LastName              string    `json:"last_name" bson:"last_name"`
	FirstName             string    `json:"first_name" bson:"first_name"`
	DateOfBirth           time.Time `json:"date_of_birth" bson:"date_of_birth"`
	PlaceOfBirth          string    `json:"place_of_birth" bson:"place_of_birth"`
	MaritalStatus         string    `json:"marital_status" bson:"marital_status"`
	Profession            string    `json:"profession" bson:"profession"`
	Citizenship           string    `json:"citizenship" bson:"citizenship"`
	Degree                string    `json:"degree" bson:"degree"`
	SalaryPerMonth        float32   `json:"salary_per_month" bson:"salary_per_month"`
	Assets                string    `json:"assets" bson:"assets"`
	SelfEmployed          bool      `json:"self_employed" bson:"self_employed"`
	TradingBudgetPerMonth float32   `json:"trading_budget_per_month" bson:"trading_budget_per_month"`
	HasTradingExperience  bool      `json:"has_trading_experience" bson:"has_trading_experience"`
	TradingExperience     string    `json:"trading_experience" bson:"trading_experience"`
	KnowledgeInTrading    string    `json:"knowledge_in_trading" bson:"knowledge_in_trading"`
	Motivation            string    `json:"motivation" bson:"motivation"`
	ExcludeCompanies      string    `json:"exclude_companies" bson:"exclude_companies"`
	RiskTaking            int       `json:"risk_taking" bson:"risk_taking"`
	Preferences           string    `json:"preferences" bson:"preferences"`
}
