package runtime

import (
	"fmt"
	"time"
)

type Cache[T any] interface {
	AddTask(task T) error

	//GetAllTask() ([]T, error)

	//Delete(workflowId string) error
}

type UpdateCommand struct {
	Modified  time.Time
	Status    string
	Variables map[string]any
}

type DB[S any] interface {
	//Save(entity S) error

	Update(command UpdateCommand) (S, error)

	//Get(workflowId string) S
}

type Engine[T, S any] struct {
	Name    string
	RCache  Cache[T]
	Storage DB[S]
}

func NewEngine[T, S any](cache Cache[T], db DB[S]) Engine[T, S] {
	return Engine[T, S]{
		Name:    "Default",
		RCache:  cache,
		Storage: db,
	}
}

func (e *Engine[T, S]) Process() {
	fmt.Println("Process")
	fmt.Println("Saving Data")

	v := make(map[string]any)
	v["1"] = "Awesome"
	c := UpdateCommand{
		Modified:  time.Now(),
		Status:    "s",
		Variables: v,
	}

	e.Storage.Update(c)

}

type CEntity struct {
	Name string
}

type IMCache struct {
}

func NewIMCache() Cache[CEntity] {
	return &IMCache{}
}

func (im *IMCache) AddTask(task CEntity) error {
	return nil
}

type Entity struct {
	Name      string
	Status    string
	Udt       time.Time
	Variables map[string]any
}

type DBStorage struct{}

// Update implements DB
func (*DBStorage) Update(command UpdateCommand) (Entity, error) {
	e := Entity{
		Name:      "A",
		Status:    command.Status,
		Udt:       command.Modified,
		Variables: command.Variables,
	}
	return e, nil
}

func NewDBStorage() DB[Entity] {
	return &DBStorage{}
}
