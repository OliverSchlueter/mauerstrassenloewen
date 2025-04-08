package natsdto

type SystemMessage string

type SimplePromptRequest struct {
	UserMsg   string        `json:"user_msg"`
	SystemMsg SystemMessage `json:"system_msg"`
}

type SimplePromptJob struct {
	JobID  string `json:"job_id"`
	Result string `json:"result"`
}
