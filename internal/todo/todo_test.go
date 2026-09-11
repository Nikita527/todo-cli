package todo_test

import (
	"errors"
	"testing"

	"github.com/Nikita527/todo-cli/internal/todo"
)

func TestStore(t *testing.T) {
	t.Run("Add + List", func(t *testing.T) {
		s := todo.New()
		task, err := s.Add("buy milk")
		if err != nil {
			t.Fatalf("Add: unexpected error: %v", err)
		}
		if task.ID != 1 || task.Text != "buy milk" || task.Done {
			t.Fatalf("Add returned %+v, want ID=1 Text=%q Done=false", task, "buy milk")
		}

		list := s.List()
		if len(list) != 1 {
			t.Fatalf("List len = %d, want 1", len(list))
		}
		got := list[0]
		if got.ID != 1 || got.Text != "buy milk" || got.Done {
			t.Errorf("List[0] = %+v, want ID=1 Text=%q Done=false", got, "buy milk")
		}
	})

	t.Run("Add empty text", func(t *testing.T) {
		s := todo.New()
		_, err := s.Add("")
		if !errors.Is(err, todo.ErrEmptyText) {
			t.Fatalf("Add(\"\") error = %v, want ErrEmptyText", err)
		}
		if len(s.List()) != 0 {
			t.Errorf("List after empty Add = %v, want empty", s.List())
		}
	})

	t.Run("Done existing", func(t *testing.T) {
		s := todo.New()
		task, err := s.Add("task")
		if err != nil {
			t.Fatalf("Add: %v", err)
		}
		if err := s.Done(task.ID); err != nil {
			t.Fatalf("Done: %v", err)
		}
		list := s.List()
		if len(list) != 1 || !list[0].Done {
			t.Errorf("after Done: %+v, want Done=true", list)
		}
	})

	t.Run("Done missing", func(t *testing.T) {
		s := todo.New()
		err := s.Done(42)
		if !errors.Is(err, todo.ErrNotFound) {
			t.Fatalf("Done(42) error = %v, want ErrNotFound", err)
		}
	})

	t.Run("Delete + List empty + Delete again", func(t *testing.T) {
		s := todo.New()
		task, err := s.Add("temp")
		if err != nil {
			t.Fatalf("Add: %v", err)
		}
		if err := s.Delete(task.ID); err != nil {
			t.Fatalf("Delete: %v", err)
		}
		if len(s.List()) != 0 {
			t.Errorf("List after Delete = %v, want empty", s.List())
		}
		err = s.Delete(task.ID)
		if !errors.Is(err, todo.ErrNotFound) {
			t.Fatalf("second Delete error = %v, want ErrNotFound", err)
		}
	})

	t.Run("List returns copy", func(t *testing.T) {
		s := todo.New()
		if _, err := s.Add("original"); err != nil {
			t.Fatalf("Add: %v", err)
		}

		list := s.List()
		list[0].Text = "mutated"
		list[0].Done = true

		again := s.List()
		if len(again) != 1 {
			t.Fatalf("List len = %d, want 1", len(again))
		}
		if again[0].Text != "original" || again[0].Done {
			t.Errorf("store changed after mutating List result: %+v", again[0])
		}
	})
}
