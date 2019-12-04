package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var instructionMatcher = regexp.MustCompile("(R|U|L|D)(\\d+)")

func main() {
	fmt.Println("Advent of Code 2019")
	problem1()
	problem2()
	problem3()
	problem4()
}

func problem4() {
	fmt.Println("Problem 4")
	pc := PasswordChecker{
		start: 172851,
		end:   675869,
	}
	fmt.Println(pc.passwordMatchCount())
	fmt.Println(pc.passwordMatchCountWithSecondCriteria())
}

type PasswordChecker struct {
	start int
	end int
}

func (pc PasswordChecker) passwordMatchCount() int {
	total := 0
	for i := pc.start; i<=pc.end; i++ {
		match, _ := pc.checkPassword(strconv.Itoa(i))
		if match {
			total += 1
		}
	}
	return total
}

func (pc PasswordChecker) passwordMatchCountWithSecondCriteria() interface{} {
	total := 0
	for i := pc.start; i<=pc.end; i++ {
		passwordString := strconv.Itoa(i)
		match, _ := pc.checkPassword(passwordString)
		if match && pc.checkPasswordSecondCriteria(passwordString){
			total += 1
		}
	}
	return total
}


func (pc PasswordChecker) checkPasswordSecondCriteria(password string) (bool) {
	digitArray := strings.Split(password, "")
	digitCountMap := make(map[string]int)
	for _, d := range digitArray {
		digitCountMap[d]++
	}
	for _, v := range digitCountMap {
		if v == 2 {
			return true
		}
	}
	return false
}

func (pc PasswordChecker) checkPassword(password string) (bool, string){
	if len(password) != 6 {
		return false, "password is 6 characters"
	}

	passwordAsInt, err := strconv.Atoi(password)
	if err != nil {
		return false, "password is not an integer" + err.Error()
	}

	if passwordAsInt < pc.start {
		return false, fmt.Sprintf("password is less than start of range %d\n", pc.start)
	}

	if passwordAsInt > pc.end {
		return false, fmt.Sprintf("password is greater than end of range %d\n", pc.start)
	}

	digits := strings.Split(password, "")

	var consecutiveDigitsMatch bool

	for i:=1; i<len(digits); i++ {
		if digits[i] == digits[i-1] {
			consecutiveDigitsMatch = true
		}

		thisDigitInteger, err := strconv.Atoi(digits[i])
		if err != nil {
			return false, fmt.Sprintf("not an integer digit: %v", err.Error())
		}
		lastDigitInteger, err := strconv.Atoi(digits[i-1])
		if err != nil {
			return false, fmt.Sprintf("not an integer digit: %v", err.Error())
		}

		if thisDigitInteger < lastDigitInteger {
			return false, fmt.Sprintf("digits decrease %d < %d", thisDigitInteger, lastDigitInteger)
		}
	}

	// if we reach this point, consecutive digits do not decrease
	return consecutiveDigitsMatch, "match"
}


func problem3() {
	fmt.Println("Problem 3")
	contents := readFileAsString("problem3_input")
	wireLines := strings.Split(contents, "\n")
	if len(wireLines) != 2 {
		log.Fatalln("input should be two lines")
	}
	wire1Coordinates := parseWireInput(wireLines[0])
	wire2Coordinates := parseWireInput(wireLines[1])
	matches := wireCoordinateMatches(wire1Coordinates, wire2Coordinates)
	fmt.Println(shortestManhattanMatch(matches))

	matchesWithStepCount := wireCoordinateMatchesAfterStepCount(wire1Coordinates, wire2Coordinates)
	fmt.Println(shortestMatchStepCount(matchesWithStepCount))
}

func problem1() {
	fmt.Println("Problem 1")
	contents := readFileAsString("problem1_input")
	inputs := strings.Split(contents, "\n")
	sum := 0
	improvedSum := 0
	for _, input := range inputs {
		sum += getFuel(input)
		improvedSum += getFuelImproved(input)
	}
	fmt.Println(sum)
	fmt.Println(improvedSum)
}

func readFileAsString(filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	contents := string(file)
	return contents
}

func getFuel(input string) int {
	if input == "" {
		return 0
	}
	in, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	in = in/3
	return int(in - 2)
}

func getFuelImproved(input string) int {
	if input == "" {
		return 0
	}

	diff := getFuel(input)
	total := 0 + diff
	for {
		diff = getFuel(strconv.Itoa(diff))
		if diff <=0 {
			break
		}
		total += diff
	}
	return total
}

func problem2() {
	fmt.Println("Problem 2")
	input := readFileAsString("problem2_input")
	intArray := createIntArray(input)
	intArray[1] = 12
	intArray[2] = 2
	code := processIntCode(intArray)
	fmt.Println(code[0])

	// b
	for i:=0; i<=99; i++ {
		for j:=0; j<=99; j++ {
			intArray := createIntArray(input)
			intArray[1] = i
			intArray[2] = j
			code := processIntCode(intArray)
			if code[0] == 19690720 {
				fmt.Println(100*i+j)
				return
			}
		}
	}
	fmt.Println("Found nothing")

}

func processIntCode(program []int) []int {
	length := len(program)
	for i:=0; i<length/4; i++ {
		opcode := program[4*i]
		in1 := program[4*i+1]
		in2 := program[4*i+2]
		out := program[4*i+3]

		if opcode == 1 {
			res := program[in1] + program[in2]
			program[out] = res
		}

		if opcode == 2 {
			res := program[in1] * program[in2]
			program[out] = res
		}

		if opcode == 99 {
			return program
		}
	}

	return program
}

func createIntArray(input string) []int {
	intArray := make([]int, len(input))
	split := strings.Split(input, ",")
	for i, s := range split{
		parseInt, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		intArray[i] = int(parseInt)
	}
	return intArray
}


func IntAbs(i int) int{
	if i < 0 {
		return -i
	}
	return i
}

func shortestManhattanMatch(matches []Coordinate) int{
	lowest := IntAbs(matches[0].x) + IntAbs(matches[0].y)

	for _, match := range matches {
		dst := IntAbs(match.x) + IntAbs(match.y)
		if dst < lowest {
			lowest = dst
		}
	}
	return lowest
}

func wireCoordinateMatches(coordinates1, coordinates2 []Coordinate ) []Coordinate{
	matches := []Coordinate{}
	coordinateHash := make(map[string]bool)
	for _, coordinate := range coordinates1 {
		key := fmt.Sprintf("%d,%d", coordinate.x, coordinate.y)
		coordinateHash[key] = true
	}

	for _, coordinate := range coordinates2 {
		key := fmt.Sprintf("%d,%d", coordinate.x, coordinate.y)
		if coordinateHash[key] {
			matches = append(matches, coordinate)
		}
	}

	return matches
}

type MatchAfterStepCount struct {
	match Coordinate
	stepCount int
}

func wireCoordinateMatchesAfterStepCount(coordinates1, coordinates2 []Coordinate ) []MatchAfterStepCount{
	matchesAfterStepCount := []MatchAfterStepCount{}

	coordinateHash := make(map[string]int)
	for i, coordinate := range coordinates1 {
		step := i+1
		key := fmt.Sprintf("%d,%d", coordinate.x, coordinate.y)
		coordinateHash[key] = step
	}

	for i, coordinate := range coordinates2 {
		step := i+1
		key := fmt.Sprintf("%d,%d", coordinate.x, coordinate.y)
		if coordinateHash[key] > 0 {
			matchesAfterStepCount = append(matchesAfterStepCount, MatchAfterStepCount{match: coordinate, stepCount: coordinateHash[key] + step})
		}
	}

	return matchesAfterStepCount
}


func parseWireInput(wireInput string) []Coordinate {
	tokens  := strings.Split(wireInput, ",")
	currentX := 0
	currentY := 0
	var coordinates []Coordinate
	for _, token := range tokens {
		submatch := instructionMatcher.FindStringSubmatch(token)
		direction := submatch[1]
		count := submatch[2]
		countI, err := strconv.Atoi(count)
		if err != nil {
			log.Fatalln(err)
		}
		for i:=0; i<countI; i++ {
			switch direction {
			case "U":
				currentY = currentY + 1
			case "D":
				currentY = currentY - 1
			case "L":
				currentX = currentX - 1
			case "R":
				currentX = currentX + 1
			default:
				log.Fatalln("Encountered unknown instructions", direction, count)
			}
			coordinates = append(coordinates, Coordinate{x: currentX, y: currentY})
		}

	}
	return coordinates
}

type Coordinate struct {
	x int
	y int
}

func shortestMatchStepCount(matches []MatchAfterStepCount) int {
	if len(matches) <= 0 {
		log.Fatalln("invalid match data")
	}
	lowest := matches[0].stepCount
	for _, match := range matches {
		if match.stepCount < lowest {
			lowest = match.stepCount
		}
	}
	return lowest
}