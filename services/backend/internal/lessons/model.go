package lessons

type Lesson struct {
	UserID string          `json:"user_id" bson:"user_id"`
	Done   map[string]bool `json:"done" bson:"done"`
}
