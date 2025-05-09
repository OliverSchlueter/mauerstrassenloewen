package authentication

import "time"

type Token struct {
	ID        string    `json:"id" bson:"tokenId"`
	UserID    string    `json:"userId" bson:"userId"`
	Hash      string    `json:"hash" bson:"hash"`
	ExpiresAt time.Time `json:"expiresAt" bson:"expiresAt"`
}
