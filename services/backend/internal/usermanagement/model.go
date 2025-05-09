package usermanagement

import "time"

type User struct {
	ID       string  `json:"id" bson:"user_id"`
	Name     string  `json:"name" bson:"name"`
	Email    string  `json:"email" bson:"email"`
	Password string  `json:"password" bson:"password"`
	Profile  Profile `json:"profile" bson:"profile"`
}

type Profile struct {
	//Personal Information
	ID                   string    `json:"id" bson:"profile_id"`
	Surname              string    `json:"surname" bson:"surname"`
	Name                 string    `json:"name" bson:"name"`
	Birthdate            time.Time `json:"birthdate" bson:"birthdate"`
	Profession           string    `json:"profession" bson:"profession"`
	EmploymentStatus     string    `json:"employmentStatus" bson:"employmentStatus"`
	Country              string    `json:"country" bson:"country"`
	MartialStatus        string    `json:"martialStatus" bson:"martialStatus"`
	FinancialObligations string    `json:"financialObligations" bson:"financialObligations"`

	//Goal and Objectives
	PrimaryGoal       string `json:"primaryGoal" bson:"primaryGoal"`
	TimeHorizon       string `json:"timeHorizon" bson:"timeHorizon"`
	StartingCapital   string `json:"startingCapital" bson:"startingCapital"`
	MonthlyInvestment string `json:"monthlyInvestment" bson:"monthlyInvestment"`

	// Commitment and Experience
	PreferredMarkets     []string `json:"preferredMarkets" bson:"preferredMarkets"`
	ToolsPlatforms       string   `json:"toolsPlatforms" bson:"toolsPlatforms"`
	WeeklyTimeCommitment string   `json:"weeklyTimeCommitment" bson:"weeklyTimeCommitment"`

	TradingExperience            string `json:"tradingExperience" bson:"tradingExperience"`
	ChartAnalysisKnowledge       string `json:"chartAnalysisKnowledge" bson:"chartAnalysisKnowledge"`
	FundamentalAnalysisKnowledge string `json:"fundamentalAnalysisKnowledge" bson:"fundamentalAnalysisKnowledge"`
	Certifications               string `json:"certifications" bson:"certifications"`

	// Psychological Profile
	RiskAffinity    string `json:"riskAffinity" bson:"riskAffinity"`
	LossBehavior    string `json:"lossBehavior" bson:"lossBehavior"`
	StressTolerance string `json:"stressTolerance" bson:"stressTolerance"`

	//Ethics and Values
	EsgPreferences            string `json:"esgPreferences" bson:"esgPreferences"`
	ExclusionCriteria         string `json:"exclusionCriteria" bson:"exclusionCriteria"`
	ExcludePoliticalInvesting bool   `json:"excludePoliticalInvesting" bson:"excludePoliticalInvesting"`
	ReligiousRestrictions     string `json:"religiousRestrictions" bson:"religiousRestrictions"`

	//Setup
	TradingJournal        bool   `json:"tradingJournal" bson:"tradingJournal"`
	InformationManagement string `json:"informationManagement" bson:"informationManagement"`
	AvailableTech         string `json:"availableTech" bson:"availableTech"`
}
