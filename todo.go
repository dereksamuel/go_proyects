package main

import "fmt"

type todo struct {
	name        string
	description string
	done        bool
}

func (t *todo) update(isDone bool, description string, name string) {
	t.done = isDone
	if description != "" {
		t.description = description
	}

	if name != "" {
		t.name = name
	}
}

type todoList struct {
	todos []*todo /* [] es un slice. *todo es para actualizar valores en el futuro */
} /* son similares a arrays excepto que si son mutables */

func (tl *todoList) addTodo(t *todo) {
	tl.todos = append(tl.todos, t)
}

func (tl *todoList) deleteTodo(index int) {
	tl.todos = append(tl.todos[:index], tl.todos[index+1:]...) // add a portion of the list, index + 1 si es el ultimo es zero value
}

func (t todoList) printTodos() {
	for _, todo := range t.todos { // i es el indice
		fmt.Println("Name:", todo.name, "Is Done:", todo.done)
	}
	fmt.Println("")
}

func (t todoList) printDoneTodos() {
	for _, todo := range t.todos { // i es el indice
		if todo.done {
			fmt.Println("****** DONE TODO ******")
			fmt.Println("Name:", todo.name, "Is Done:", todo.done)
		}
	}
	fmt.Println("")
}

func main() {
	t1 := &todo{
		name:        "Done Go's course",
		description: "Done this course on this weekend",
	} /* apuntamos a la referencia no a una copia */
	t2 := &todo{
		name:        "DK artist notebook buy",
		description: "Go to DollarCity",
	}
	t3 := &todo{
		name:        "To complete work space",
		description: "doh",
		done:        true,
	}

	list := &todoList{
		todos: []*todo{
			t1, t2,
		},
	}

	list.printTodos()
	list.addTodo(t3)
	list.printTodos()
	list.deleteTodo(0)
	list.printTodos()
	list.printDoneTodos()

	// fmt.Printf("%+v\n", list)
	// fmt.Printf("%+v\n", t) // vamos a mostrar interface key value: %+v

	todoMaps := make(map[string]*todoList)
	todoMaps["Derek"] = list

	t4 := &todo{
		name:        "Done Go's course GOOGLE",
		description: "Done this course on this weekend",
	} /* apuntamos a la referencia no a una copia */
	t5 := &todo{
		name:        "DK artist notebook buy APPLE",
		description: "Go to DollarCity",
	}
	t6 := &todo{
		name:        "To complete work space SUBSTANCE",
		description: "doh",
		done:        true,
	}

	list2 := &todoList{
		todos: []*todo{
			t4, t5, t6,
		},
	}

	todoMaps["Hector"] = list2

	todoMaps["Hector"].printDoneTodos()
}
