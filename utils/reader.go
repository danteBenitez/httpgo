package readerutils

import (
	"io"
	"strings"
)

// Reads the passed io.Reader until it encounters the
// passed byte. If the read has an error, it returns it.
// It includes the matching byte.
func ReadUntilChar(r io.Reader, ch byte) (string, error) {
   buf := make([]byte, 1)
   var builder strings.Builder
   for {
      _, err := r.Read(buf)
      if err != nil { return builder.String(), err }
      if buf[0] == ch {
        builder.Write(buf)
        return builder.String(), nil
      }
      builder.Write(buf)
   }
}
