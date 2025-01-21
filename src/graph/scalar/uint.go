package scalar

import (
	"fmt"
	"io"
	"strconv"
)

// Uint custom scalar type
type Uint uint

func (u *Uint) UnmarshalGQL(v interface{}) error {
	val, ok := v.(uint)

	if !ok {
		return fmt.Errorf("%T is not an uint", val)
	}

	*u = Uint(val)

	return nil
}

func (u Uint) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.FormatUint(uint64(u), 10))
}
