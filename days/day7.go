package days

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"regexp"
	"math"
)

type filesystemObject interface {
	isDirectory() bool
	getParent() *directory
	getSize() int
	String() string
}

/////////////////
// file struct //
/////////////////

type file struct {
	parent *directory
	name string
	size int
}

func (this *file) getParent() *directory {
	return this.parent
}

func (this *file) getSize() int {
	return this.size
}

func (this *file) isDirectory() bool {
	return false
}

func (this *file) String() string {
	return fmt.Sprintf("- %s (file, %d)", this.name, this.size)
}

//////////////////////
// directory struct //
//////////////////////

type directory struct {
	children map[string]filesystemObject
	parent *directory
	name string
}

func (this *directory) getChildren() map[string]filesystemObject {
	return this.children
}

func (this *directory) isDirectory() bool {
	return true
}

func (this *directory) getParent() *directory {
	return this.parent
}

func (this *directory) getSize() int {
	sum := 0
	for _, child := range this.children {
		sum += child.getSize()
	}

	return sum
}

func (this *directory) String() string{
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("- %s (dir)\n", this.name))
	children := make([]string, len(this.children))
	for _, child := range this.children {
		children = append(children, "  " + strings.Replace(child.String(), "\n", "\n  ", -1)+"\n")
	}


	sort.Strings(children)
	for i, child := range children {
		if i >= len(children) - 1 {
			builder.WriteString(child[0:len(child)-1])
			break
		}

		builder.WriteString(child)
		// builder.WriteString("\n")
	}

	return builder.String()
}

func (this *directory) sumOfDirSizesLessThen100000() int {
	const MAX_SIZE = 100000
	sum := 0
	thisSize := this.getSize()
	if thisSize < MAX_SIZE {
		sum += thisSize
	}

	for _, child := range this.children {
		if child.isDirectory() {
			sum += (child).(*directory).sumOfDirSizesLessThen100000()
		}
	}

	return sum
}

func (this *directory) smalestDirMoreThen(MIN_SIZE int) int {
	lowestSize := math.MaxInt
	thisSize := this.getSize()
	if MIN_SIZE < thisSize && thisSize < lowestSize {
		lowestSize = thisSize
	}

	for _, child := range this.children {
		if (!child.isDirectory()) {
			continue
		}

		childSmallestDir := (child).(*directory).smalestDirMoreThen(MIN_SIZE)
		if MIN_SIZE < childSmallestDir && childSmallestDir < lowestSize {
			lowestSize = childSmallestDir
		}
	}

	return lowestSize
}

///////////////
// main code //
///////////////

func (day *Day) Day7Part1(filepath string) {
	inputFile,_ := os.Open(filepath)
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(splitByCommand)

	root := &directory {
		make(map[string]filesystemObject),
		nil,
		"/",
	}

	var currentDir *directory
	whitespaceSplitter := regexp.MustCompile("\\s+")
	for scanner.Scan() {
		line := whitespaceSplitter.Split(scanner.Text(), -1)
		if line[0][0] != '$' {
			fmt.Fprintln(os.Stderr, "unreachable")
			os.Exit(1)
		}

		if line[1] == "cd" {
			currentDir = parseCd(line[2:], currentDir, root)
			continue;
		}

		parseLs(line[2:], currentDir)
	}

	fmt.Println(root.sumOfDirSizesLessThen100000())
}


func (d *Day) Day7Part2(filepath string) {
	inputFile,_ := os.Open(filepath)
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(splitByCommand)

	root := &directory {
		make(map[string]filesystemObject),
		nil,
		"/",
	}

	var currentDir *directory
	whitespaceSplitter := regexp.MustCompile("\\s+")
	for scanner.Scan() {
		line := whitespaceSplitter.Split(scanner.Text(), -1)
		if line[0][0] != '$' {
			fmt.Fprintln(os.Stderr, "unreachable")
			os.Exit(1)
		}

		if line[1] == "cd" {
			currentDir = parseCd(line[2:], currentDir, root)
			continue;
		}

		parseLs(line[2:], currentDir)
	}

	const TOTAL_SPACE = 70000000
	const FREE_SPACE_NEEDED = 30000000
	totalFreeSpace := TOTAL_SPACE - root.getSize()
	fmt.Println(root.smalestDirMoreThen(FREE_SPACE_NEEDED - totalFreeSpace))
}

func parseLs(line []string, currentDir *directory) {
	for i := 0; i+1 < len(line); i += 2 {
		if line[i] == "dir" {
			newDir := &directory {
				make(map[string]filesystemObject),
				currentDir,
				line[i+1],
			}

			currentDir.children[line[i+1]] = newDir
			continue
		}

		fileSize, _ := strconv.Atoi(line[i])
		currentDir.children[line[i+1]] = &file {
			currentDir,
			line[i+1],
			fileSize,
		}
	}
}

// the line, starting after the cd token
func parseCd(line []string, currentDir, root *directory) *directory {
	if line[0][0] == '/' {
		currentDir = root
	}

	for _, directoryName := range strings.Split(line[0], "/") {
		if directoryName == "" {
			continue
		}

		if directoryName == ".." {
			currentDir = currentDir.getParent()
			continue
		}

		currentDir = (currentDir.children[directoryName]).(*directory)
	}

	return currentDir
}

func splitByCommand(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) == 0 && atEOF {
		return 0, nil, nil
	}

	// "ab\n$cd"
	//  01 2345"
	if index := strings.Index(string(data), "\n$"); index >= 0 {
		return index + 1, data[0:index], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}
