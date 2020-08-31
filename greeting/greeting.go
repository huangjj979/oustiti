package greeting

import "fmt"

func Greet(somebody string) string {
	return fmt.Sprintf("Hello %s!", somebody)
}
