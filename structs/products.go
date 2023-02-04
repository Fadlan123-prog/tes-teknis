package structs

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	ID                  uuid.UUID
	Product_id          uuid.UUID
	Product_name        string
	Product_description string
}
