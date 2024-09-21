package model

import "task-api/internal/constant"

type Item struct {
	ID       uint                `json:"id" gorm:"primaryKey"`
	Title    string              `json:"title"`
	Price    float64             `json:"price"`
	Quantity uint                `json:"quantity"`
	Status   constant.ItemStatus `json:"status"`
}