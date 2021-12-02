package main

import (
	"fmt"
	"strings"

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

	// rootCmd.Execute()
	var shellArgs string
	n, err := fmt.Scan(&shellArgs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
	args := strings.Fields(shellArgs)
	// fmt.Println(shellArgs)
	fmt.Println(args[0])

	// for {
	// 	fmt.Scan(&shellArgs)
	// 	// fmt.Println(shellArgs)

	// 	// args := strings.Fields(shellArgs)
	// 	// fmt.Println(args)
	// 	args := strings.Fields(shellArgs)
	// 	// fmt.Println(args)
	// 	fmt.Println("====args[0]:", args[0])

	// }
}
