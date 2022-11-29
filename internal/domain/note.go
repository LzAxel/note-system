package domain

import "time"

type Note struct {
	Id        int       `json:"_id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Text      string    `json:"text" db:"text"`
	Tag       string    `json:"tag" db:"tag"`
	Url       string    `json:"url" db:"url"`
	IsPublic  bool      `json:"isPublic" db:"is_public"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	AccountId int       `json:"accountId" db:"account_id"`
}

type CreateNoteDTO struct {
	Name      string `json:"name" binding:"required"`
	Text      string `json:"text" binding:"required"`
	Tag       string `json:"tag"`
	IsPublic  bool   `json:"isPublic"`
	AccountId int    `json:"-"`
}

type UpdateNoteDTO struct {
	Id        int
	Name      *string `json:"name"`
	Text      *string `json:"text"`
	Tag       *string `json:"tag"`
	IsPublic  *bool   `json:"isPublic"`
	AccountId int
}

type GetDeleteNoteDTO struct {
	Id, AccountId int
}
