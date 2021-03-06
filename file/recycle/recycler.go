package recycle

import (
	"errors"
	"io"
	"strings"
	"time"
	"unicode/utf16"

	"github.com/jtsylve/DFRWS-Go-Workshop/timestamp"
)

// ErrInvalid is returned when the data passed to a Decode function
// is not properly formatted.
var ErrInvalid = errors.New("Invalid recycle bin file.")

// Metadata gives information about a deleted file
type Metadata struct {
	Name    string    // Original File Name
	Size    int64     // Original File Size
	Deleted time.Time // The Date And Time The File Was Deleted
}

// recycler represents the on-disk structure of a Windows Vista+
// recycle bin file.
type recycler struct {
	Signature   uint64
	Size        int64
	DeletedTime timestamp.FileTime
	Name        [260]uint16
}

// DecodeI takes a Windows Vista+ $RECYCLE.BIN $I file as r and
// returns associated Metadata for the deleted file. ErrInvalid
// is returned if r is not a valid $I file.
func DecodeI(r io.Reader) (Metadata, error) {
	// TODO: Implement me!
	return Metadata{}, nil
}

// parseUTF16String parses and returns a UTF-16 NULL-Terminated
// string from b.
func parseUTF16String(b []uint16) string {
	s := string(utf16.Decode(b))

	// The following will only return the portion of the string
	// that occurs before the first NULL byte.
	return strings.SplitN(s, "\x00", 2)[0]
}
