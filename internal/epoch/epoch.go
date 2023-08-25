package epoch

import (
	"strconv"
	"time"
)

type EpochMillisTime time.Time

func (e *EpochMillisTime) UnmarshalJSON(b []byte) error {
	timestamp, err := strconv.ParseInt(string(b), 10, 64) // Convert string to int64
	if err != nil {
		return err
	}

	seconds := timestamp / 1000
	ns := (timestamp % 1000) * int64(time.Millisecond)
	*e = EpochMillisTime(time.Unix(seconds, ns))

	return nil
}
