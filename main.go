package main

import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
)

//docker: docker run <container_name> cmd args
//golang: go run main.go cmd args

func main() {
    switch os.Args[1] {
        case "run":
            run()
        default:
            panic("hwat?")
    }
}

func run() {
    // print what we're running to stdout
    fmt.Printf("running %v\n", os.Args[2:])

    //building the command to run it
    cmd := exec.Command(os.Args[2], os.Args[3:]...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.SysProcAttr = &syscall.SysProcAttr {
        Cloneflags: syscall.CLONE_NEWUTS,
    }

    must(cmd.Run())
}

func must(err error) {
    if err != nil {
        panic(err)
    }
}
