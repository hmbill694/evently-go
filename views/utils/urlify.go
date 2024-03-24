package utils

import (
	"fmt"

	"github.com/a-h/templ"
)

func Urlify(format string, a ...any) templ.SafeURL {
	return templ.URL(fmt.Sprintf(format, a...))
}
