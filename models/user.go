package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	ID          primitive.ObjectID `json:"_id"   bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Cargo       primitive.ObjectID `json:"cargo"   bson:"cargo,omitempty"`
	EstadoEval  bool               `json:"estado_eval"   bson:"estado_eval,omitempty"`
	EstadoAuto  bool               `json:"estado_auto"   bson:"estado_auto,omitempty"`
	EstadoRetro bool               `json:"estado_retro"   bson:"estado_retro,omitempty"`
	NameCargo   string             `json:"name_cargo,omitempty"`
}

type User struct {
	ID          primitive.ObjectID `json:"_id"   bson:"_id,omitempty"`
	Email       string             `json:"email"  bson:"email,omitempty" valid:"required,email"`
	Name        string             `json:"name" bson:"name,omitempty" valid:"required,minstringlength(2)"`
	Rol         string             `json:"rol"   bson:"rol,omitempty"`
	Hash        string             `json:"_hash" bson:"_hash,omitempty"`
	Password    string             `json:"password,omitempty"  bson:"password,omitempty" validate:"required,min=5"`
	Cargo       primitive.ObjectID `json:"cargo"   bson:"cargo,omitempty"`
	Team        primitive.ObjectID `json:"team,omitempty" bson:"team,omitempty"`
	EstadoEval  bool               `json:"estado_eval"   bson:"estado_eval,omitempty"`
	EstadoAuto  bool               `json:"estado_auto"   bson:"estado_auto,omitempty"`
	EstadoRetro bool               `json:"estado_retro"   bson:"estado_retro,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}
