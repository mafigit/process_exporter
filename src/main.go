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

func getResponse() string {
	var processList []*process.Process

	var processListError error
	processList, processListError = process.Processes()

	if processListError != nil {
		fmt.Println(processListError)
		os.Exit(1)
	}

	var prometheusExpose strings.Builder
	for _, processElement := range processList {

		processName, error := processElement.Name()
		if error != nil {
			continue
		}

		cpuUsage, err := processElement.CPUPercent()

		if err != nil {
			continue
		}

		prometheusExpose.WriteString("process_exporter_process")
		prometheusExpose.WriteString("{")

		prometheusExpose.WriteString("name=\"")
		prometheusExpose.WriteString(processName)
		prometheusExpose.WriteString("\"")

		prometheusExpose.WriteString(",")
		prometheusExpose.WriteString("pid=\"")
		prometheusExpose.WriteString(strconv.FormatInt(int64(processElement.Pid), 10))
		prometheusExpose.WriteString("\"")

		prometheusExpose.WriteString("}")
		prometheusExpose.WriteString(" ")

		convertedCPUUsage := fmt.Sprintf("%.2f", cpuUsage)
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
