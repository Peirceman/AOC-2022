package days

import (
	"bytes"
	"io"
	"math"
	"os"
	"strconv"
	"reflect"
)

type Day struct {}

var d Day = Day{}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func clamp(a, minVal, maxVal int) int {
	return max(min(a, maxVal), minVal)
}

func getFileLength(filePath string) int {
	file, _ := os.Open(filePath)
	defer file.Close()
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count

		case err != nil:
			os.Exit(1)
			return count
		}
	}
}

func Run(day uint, part uint, filePath string) {
	toCall := "Day" + strconv.FormatUint(uint64(day), 10) +
			"Part" + strconv.FormatUint(uint64(part), 10)
	reflect.ValueOf(&d).MethodByName(toCall).
			Call([]reflect.Value{reflect.ValueOf(filePath)})
}
