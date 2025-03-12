package model

import "time"

type Shift struct {
	BaseModel
	DoctorID   int64         `json:"doctor_id" bun:",notnull"`
	LocationID int64         `json:"location_id" bun:",notnull"`
	ShiftDate  time.Time     `json:"shift_date" bun:",notnull"`
	StartTime  string        `json:"start_time" bun:",notnull"`
	EndTime    string        `json:"end_time" bun:",notnull"`
	Doctor     Doctor        `json:"doctor" bun:"rel:belongs-to,join:doctor_id=id"`
	Location   ShiftLocation `json:"location" bun:"rel:belongs-to,join:location_id=id"`

	tableName struct{} `bun:"shifts"`
}

type Holiday struct {
	BaseModel
	DoctorID    int64         `json:"doctor_id" bun:",notnull"`
	LocationID  int64         `json:"location_id" bun:",notnull"`
	HolidayDate time.Time     `json:"holiday_date" bun:",notnull"`
	Doctor      Doctor        `json:"-" bun:"rel:belongs-to,join:doctor_id=id"`
	Location    ShiftLocation `json:"-" bun:"rel:belongs-to,join:location_id=id"`

	tableName struct{} `bun:"holidays"`
}

type ShiftsStatus struct {
	BaseModel
	Year       int           `json:"year" bun:",notnull"`
	Month      int           `json:"month" bun:",notnull"`
	Done       bool          `json:"done" bun:",notnull,default:false"`
	LocationID int64         `json:"location_id" bun:",notnull"`
	Location   ShiftLocation `json:"-" bun:"rel:belongs-to,join:location_id=id"`

	tableName struct{} `bun:"shifts_status"`
}
