package json

// IsZeroer is an interface implemented by any type which wishes to
// convey whether its current value is the zero value. This is used by
// encoding/json and encoding/xml through its use of the "omitempty" option
// in struct tags.
type IsZeroer interface {
	IsZero() bool
}
