package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-ps"
	"github.com/struCoder/pidusage"
)

type MyProcess struct {
	name     string
	cpuUsage float64
	pid      int
}

func main() {
	var processList []ps.Process
	var processListError error

	processList, processListError = ps.Processes()
	myProcessList := make([]MyProcess, len(processList))

	if processListError != nil {
		fmt.Println(processListError)
		os.Exit(1)
	}

	for index, processElement := range processList {
		myPid := processElement.Pid()
		processName := processElement.Executable()
		sysInfo, err := pidusage.GetStat(myPid)
		if err == nil {
			myProcessList[index] = MyProcess{name: processName, cpuUsage: sysInfo.CPU, pid: myPid}
		}
	}
	fmt.Println(myProcessList[0])

	//	var processList []*process.Process
	//	var processListError error
	//
	//	processList, processListError = process.Processes()
	//
	//	if processListError != nil {
	//		fmt.Println(processListError)
	//		os.Exit(1)
	//	}
	//
	//	myProcessList := make([]MyProcess, len(processList))
	//
	//	for _, processElement := range processList {
	//		cpuPercent, error := processElement.CPUPercent()
	//		processName, error := processElement.Name()
	//		if error == nil {
	//			myProcessList = append(myProcessList,
	//				MyProcess{name: processName, cpuUsage: cpuPercent})
	//		}
	//	}
	//
	//	sort.Slice(myProcessList, func(i, j int) bool {
	//		return myProcessList[i].cpuUsage > myProcessList[j].cpuUsage
	//	})
	//
	//	fmt.Println(myProcessList)

}
