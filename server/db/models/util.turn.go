package models

// WeekDay ...
type WeekDay string

const (
	// Monday ...
	Monday WeekDay = "Monday"
	// Thursday ...
	Thursday WeekDay = "Thursday"
	// Wednesday ...
	Wednesday WeekDay = "Wednesday"
	// Tuesday ...
	Tuesday WeekDay = "Tuesday"
	// Friday ...
	Friday WeekDay = "Friday"
	// Saturday ...
	Saturday WeekDay = "Saturday"
	// Sunday ...
	Sunday WeekDay = "Sunday"
)

// GetWeekDays ...
func GetWeekDays() []WeekDay {
	return []WeekDay{
		Monday,
		Thursday,
		Wednesday,
		Tuesday,
		Friday,
		Saturday,
		Sunday,
	}
}

// ValidateWeekDay ...
func ValidateWeekDay(o WeekDay) bool {
	l := GetWeekDays()
	for _, x := range l {
		if x == o {
			return true
		}
	}
	return false
}

// TurnNum ...
type TurnNum int

const (
	// Turn1 ...
	Turn1 TurnNum = 1
	// Turn2 ...
	Turn2 TurnNum = 2
	// Turn3 ...
	Turn3 TurnNum = 3
	// Turn4 ...
	Turn4 TurnNum = 4
	// Turn5 ...
	Turn5 TurnNum = 5
)

// GetTurnNums ...
func GetTurnNums() []TurnNum {
	return []TurnNum{
		Turn1,
		Turn2,
		Turn3,
		Turn4,
		Turn5,
	}
}

// ValidateTurnNum ...
func ValidateTurnNum(o TurnNum) bool {
	l := GetTurnNums()
	for _, x := range l {
		if x == o {
			return true
		}
	}
	return false
}

// Turn ...
type Turn struct {
	TurnWeekDay WeekDay `gorm:"type:varchar(10); not null"`
	TurnNum     TurnNum `gorm:"not null"`
}

// Valid ...
func (o *Turn) Valid() *[]ValidationError {
	es := make([]ValidationError, 0)
	validateTrue(&es, "TurnWeekDay", ValidateWeekDay(o.TurnWeekDay))
	validateTrue(&es, "TurnNum", ValidateTurnNum(o.TurnNum))
	return &es
}
