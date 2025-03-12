package model

type Doctor struct {
	BaseModel
	UserID         int64  `json:"user_id" bun:",notnull"`
	Specialization string `json:"specialization"`
	Title          string `json:"title"`
	ShiftLimit     int    `json:"shift_limit"`
	User           User   `json:"-" bun:"rel:belongs-to,join:user_id=id"`

	tableName struct{} `bun:"doctors"`
}

type ShiftLocation struct {
	BaseModel
	Name        string `json:"name" bun:",notnull"`
	Description string `json:"description,omitempty"`

	tableName struct{} `bun:"shift_locations"`
}

type DoctorShiftLocation struct {
	BaseModel
	DoctorID   int64         `json:"doctor_id" bun:",notnull"`
	LocationID int64         `json:"location_id" bun:",notnull"`
	Doctor     Doctor        `json:"-" bun:"rel:belongs-to,join:doctor_id=id"`
	Location   ShiftLocation `json:"-" bun:"rel:belongs-to,join:location_id=id"`

	tableName struct{} `bun:"doctor_shift_locations"`
}
