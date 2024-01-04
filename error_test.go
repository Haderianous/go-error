package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestErr_Error(t *testing.T) {
	fmt.Println(UnProcessable(errors.New("validation error")))
}
