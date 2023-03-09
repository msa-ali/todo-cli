package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/msa-ali/todo-cli"
)

const todoFilename = ".todo.json"

func main() {
	// go run main.go -h
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed with ❤️ by Core Team.\n\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright@2023\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage Information:\n\n")
		flag.PrintDefaults()
	}
	// command line flags
	task := flag.String("task", "", "Task to be included in the ToDo List")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	l := &todo.List{}
	if err := l.Get(todoFilename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		fmt.Println(l)
		// for _, item := range *l {
		// 	if !item.Done {
		// 		fmt.Println(item.Task)
		// 	}
		// }
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		l.Add(*task)
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

}
