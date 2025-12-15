package puzzles

import (
	"dec8/helpers"
	"log"
	"math"
	"sort"
)

// Reuse the structures from part1
type Connection2 struct {
	Box1     int     
	Box2     int     
	Distance float64 
}

type UnionFind2 struct {
	parent []int
	rank   []int
}

func NewUnionFind2(n int) *UnionFind2 {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &UnionFind2{parent: parent, rank: rank}
}

func (uf *UnionFind2) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind2) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	
	if rootX == rootY {
		return false
	}
	
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

func (uf *UnionFind2) GetCircuitSizes() []int {
	circuits := make(map[int]int)
	for i := 0; i < len(uf.parent); i++ {
		root := uf.Find(i)
		circuits[root]++
	}
	
	sizes := make([]int, 0, len(circuits))
	for _, size := range circuits {
		sizes = append(sizes, size)
	}
	
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})
	
	return sizes
}

func calculateDistance2(box1, box2 helpers.JunctionBox) float64 {
	dx := float64(box1.XAxis - box2.XAxis)
	dy := float64(box1.YAxis - box2.YAxis)
	dz := float64(box1.ZAxis - box2.ZAxis)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func generateAllConnections2(boxes []helpers.JunctionBox) []Connection2 {
	var connections []Connection2
	
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			distance := calculateDistance2(boxes[i], boxes[j])
			connections = append(connections, Connection2{
				Box1:     i,
				Box2:     j,
				Distance: distance,
			})
		}
	}
	
	sort.Slice(connections, func(i, j int) bool {
		return connections[i].Distance < connections[j].Distance
	})
	
	return connections
}

func Part2() {
    boxes := helpers.ValidateInput()
    connections := generateAllConnections2(boxes)
    uf := NewUnionFind2(len(boxes))
    
    log.Printf("Part 2 - Processing %d junction boxes with %d possible connections", 
        len(boxes), len(connections))
    
    connectionsUsed := 0
    var lastConnectionBoxes [2]helpers.JunctionBox
    
    // Continue making connections until all boxes are in one circuit
    for _, conn := range connections {
        circuitSizes := uf.GetCircuitSizes()
        
        // Check if we already have everything in one circuit
        if len(circuitSizes) == 1 {
            break
        }
        
        if uf.Union(conn.Box1, conn.Box2) {
            connectionsUsed++
            lastConnectionBoxes[0] = boxes[conn.Box1]
            lastConnectionBoxes[1] = boxes[conn.Box2]
            
            // Check after this connection if we now have one circuit
            newCircuitSizes := uf.GetCircuitSizes()
            if len(newCircuitSizes) == 1 {
                log.Printf("Final connection (#%d): Box[%d] (%d,%d,%d) to Box[%d] (%d,%d,%d) - distance %.2f", 
                    connectionsUsed, conn.Box1, lastConnectionBoxes[0].XAxis, lastConnectionBoxes[0].YAxis, lastConnectionBoxes[0].ZAxis,
                    conn.Box2, lastConnectionBoxes[1].XAxis, lastConnectionBoxes[1].YAxis, lastConnectionBoxes[1].ZAxis, conn.Distance)
                break
            }
        }
    }
    
    finalCircuitSizes := uf.GetCircuitSizes()
    
    if len(finalCircuitSizes) == 1 {
        result := lastConnectionBoxes[0].XAxis * lastConnectionBoxes[1].XAxis
        log.Printf("Part 2 - All %d boxes connected in one circuit", finalCircuitSizes[0])
        log.Printf("Last connection X coordinates: %d × %d = %d", 
            lastConnectionBoxes[0].XAxis, lastConnectionBoxes[1].XAxis, result)
    } else {
        log.Printf("Part 2 - Failed to connect all boxes. %d circuits remain", len(finalCircuitSizes))
    }
}
