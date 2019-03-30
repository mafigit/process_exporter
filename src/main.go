package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/process"
)

type MyProcess struct {
	name     string
	cpuUsage float64
	pid      int32
}

func truncateString(str string, num int) string {
	bnoden := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		bnoden = str[0:num] + "..."
	}
	return bnoden
}

func getResponse() string {
	var processList []*process.Process

	var processListError error
	processList, processListError = process.Processes()
	myProcessList := make([]MyProcess, len(processList))
	if processListError != nil {
		fmt.Println(processListError)
		os.Exit(1)
	}
	for index, processElement := range processList {

		processName, error := processElement.Name()
		if error != nil {
			continue
		}

		cpuUsage, err := processElement.CPUPercent()

		if err == nil {
			myProcessList[index] = MyProcess{name: processName, cpuUsage: cpuUsage, pid: processElement.Pid}
		}
	}

	var prometheusExpose strings.Builder

	for _, processElement := range myProcessList {

		prometheusExpose.WriteString("process_collector_process")
		prometheusExpose.WriteString("{")

		prometheusExpose.WriteString("name=\"")
		prometheusExpose.WriteString(processElement.name)
		prometheusExpose.WriteString("\"")

		prometheusExpose.WriteString(",")
		prometheusExpose.WriteString("pid=\"")
		prometheusExpose.WriteString(strconv.FormatInt(int64(processElement.pid), 10))
		prometheusExpose.WriteString("\"")

		prometheusExpose.WriteString("}")
		prometheusExpose.WriteString(" ")

		convertedCPUUsage := fmt.Sprintf("%.2f", processElement.cpuUsage)
		prometheusExpose.WriteString(convertedCPUUsage)
		prometheusExpose.WriteString("\n")
	}

	return prometheusExpose.String()
}

func main() {
	portArgumentPtr := flag.String("port", "9020", "Port number")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := getResponse()
		fmt.Fprintf(w, response)
	})

	var listentingPort strings.Builder

	listentingPort.WriteString(":")
	listentingPort.WriteString(*portArgumentPtr)

	fmt.Println(listentingPort.String())

	http.ListenAndServe(listentingPort.String(), nil)
}
