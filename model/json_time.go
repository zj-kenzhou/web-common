package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type JsonTime struct {
	time.Time
}

func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = JsonTime{Time: time.Time{}}
		return
	}
	now, err := time.Parse(`"`+"2006-01-02 15:04:05"+`"`, string(data))
	*t = JsonTime{Time: now}
	return
}

// MarshalJSON on JSONTime format Time field with Y-m-d H:i:s
func (t JsonTime) MarshalJSON() ([]byte, error) {
	res := t.ToString()
	if res == "" {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", res)), nil
}

// Value insert timestamp into mysql need this function.
func (t JsonTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan value of time.Time
func (t *JsonTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JsonTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t JsonTime) GobEncode() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t *JsonTime) GobDecode(data []byte) (err error) {
	if len(data) == 2 {
		*t = JsonTime{Time: time.Time{}}
		return
	}
	now, err := time.Parse(`"`+"2006-01-02 15:04:05"+`"`, string(data))
	*t = JsonTime{Time: now}
	return
}

func (t *JsonTime) ToString() string {
	return t.Time.Format("2006-01-02 15:04:05")
}
