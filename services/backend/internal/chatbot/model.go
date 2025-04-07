package chatbot

type SystemMessage string

const (
	FinancialAdvisor = "You are a financial advisor. You will help the user with their financial questions and provide them with the best advice possible."
)

type SimplePromptRequest struct {
	UserMsg   string        `json:"user_msg"`
	SystemMsg SystemMessage `json:"system_msg"`
}

type SimplePromptJob struct {
	JobID  string `json:"job_id"`
	Result string `json:"result"`
}
