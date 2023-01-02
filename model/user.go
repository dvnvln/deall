package model

import (
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username  string             `bson:"username,omitempty"`
	Password  string             `bson:"password,omitempty"`
	Role      Details            `bson:"inline"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type Details struct {
	RoleName string
}

func (u *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (u *User) Bind(r *http.Request) error {
	return nil
}
