package repository

import (
	"database/sql"
	"film_library/model"
)

type Repository struct {
	Actor
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{Actor: NewActorRepository(db)}
}

type Actor interface {
	Add(actorId int, actor *model.Actor) (*model.Actor, error)
	Update(actorId int, actor *model.Actor) (*model.Actor, error)
	Delete(actorId int) (*model.Actor, error)
}
