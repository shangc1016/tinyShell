package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"os/exec"
	"syscall"

	"log"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "tinyShell",
		Short: "a tiny shell tool",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				fmt.Println(args)
			}
			var tinyShellArgs string
			// env := os.Environ()
			for {
				fmt.Scan(&tinyShellArgs)
				// fmt.Println(tinyShellArgs)
				arg := strings.Fields(tinyShellArgs)
				for _, value := range arg {
					fmt.Println("value:", value)
				}
				// pid, _, _ := syscall.Syscall(57, 0, 0, 0)
				// if pid == 0 {
				// 	path, err := exec.LookPath(arg[0])
				// 	fmt.Println("====Path:", path)
				// 	if err != nil {
				// 		log.Fatal(err.Error())
				// 		continue
				// 	}
				// 	fmt.Println("====arg:", arg)
				// 	if err := syscall.Exec(path, arg, env); err != nil {
				// 		log.Fatal(err.Error())
				// 		continue
				// 	}

				// } else {
				// 	syscall.Syscall(61, 0, 0, 0)
				// }
			}
		},
	}
)

func main() {

	for {
		userName := os.Getenv("USER")
		hostName, err := os.Hostname()
		if err != nil{
			hostName = "NIL"
		}
		pwd, err := os.Getwd()
		if err != nil{
			pwd = "NIL"
		}
		fmt.Printf("%s:%s:%s$",userName, hostName, pwd)

		inputReader := bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadString('\n')
		if err != nil{
			log.Println(err.Error())
		}
		shellArgs := strings.Fields(input)

		if shellArgs[0] == "quit" || shellArgs[0] == "q" {
			break
		}

		commandPath, err := exec.LookPath(shellArgs[0])
		env := os.Environ()
		if err != nil{
			log.Println(err.Error())
		}
		
		pid, _, _ := syscall.Syscall(57, 0, 0, 0)
		if pid == 0 {
			syscall.Exec(commandPath, shellArgs, env)
		} else{
			syscall.Syscall(61, 0, 0, 0)
		}
	}
}
