package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
  
	"golang.org/x/sys/windows"
)

func daemon() {
	cmd := exec.Command(os.Args[0], os.Args...)
	cmd.SysProcAttr = &windows.SysProcAttr{
		CreationFlags: windows.CREATE_NEW_PROCESS_GROUP | windows.DETACHED_PROCESS,
	}
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil

	cmd.Env = append(os.Environ(), "GO_ENV_DAEMON=true")

	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	os.Exit(0)

}

func main() {
  if os.Getenv("GO_ENV_DAEMON") == "" {
        daemon() // spawn child and exit parent
    }

  // Rest of your application logic goes here.
  // This is just an example of a long-running task.
  for {
    time.Sleep(5 * time.Second) // busy loop
  
  }
  
}
