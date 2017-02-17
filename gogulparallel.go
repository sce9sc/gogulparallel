package main

import (
	"log"
	"os/exec"
	"fmt"
	"sync"
	"os"
	"gopkg.in/square/go-jose.v1/json"
)

var wg sync.WaitGroup

func test(fs string){
	log.Printf("Command start")

	cmd := exec.Command("gulp","--gulpfile","./gulpfile.js",fs)
	//cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	log.Printf("Just ran subprocess %d, exiting\n", cmd.Process.Pid)
	//err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
	wg.Done()
}

type ConfData struct {
	Tasks []string `json:"tasks"`
}

func main() {
	f,err :=os.Open("gulparallel.conf")
	if err !=nil{
		panic(err)
	}
	defer f.Close()

	var obj ConfData
	err = json.NewDecoder(f).Decode(&obj)
	if err !=nil{
		panic(err)
	}
	fmt.Println(obj.Tasks)
	tasksToRun := obj.Tasks

	for f:=range tasksToRun{
		wg.Add(1)
		fs:=tasksToRun[f]
		go test(fs)
	}

	wg.Wait()
	fmt.Println("Finished all")



}