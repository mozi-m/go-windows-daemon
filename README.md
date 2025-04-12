# go-windows-daemon
Go library for creating child processes (daemons) in Windows.

This repository demonstrates how to create a detached child process in Go on Windows using 'golang.org/x/sys/windows'. This is particularly useful for creating background services or daemons on Windows.


## Why `golang.org/x/sys/windows`?

This package provides low-level access to Windows APIs such as `DETACHED_PROCESS` and `CREATE_NEW_PROCESS_GROUP`, which are necessary for creating detached child processes. There isn’t much official documentation for using these APIs to manage Windows process behavior.


## Features

- Detaches the child process from the terminal.
- Uses environment variables to distinguish between parent and child processes.

## Usage

```go

func main() {
    // init code

    if os.Getenv("GO_ENV_DAEMON") == "" {
        daemon() // spawn child and exit parent
    }

    // rest of logic code
   
} 

