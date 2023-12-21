package model

// Updates to this file will require re-running `go generate`

// !! NOTE as a generated type value that will be inserted into the
// database, the order of these enums is strict, and any changes should
// be additive

//go:generate go run github.com/dmarkham/enumer -type=Action
type Action int

const (
	Join Action = iota + 1
	Leave
	Log
	Add
	Remove
	Update
  Switch
)

var ignoreActivityEvent = map[Action]bool {
  Switch: true,
}

func IsIgnoreActivityEvent(action Action) bool {
  flag := ignoreActivityEvent[action]
  return flag
}

// naming is based off the Richter Scale
//
//go:generate go run github.com/dmarkham/enumer -type=Tier
type Tier int

const (
	Great Tier = iota + 100
	Major
	Strong
	Moderate
	Light
	Slight
	Minor
	Micro
)
