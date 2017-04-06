package multus

import "github.com/jinzhu/gorm"

// Handler holds app state
type Handler struct {
	DB *gorm.DB
}
