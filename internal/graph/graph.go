package graph

import (
	"fmt"
	"graphs/internal/queue"
	"graphs/internal/stack"
	"strings"
)

type Vertex[T comparable] struct {
	Value T
	Edges []*Vertex[T]
}

type Graph[T comparable] struct {
	Vertices []*Vertex[T]
}

func (g *Graph[T]) AddVertex(value T) *Vertex[T] {
	vertex := &Vertex[T]{
		Value: value,
	}
	g.Vertices = append(g.Vertices, vertex)
	return vertex
}

func (v *Vertex[T]) AddEdge(vertices ...*Vertex[T]) {
	for _, vertex := range vertices {
		v.Edges = append(v.Edges, vertex)
	}
}

func (g *Graph[T]) ToString() string {
	var sb strings.Builder
	for _, vertex := range g.Vertices {
		sb.WriteString(fmt.Sprintf("%v: ", vertex.Value))
		for _, v := range vertex.Edges {
			sb.WriteString(fmt.Sprintf("%v,", v.Value))
		}
		sb.WriteString(fmt.Sprintf("\n"))
	}
	return sb.String()
}

func (g *Graph[T]) DFS(start, end *Vertex[T]) {
	if start == nil || end == nil {
		return
	}

	s := stack.NewStack[*Vertex[T]]()
	visited := make(map[*Vertex[T]]struct{})
	s.Push(start)

	for s.Len() > 0 {
		vertex, _ := s.Pop()

		if _, ok := visited[vertex]; !ok {
			visited[vertex] = struct{}{}
			fmt.Printf("%v", vertex.Value)

			if vertex == end {
				fmt.Println()
				return
			}

			fmt.Printf(" -> ")

			for _, v := range vertex.Edges {
				if _, ok := visited[v]; !ok {
					s.Push(v)
				}
			}
		}
	}
}

func (g *Graph[T]) BFS(start, end *Vertex[T]) {
	if start == nil || end == nil {
		return
	}

	q := queue.New[*Vertex[T]]()
	visited := make(map[*Vertex[T]]struct{})

	visited[start] = struct{}{}
	q.Enqueue(start)

	for q.Len() > 0 {
		vertex, _ := q.Dequeue()

		fmt.Printf("%v", vertex.Value)

		if vertex == end {
			fmt.Println()
			return
		}

		fmt.Printf(" -> ")

		for _, v := range vertex.Edges {

			if _, ok := visited[v]; !ok {
				visited[v] = struct{}{}
				q.Enqueue(v)
			}
		}
	}
}
