package utils

import "fmt"

// NewImageName Splicing time to make unique Image name
func NewImageName(time int64) string {
	return fmt.Sprintf("%d-fusion-test.png", time)
}
