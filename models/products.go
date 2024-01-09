package models

import "time"

type ProductCategory struct {
	Id        uint64
	Name      string
	CreatedAt time.Time
}
