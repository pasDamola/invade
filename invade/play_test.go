package invade

import (
	"testing"

	"reflect"
)

func TestMoveDirections(t *testing.T) {
	value := MoveDirections("north")
	valueType := reflect.TypeOf(value)
	if k := valueType.Kind(); k != reflect.Int {
		t.Errorf("This function must return an integer type i.e. north would return 1, south would return 2")
	}
}

func TestMoveDirectionsBack(t *testing.T) {
	value := MoveDirectionsBack(1)
	valueType := reflect.TypeOf(value)
	if k := valueType.Kind(); k != reflect.String {
		t.Errorf("This function must return an string type i.e. 1 would return 'north', 2 would return 'south'")
	}
}


func TestRun(t *testing.T){
	value := CheckEnd()
	valueType := reflect.TypeOf(value)
	if k := valueType.Kind(); k != reflect.Bool {
		t.Fail()
	}
}
