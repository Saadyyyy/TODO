package respons

import "time"

type CreateTugasRespon struct {
	ID          uint `gorm:"type:auto_increment"`
	Task        string
	Level       string
	Deadline    string
	Description string
	Status      bool
	Created_at  time.Time
}
type UpdateTugasRespon struct {
	ID          uint
	Task        string
	Level       string
	Deadline    string
	Description string
	Status      bool
	Update_at   time.Time
}

type DeleteTugasRespon struct {
	ID         uint
	Task       string
	Deleted_at time.Time
}

type GetIdTugasRespon struct {
	ID          uint
	Task        string
	Level       string
	Deadline    string
	Description string
	Status      bool
}
