package types

import "fmt"

type Error struct {
	IsSuccessful bool     `json:"success"`
	Message      string   `json:"message"`
	StackTrace   []string `json:"backtrace"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Danbooru request failed: %s", e.Message)
}
