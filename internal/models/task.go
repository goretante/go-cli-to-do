/*
Copyright © 2025 Ante Goreta agoret00@fesb.hr
*/
package models

import (
	"time"
)

type Task struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Description string    `gorm:"not null" json:"description"`
	IsDone      bool      `gorm:"default:false" json:"isDone"`
	CreatedOn   time.Time `gorm:"autoCreateTime" json:"createdOn"`
	UpdatedOn   time.Time `gorm:"autoUpdateTime" json:"updatedOn"`
	DueDate     time.Time `gorm:"dueDate"`
}
