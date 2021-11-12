package input_files

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) []string{
	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed to open because: %s ", err.Error())

	}

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()
	return text
}

func ReadLayout(path string) (int64, []int64, []int64, error) {
	getValues := readFile(path)
	floors, _ := strconv.ParseInt(getValues[0], 10, 64)
	mainCorridors, _ := strconv.ParseInt(getValues[1], 10, 64)
	subCorridors, _ := strconv.ParseInt(getValues[2], 10, 64)

	mainCorridorList := make([]int64, floors)
	subCorridorList := make([]int64, floors)
	for i := 0; i<int(floors); i++ {
		mainCorridorList[i], subCorridorList[i] = mainCorridors, subCorridors
	}
	return floors, mainCorridorList, subCorridorList, nil
}

func ReadEvents(path string) ([]int64, []string, []int64, []string, error) {
	lines := readFile(path)

	length := len(lines)
	floorList, crdTypeList, crdNumList, sensorInp := make([]int64,length),make([]string,length),make([]int64,length),make([]string,length)

	// Floor corridor_type corridor_number sensor_input
	for idx, line := range lines {
		splitValues := strings.Split(line, " ")
		floor, _ := strconv.ParseInt(splitValues[0], 10, 64)
		crdNum, _ := strconv.ParseInt(splitValues[2], 10, 64)

		floorList[idx],crdTypeList[idx],crdNumList[idx],sensorInp[idx] = floor, splitValues[1], crdNum, splitValues[3]
	}

	return floorList, crdTypeList, crdNumList, sensorInp, nil
}