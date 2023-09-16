package model

import "time"

type TodoList struct {
	ID           int       `json:"id" db:"id"`
	Title        string    `json:"title" db:"title"`
	Descriptions string    `json:"descriptions" db:"descriptions"`
	CreateAt     time.Time `json:"create_at" db:"create_at"`
	UpdateAt     time.Time `json:"update_at" db:"update_at"`
	CreateTime   string    `json:"create_time"`
	UpdateTime   string    `json:"update_time"`
}

func (tl *TodoList) SetTimes() {
	format := "2006-01-02"

	tl.CreateTime = tl.CreateAt.Format(format)
	tl.UpdateTime = tl.UpdateAt.Format(format)
}
