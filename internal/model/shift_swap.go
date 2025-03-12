package model

import "time"

type ShiftSwapRequest struct {
	BaseModel
	LocationID       int64     `json:"location_id" bun:",notnull"`
	RequesterID      int64     `json:"requester_id" bun:",notnull"`
	RequestShiftDate time.Time `json:"request_shift_date" bun:",notnull"`
	RequestedShiftID int64     `json:"requested_shift_id" bun:",notnull"`
	AcceptorID       int64     `json:"acceptor_id" bun:",notnull"`
	OfferedShiftDate time.Time `json:"offered_shift_date" bun:",notnull"`
	OfferedShiftID   int64     `json:"offered_shift_id" bun:",notnull"`
	MutualAgreement  bool      `json:"mutual_agreement"`
	RequesterComment string    `json:"requester_comment"`
	AcceptorComment  string    `json:"acceptor_comment"`
	Status           string    `json:"status" bun:",notnull,default:'pending'"`

	Location       ShiftLocation `json:"-" bun:"rel:belongs-to,join:location_id=id"`
	Requester      Doctor        `json:"requester" bun:"rel:belongs-to,join:requester_id=id"`
	RequestedShift Shift         `json:"requested_shift" bun:"rel:belongs-to,join:requested_shift_id=id"`
	Acceptor       Doctor        `json:"acceptor" bun:"rel:belongs-to,join:acceptor_id=id"`
	OfferedShift   Shift         `json:"offered_shift" bun:"rel:belongs-to,join:offered_shift_id=id"`

	tableName struct{} `bun:"shift_swap_requests"`
}

type Notification struct {
	BaseModel
	Message  string `json:"message" bun:",notnull"`
	DoctorID int64  `json:"doctor_id" bun:",notnull"`
	IsRead   bool   `json:"is_read" bun:",notnull,default:false"`
	Doctor   Doctor `json:"-" bun:"rel:belongs-to,join:doctor_id=id"`

	tableName struct{} `bun:"notifications"`
}
