package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	DFS = iota
	BFS
	GBFS
	ASTAR
	DIJKSTRA
)

type Point struct {
	Row int
	Col int
}
type Wall struct {
	State Point
	wall  bool
}
type MAZE struct {
	Height int
	Width  int
	Start  Point
	End    Point
	Walls  [][]Wall
}

func main() {
	var m MAZE
	var maze, searchType string

	flag.StringVar(&maze, "file", "maze-100-steps.txt", "maze file")
	flag.StringVar(&searchType, "search", "dfs", "search type")
	flag.Parse()

	err := m.Load(maze)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("maze height/width:", m.Height, m.Width)

}

func (g *MAZE) Load(fileName string) error {
	f, error := os.Open(fileName)
	if error != nil {
		fmt.Printf("Error opening %s: %s", fileName, error)
	}
	defer f.Close()

	var fileContent []string

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("Error reading %s: %s", fileName, err)
		}
		fileContent = append(fileContent, line)
	}

	foundStart, foundEnd := false, false

	for _, line := range fileContent {
		if strings.Contains(line, "A") {
			foundStart = true
		}
		if strings.Contains(line, "B") {
			foundEnd = true
		}
	}

	if !foundStart {
		return fmt.Errorf("Start point 'A' not found in %s", fileName)
	}
	if !foundEnd {
		return fmt.Errorf("End point 'B' not found in %s", fileName)
	}

	g.Height = len(fileContent)

	g.Width = len(fileContent[0])

	var rows [][]Wall
	for i, row := range fileContent {
		var cols []Wall
		for j, col := range row {
			currLetter := fmt.Sprintf("%c", col)
			var wall Wall
			switch currLetter {
			case "A":
				g.Start = Point{Row: i, Col: j}
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = false
			case "B":
				g.End = Point{Row: i, Col: j}
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = false
			case " ":
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = false

			case "#":
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = true
			default:
				continue
			}
			cols = append(cols, wall)
		}
		rows = append(rows, cols)
	}
	g.Walls = rows
	return nil

}
