package repositories

import "github.com/ara-ta3/TodoAppWithGolang/models"

type TodoRepository interface {
	FindAll() (map[int]*models.Todo, error)
	FindTodo(id int) (*models.Todo, error)
	PutTodo(t *models.Todo) error
	RemoveTodo(id int) error
}

type TodoRepositoryOnMemory struct {
	Data map[int]*models.Todo
}

func (r TodoRepositoryOnMemory) FindAll() (map[int]*models.Todo, error) {
	return r.Data, nil
}

func (r TodoRepositoryOnMemory) FindTodo(id int) (*models.Todo, error) {
	t := r.Data[id]
	return t, nil
}

func (r TodoRepositoryOnMemory) PutTodo(t *models.Todo) error {
	if t.ID == 0 {
		t.ID = r.newID()
	}
	r.Data[t.ID] = t
	return nil
}

func (r TodoRepositoryOnMemory) RemoveTodo(id int) error {
	delete(r.Data, id)
	return nil
}

func (r TodoRepositoryOnMemory) newID() int {
	if len(r.Data) <= 0 {
		return 1
	}
	keys := make([]int, 0, len(r.Data))
	for k := range r.Data {
		keys = append(keys, k)
	}
	return max(keys) + 1
}

func max(a []int) int {
	max := a[0]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}
