package domain

type Note struct {
	Name      string
	Text      string
	Tag       string
	Url       string
	IsPublic  bool
	CreatedAt int
	UpdatedAt int
}
