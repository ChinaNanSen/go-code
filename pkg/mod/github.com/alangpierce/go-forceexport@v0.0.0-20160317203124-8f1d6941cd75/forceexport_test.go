package forceexport

import (
	"fmt"
	"testing"
)

func TestTimeNow(t *testing.T) {
	var timeNowFunc func() (int64, int32)
	GetFunc(&timeNowFunc, "time.now")
	sec, nsec := timeNowFunc()
	if sec == 0 || nsec == 0 {
		t.Error("Expected nonzero result from time.now().")
	}
}

// Note that we need to disable inlining here, or else the function won't be
// compiled into the binary. We also need to call it from the test so that the
// compiler doesn't remove it because it's unused.
//go:noinline
func addOne(x int) int {
	return x + 1
}

func TestAddOne(t *testing.T) {
	if addOne(3) != 4 {
		t.Error("addOne should work properly.")
	}

	var addOneFunc func(x int) int
	err := GetFunc(&addOneFunc, "github.com/alangpierce/go-forceexport.addOne")
	if err != nil {
		t.Error("Expected nil error.")
	}
	if addOneFunc(3) != 4 {
		t.Error("Expected addOneFunc to add one to 3.")
	}
}

func TestGetSelf(t *testing.T) {
	var getFunc func(interface{}, string) error
	err := GetFunc(&getFunc, "github.com/alangpierce/go-forceexport.GetFunc")
	if err != nil {
		t.Error("Error: %s", err)
	}
	// The two functions should share the same code pointer, so they should
	// have the same string representation.
	if fmt.Sprint(getFunc) != fmt.Sprint(GetFunc) {
		t.Errorf("Expected ")
	}
	// Call it again on itself!
	err = getFunc(&getFunc, "github.com/alangpierce/go-forceexport.GetFunc")
	if err != nil {
		t.Error("Error: %s", err)
	}
	if fmt.Sprint(getFunc) != fmt.Sprint(GetFunc) {
		t.Errorf("Expected ")
	}
}

func TestInvalidFunc(t *testing.T) {
	var invalidFunc func()
	err := GetFunc(&invalidFunc, "invalidpackage.invalidfunction")
	if err == nil {
		t.Error("Expected an error.")
	}
	if invalidFunc != nil {
		t.Error("Expected a nil function.")
	}
}
