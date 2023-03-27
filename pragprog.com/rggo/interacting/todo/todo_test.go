package todo_test

import (
	"testing"

	"pragprog.com/rggo/interacting/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should not be completed.")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New task should be completed.")
	}

}

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"New task 1",
		"New task 2",
		"New task 3",
	}

	for _, v := range tasks {
		l.Add(v)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead", tasks[0], l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected list length %q, got %q instead", 2, len(l))
	}
}
