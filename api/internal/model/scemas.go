package model

type Relationship struct {
	Name  string `gorm:"primaryKey" json:"name"`
	Rel   string `json:"rel"`
	Email string `json:"email"`
}

type Camlog struct {
	ID               int64        `json:"id"`
	RelationshipName string       `json:"relname"`
	Relationship     Relationship `json:"-" gorm:"foreignKey:RelationshipName"`
	IPaddr           string       `json:"ipaddr"`
	Datetime         string       `json:"datetime"`
}
