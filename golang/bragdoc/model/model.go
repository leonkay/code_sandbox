package model

import "time"

type RequestContext struct {
	Company    string
	Project    string
	Title      string
	Role       string
	DateStr    string
	Action     string
	TrackerKey string
	Args       []string
}

type ProcessContext struct {
	Company *Company
	Title   *Title
}

type Company struct {
	Id         int       `json:"id" db:"company_id"`
	Name       string    `json:"name" db:"company_name"`
	CreateDate time.Time `json:"create_date" db:"create_date"`
}

type Title struct {
	Id        int    `json:"id" db:"title_id"`
	CompanyId int    `json:"company_id" db:"company_id"`
	Title     string `json:"title" db:"title"`
	Level     string `json:"level" db:"level"`
}

type TrackerType struct {
	Id              int    `json:"id" db:"tracker_type_id"`
	Tier            int    `json:"tier" db:"display_tier"`
	ContextEligible bool   `json:"context_eligible" db:"context_eligible"`
	Key             string `json:"key" db:"activity_key"`
}

type TrackerRelation struct {
	ParentId int `json:"parent_id" db:"parent_id"`
	ChildId  int `json:"child_id" db:"child_id"`
}

type ActivityEvent struct {
	Id            int       `json:"id" db:"activity_id"`
	Time          time.Time `json:"time" db:"time"`
	TrackerTypeId int       `json:"tracker_type_id" db:"tracker_type_id"`
	Action        Action    `json:"action" db:"action_id"`
	Title         string    `json:"title" db:"title"`
}

type ActivityEventExt struct {
	Id              int    `json:"id" db:"activity_ext_id"`
	ActivityEventId int    `json:"activity_id" db:"activity_id"`
	Description     string `json:"description" db:"description"`
}
