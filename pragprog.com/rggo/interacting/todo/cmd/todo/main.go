package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"pragprog.com/rggo/interacting/todo"
)

const todoFileName = ".todo.json"

func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for the Pragmatic Bookshelf\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2020\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage information:\n")
		flag.PrintDefaults()
	}

	task := flag.String("task", "", "Task to be included in the Todo list")
	add := flag.Bool("add", false, "Add task the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:

		fmt.Print(l)

	case *complete > 0:

		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *task != "":
		l.Add(*task)

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *add:
		tasks, err := getTask(os.Stdin, flag.Args()...)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		for _, t := range tasks {
			l.Add(t)
		}

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

func getTask(r io.Reader, args ...string) ([]string, error) {

	tasks := []string{}
	// Concate os argumentos excedentes do -add
	if len(args) > 0 {
		for _, arg := range args {
			tasks = append(tasks, arg)
		}
		return tasks, nil
	}

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {

		if err := s.Err(); err != nil {
			return []string{}, err
		}

		if len(s.Text()) == 0 {
			return []string{}, fmt.Errorf("Task cannot be blank")
		}
		tasks = append(tasks, s.Text())
	}
	return tasks, nil

}
