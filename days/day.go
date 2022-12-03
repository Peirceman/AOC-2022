package days

import (
	"strconv"
	"reflect"
)

type Day struct {}

var d Day = Day{}

func Run(day uint, part uint, filePath string) {
	toCall := "Day" + strconv.FormatUint(uint64(day), 10) + "Part" + strconv.FormatUint(uint64(part), 10)
	reflect.ValueOf(&d).MethodByName(toCall).Call([]reflect.Value{reflect.ValueOf(filePath)})
}
