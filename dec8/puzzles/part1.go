package puzzles

import (
	"dec8/helpers"
	"log"
	"math"
	"sort"
)

type Connection struct {
	Box1     int     // index of first box
	Box2     int     // index of second box
	Distance float64 // Euclidean distance between boxes
}

// UnionFind represents a disjoint set data structure
type UnionFind struct {
	parent []int
	rank   []int
}

// NewUnionFind creates a new UnionFind with n elements
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i // each element is its own parent initially
		rank[i] = 0
	}
	return &UnionFind{parent: parent, rank: rank}
}

// Find returns the root of the set containing x (with path compression)
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // path compression
	}
	return uf.parent[x]
}

// Union merges the sets containing x and y, returns true if they were different sets
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	
	if rootX == rootY {
		return false // already in same set
	}
	
	// Union by rank
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
	
	return true
}

// GetCircuitSizes returns the sizes of all circuits
func (uf *UnionFind) GetCircuitSizes() []int {
	circuits := make(map[int]int)
	for i := 0; i < len(uf.parent); i++ {
		root := uf.Find(i)
		circuits[root]++
	}
	
	sizes := make([]int, 0, len(circuits))
	for _, size := range circuits {
		sizes = append(sizes, size)
	}
	
	// Sort in descending order
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})
	
	return sizes
}

// calculateDistance computes the Euclidean distance between two junction boxes
func calculateDistance(box1, box2 helpers.JunctionBox) float64 {
	dx := float64(box1.XAxis - box2.XAxis)
	dy := float64(box1.YAxis - box2.YAxis)
	dz := float64(box1.ZAxis - box2.ZAxis)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// generateAllConnections creates a list of all possible connections between boxes
func generateAllConnections(boxes []helpers.JunctionBox) []Connection {
	var connections []Connection
	
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			distance := calculateDistance(boxes[i], boxes[j])
			connections = append(connections, Connection{
				Box1:     i,
				Box2:     j,
				Distance: distance,
			})
		}
	}
	
	// Sort by distance (shortest first)
	sort.Slice(connections, func(i, j int) bool {
		return connections[i].Distance < connections[j].Distance
	})
	
	return connections
}

func Part1() {
	boxes := helpers.ValidateInput()
	connections := generateAllConnections(boxes)
	uf := NewUnionFind(len(boxes))
	
	log.Printf("Processing %d junction boxes with %d possible connections", 
		len(boxes), len(connections))
	
	// Make the 1000 shortest connections (change to 10 for test input debugging)
	connectionsUsed := 0
	actualConnections := 0 // Track both successful unions AND redundant connections
	maxConnections := 1000 // We need to make exactly 1000 connections total
	
	for _, conn := range connections {
		if actualConnections >= maxConnections {
			break
		}
		
		actualConnections++ // Count every connection attempt
		if uf.Union(conn.Box1, conn.Box2) {
			connectionsUsed++ // Count successful unions
			if connectionsUsed <= 5 || actualConnections >= maxConnections-5 || connectionsUsed%100 == 0 {
				log.Printf("Successful connection %d (attempt %d): Box[%d] to Box[%d] (distance %.2f)", 
					connectionsUsed, actualConnections, conn.Box1, conn.Box2, conn.Distance)
			}
		} else {
			if actualConnections >= maxConnections-5 {
				log.Printf("Redundant connection (attempt %d): Box[%d] to Box[%d] already connected (distance %.2f)", 
					actualConnections, conn.Box1, conn.Box2, conn.Distance)
			}
		}
	}
	
	// Calculate result
	circuitSizes := uf.GetCircuitSizes()
	
	log.Printf("After %d total connection attempts (%d successful): %d circuits with sizes %v", 
		actualConnections, connectionsUsed, len(circuitSizes), circuitSizes[:min(len(circuitSizes), 5)])
	
	if len(circuitSizes) >= 3 {
		result := circuitSizes[0] * circuitSizes[1] * circuitSizes[2]
		log.Printf("Answer: %d × %d × %d = %d", 
			circuitSizes[0], circuitSizes[1], circuitSizes[2], result)
	} else if len(circuitSizes) == 1 {
		log.Printf("Only 1 circuit with %d boxes - all connected!", circuitSizes[0])
		log.Printf("Answer: Cannot multiply three circuit sizes (only have 1 circuit)")
	} else {
		log.Printf("Not enough circuits (only %d)", len(circuitSizes))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
