package days

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type packetNode interface {
	compare(packetNode) int
	isInt() bool
	String() string
}

type packetList []packetNode

// positive if a is bigger, 0 if equal and negative if b is bigger
func (a packetList) compare(b packetNode) int {
	var bList packetList
	if b.isInt() {
		bList = []packetNode{b}
	} else {
		bList = b.(packetList)
	}

	minLength := min(len(a), len(bList))

	for i := 0; i < minLength; i++ {
		diff := a[i].compare(bList[i])
		if diff != 0 {
			return diff
		}
	}

	return len(a) - len(bList)
}

func (list packetList) isInt() bool {
	return false
}

func (list packetList) String() string {
	if len(list) < 1 {
		return "[]"
	}

	result := strings.Builder{}

	result.WriteString("[")

	for _, node := range list[:len(list)-1] {
		result.WriteString(node.String())
		result.WriteString(",")
	}

	result.WriteString(list[len(list)-1].String())

	result.WriteString("]")
	return result.String()
}

type packetInt int

func (a packetInt) compare(b packetNode) int {
	if b.isInt() {
		return int(a) - int(b.(packetInt))
	}

	return packetList([]packetNode{a}).compare(b)
}

func (int packetInt) isInt() bool {
	return true
}

func (packetInt packetInt) String() string {
	return strconv.Itoa(int(packetInt))
}

type packetPair struct {
	left, right packetList
}

func (pair packetPair) isInOrder() bool {
	return pair.left.compare(pair.right) < 1
}

func (pair packetPair) String() string {
	return pair.left.String() + "\n" + pair.right.String()
}

func parsePacketPair(scanner *bufio.Scanner) packetPair {
	result := packetPair{}

	scanner.Scan()
	result.left, _ = parseList(scanner.Text())

	scanner.Scan()
	result.right, _ = parseList(scanner.Text())

	return result
}

func parseList(line string) (packetList, string) {
	result := make([]packetNode, 0, len(line)/3)
	var newNode packetNode
	line = line[1:]

	for len(line) > 0 && line[0] != ']' {
		if line[0] == ',' {
			line = line[1:]
		}

		if line[0] == '[' {
			newNode, line = parseList(line)
		} else {
			newNode, line = parseInt(line)
		}

		result = append(result, newNode)
	}

	return result, line[1:]
}

var splitRegex = regexp.MustCompile("(],?)|(]?,)")

func parseInt(line string) (packetInt, string) {
	splitLine := splitBeforeRegex(splitRegex, line, 2)
	i, _ := strconv.Atoi(splitLine[0])
	return packetInt(i), splitLine[1]
}

func splitBeforeRegex(regex *regexp.Regexp, s string, splitCount int) []string {
	if splitCount == 0 {
		return nil
	}

	splitCount--
	start := 0
	matches := regex.FindAllStringIndex(s, splitCount)
	result := make([]string, len(matches), len(matches)+1)
	for i, match := range matches {
		begin := match[0]
		result[i] = s[start:begin]
		start = begin
	}

	return append(result, s[start:])
}

type packetLists []packetList

func (list packetLists) Len() int {
	return len(list)
}

func (list packetLists) Less(i, j int) bool {
	return list[i].compare(list[j]) < 0
}

func (list packetLists) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (d Day) Day13Part1(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	packetPairs := make([]packetPair, (getFileLength(filePath)+1)/3)

	for i := range packetPairs {
		packetPairs[i] = parsePacketPair(scanner)
		if !scanner.Scan() {
			break
		}
	}

	sumOfIndecies := 0
	for i, pair := range packetPairs {
		if pair.isInOrder() {
			sumOfIndecies += i + 1
		}
	}

	fmt.Println("sum of indecies of ordered pairs:", sumOfIndecies)
}

func (d Day) Day13Part2(filePath string) {
	listCompareFunc := func(a, b packetList) int { return a.compare(b) }
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lists := make([]packetList, 0, (getFileLength(filePath)+1)/3)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		nextList, _ := parseList(scanner.Text())
		lists, _ = insertSorted(lists, nextList, listCompareFunc)
	}

	lists, index1 := insertSorted(lists,
		packetList{packetList{packetInt(2)}}, listCompareFunc)

	lists, index2 := insertSorted(lists,
		packetList{packetList{packetInt(6)}}, listCompareFunc)

	fmt.Println("product of indexes:", (index1+1)*(index2+1))
}
