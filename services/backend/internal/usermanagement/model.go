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
	ID                           string    `json:"id" bson:"profile_id"`
	Surname                      string    `json:"surname" bson:"surname"`
	Name                         string    `json:"name" bson:"name"`
	Birthdate                    time.Time `json:"birthdate" bson:"birthdate"`
	Profession                   string    `json:"profession" bson:"profession"`
	EmploymentStatus             string    `json:"employmentStatus" bson:"employmentStatus"`
	WorkHoursFlexibility         string    `json:"workHoursFlexibility" bson:"workHoursFlexibility"`
	Residence                    string    `json:"residence" bson:"residence"`
	MartialStatus                string    `json:"martialStatus" bson:"martialStatus"`
	FinancialObligations         string    `json:"financialObligations" bson:"financialObligations"`
	PrimaryGoal                  string    `json:"primaryGoal" bson:"primaryGoal"`
	TimeHorizon                  string    `json:"timeHorizon" bson:"timeHorizon"`
	StartingCapital              string    `json:"startingCapital" bson:"startingCapital"`
	MonthlyInvestment            string    `json:"monthlyInvestment" bson:"monthlyInvestment"`
	AnnualProfitTarget           string    `json:"annualProfitTarget" bson:"annualProfitTarget"`
	PreferredMarkets             string    `json:"preferredMarkets" bson:"preferredMarkets"`
	ToolsPlatforms               string    `json:"toolsPlatforms" bson:"toolsPlatforms"`
	WeeklyTimeCommitment         string    `json:"weeklyTimeCommitment" bson:"weeklyTimeCommitment"`
	TradingExperience            string    `json:"tradingExperience" bson:"tradingExperience"`
	PastPerformance              string    `json:"pastPerformance" bson:"pastPerformance"`
	ChartAnalysisKnowledge       string    `json:"chartAnalysisKnowledge" bson:"chartAnalysisKnowledge"`
	FundamentalAnalysisKnowledge string    `json:"fundamentalAnalysisKnowledge" bson:"fundamentalAnalysisKnowledge"`
	Certifications               string    `json:"certifications" bson:"certifications"`
	RiskType                     string    `json:"riskType" bson:"riskType"`
	LossBehavior                 string    `json:"lossBehavior" bson:"lossBehavior"`
	Discipline                   string    `json:"discipline" bson:"discipline"`
	TradingJournal               bool      `json:"tradingJournal" bson:"tradingJournal"`
	StressTolerance              string    `json:"stressTolerance" bson:"stressTolerance"`
	EsgPreferences               string    `json:"esgPreferences" bson:"esgPreferences"`
	ExclusionCriteria            string    `json:"exclusionCriteria" bson:"exclusionCriteria"`
	PoliticalInvesting           bool      `json:"politicalInvesting" bson:"politicalInvesting"`
	ReligiousRestrictions        string    `json:"religiousRestrictions" bson:"religiousRestrictions"`
	BrokerFeesModel              string    `json:"brokerFeesModel" bson:"brokerFeesModel"`
	TaxSituation                 string    `json:"taxSituation" bson:"taxSituation"`
	Bookkeeping                  string    `json:"bookkeeping" bson:"bookkeeping"`
	AutomationInterest           bool      `json:"automationInterest" bson:"automationInterest"`
	AvailableTech                string    `json:"availableTech" bson:"availableTech"`
}
