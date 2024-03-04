package readerutils

import ( 
   "testing"
   "strings"
)

func TestReadUntilChar(t *testing.T) {
   t.Run("Reads until it encounters the passed char and incluedes it", func(t *testing.T) {
      r := strings.NewReader("Hello, World!")
      result, err := ReadUntilChar(r, ',')
      if err != nil {
         t.Errorf("Expected no error, got %s", err)
      }
      if result != "Hello," {
         t.Errorf("Expected 'Hello', got %s", result)
      }
      buf := make([]byte, 7)
      r.Read(buf)
      if string(buf) != " World!" {
         t.Errorf("Expected ' World!', got %s", string(buf))
      }
   })

   t.Run("Includes matching char", func(t *testing.T) {
      r := strings.NewReader("Hello, World!")
      result, err := ReadUntilChar(r, '!')
      if err != nil {
         t.Errorf("Expected error, got none")
      }
      if result != "Hello, World!" {
         t.Errorf("Expected 'Hello, World', got %s", result)
      }
   })

   t.Run("Returns error if the character is not found", func(t *testing.T) {
      r := strings.NewReader("Hello, World!")
      result, err := ReadUntilChar(r, 'z')
      if err == nil {
         t.Errorf("Expected error, got none")
      }
      if result != "Hello, World!" {
         t.Errorf("Expected 'Hello, World', got %s", result)
      }
   })
}
