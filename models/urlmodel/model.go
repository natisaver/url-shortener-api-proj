package urlmodel

type URL struct {
	// gorm does provide a model you can embed with default fields but im not going to use it
	// gorm.Model

	// I also already set ID as autoincrementing and thus dont have to include it under gorm metadata tag
	// ID        uint64 `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	// the json metadata tag indicates how the field should be marshalled as, when in a json file

	ID        uint64 `json:"id" gorm:"primaryKey;column:id"`
	ShortURL  string `json:"shorturl" gorm:"column:shorturl"`
	LongURL   string `json:"longurl" gorm:"column:longurl"`
	CreatedAt uint64 `json:"createdat" gorm:"column:createdat"`
	// store time as uint64 as a standard, unsigned long integer without timezone, do processing outside when retrieving
}
