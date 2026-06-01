package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	colors := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
	}

	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		colors[idx] = 1
		for _, nbr := range graph[idx] {
			if colors[nbr] == 1 || colors[nbr] == 0 && dfs(nbr) {
				return true
			}
		}
		colors[idx] = 2
		return false
	}

	for i, c := range colors {
		if c == 0 && dfs(i) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}
