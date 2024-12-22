package user_details_repo

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`                   // Primary Key
	UniqueID     uuid.UUID          `bson:"unique_id" json:"unique_id"`                 // Unique
	UserName     string             `bson:"user_name" json:"user_name"`                 // Unique
	Email        string             `bson:"email" json:"email"`                         // Unique
	PasswordHash string             `bson:"password_hash" json:"password_hash"`         // Password hash
	PhoneNumber  *string            `bson:"phone_number,omitempty" json:"phone_number"` // Optional, Unique
	Role         []string           `bson:"role" json:"role"`                           // Enum: User, Admin
	IsVerified   bool               `bson:"is_verified" json:"is_verified"`             // Boolean
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`               // Timestamp
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`               // Timestamp
}
