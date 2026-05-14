package main

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

var visited map[int]*Employee

func getImportance(employees []*Employee, id int) int {
	visited = make(map[int]*Employee, 0)
	for _, employee := range employees {
		visited[employee.Id] = employee
	}

	total := DFS(id)
	return total
}

func DFS(id int) int {
	employee := visited[id]

	total := employee.Importance

	for _, sub := range employee.Subordinates {
		total += DFS(sub)
	}

	return total

}
