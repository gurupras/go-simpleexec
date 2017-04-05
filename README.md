# go-simpleexec
Wrapper around os/exec allowing easy piping of commands

## Install
    go get -u github.com/gurupras/go-simpleexec

## Usage
    package main

    import (
            "bytes"
            "fmt"
            "strings"

            "github.com/gurupras/simple-exec"
    )                                                               

    func main() {
            buf := bytes.NewBuffer(nil)                             

            cmd := simpleexec.ParseCmd("man df").Pipe("grep inode").Pipe("wc -l").Pipe(`sed -e 's/2/4/g'`)
            cmd.Stdout = buf                                        

            cmd.Start()                                             
            cmd.Wait()                                              

            fmt.Printf("Found %v lines containing the word inode\n", strings.TrimSpace(buf.String()))
    }

`Found 4 lines containing the word inode`
