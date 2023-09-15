package models

type Mediable struct {
	MediaID      uint
	MediableType string
	MediableID   uint
	Tag          string
	Order        uint
	Media        Media
}

type Media struct {
	ID              uint
	Disk            string
	Directory       string
	UserID          uint
	Filename        string
	Extension       string
	MimeType        string
	AggregateType   string
	Size            uint
	VariantName     string
	OriginalMediaID uint
	CreatedAt       string // Cambia esto al tipo de dato correcto (por ejemplo, time.Time)
	UpdatedAt       string // Cambia esto al tipo de dato correcto (por ejemplo, time.Time)
}

type RequestTest struct {
	Key string `form:"key" binding:"required"`
}
