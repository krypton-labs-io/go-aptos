package common

import (
	"bytes"
	"fmt"
	"strconv"
)

// Uint64 represents a uint64 value for JSON string format.
type Uint64 uint64

func (u *Uint64) UnmarshalJSON(b []byte) error {
	b = bytes.Trim(b, "\"")
	v, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return err
	}
	*u = Uint64(v)
	return nil
}

func (u Uint64) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%d\"", u)), nil
}
