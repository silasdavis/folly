package folly

import (
	"testing"
	"reflect"
	"fmt"
)

type folly struct {
	Turrets int
}

func TestTypeName(t *testing.T) {
	fmt.Println(reflect.TypeOf(folly{}).Name())
}