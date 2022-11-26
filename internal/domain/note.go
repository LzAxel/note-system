package domain

type Note struct {
	Id        int    `db:"id"`
	Name      string `json:"name" db:"name"`
	Text      string `json:"text" db:"text"`
	Tag       string `json:"tag" db:"tag"`
	Url       string `json:"url" db:"url"`
	IsPublic  bool   `json:"isPublic" db:"is_public"`
	CreatedAt int    `json:"createdAt" db:"created_at"`
	UpdatedAt int    `json:"updatedAt" db:"updated_at"`
	AccountId int    `json:"-" db:"account_id"`
}

type CreateNoteDTO struct {
	Name      string `json:"name" binding:"required"`
	Text      string `json:"text" binding:"required"`
	Tag       string `json:"tag"`
	IsPublic  bool   `json:"isPublic"`
	AccountId int    `json:"-"`
}
