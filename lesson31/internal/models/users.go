package models

type User struct {
	ID      string  `json:"id,omitempty"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Friends []*User `json:"friends"`
}

type UserInRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type FriendRequest struct {
	SourceID string `bson:"_id" json:"source_id,omitempty"`
	TargetID string `bson:"_id" json:"target_id,omitempty"`
}
