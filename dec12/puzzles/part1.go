package puzzles

import (
	"dec12/helpers"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Part1() {
    blocks, summaries, err := helpers.ValidateInput()
    if err != nil {
        log.Println("Error:", err)
        return
    }

    // Parse regions
    type Region struct {
        W, H int
        ShapeCounts []int
    }
    var regions []Region
    for _, line := range summaries {
        parts := strings.Split(line, ":")
        wh := strings.Split(strings.TrimSpace(parts[0]), "x")
        w, _ := strconv.Atoi(wh[0])
        h, _ := strconv.Atoi(wh[1])
        counts := []int{}
        for _, c := range strings.Fields(parts[1]) {
            v, _ := strconv.Atoi(c)
            counts = append(counts, v)
        }
        regions = append(regions, Region{W: w, H: h, ShapeCounts: counts})
    }

    // Prepare all orientations for each shape
    shapeOrientations := make([][][][]rune, len(blocks))
    shapeAreas := make([]int, len(blocks))
    for i, block := range blocks {
        shapeOrientations[i] = helpers.UniqueOrientations(block.Grid)
        // Calculate area
        area := 0
        for _, row := range block.Grid {
            for _, c := range row {
                if c == '#' {
                    area++
                }
            }
        }
        shapeAreas[i] = area
    }

    type Present struct {
        ShapeIdx int
        OrientIdx int
        Area int
    }

    // Try to fit all presents in a region
    canFit := func(region Region) bool {
        grid := make([][]int, region.H)
        for i := range grid {
            grid[i] = make([]int, region.W)
        }
        // Build presents list
        presents := []Present{}
        totalArea := 0
        for shapeIdx, count := range region.ShapeCounts {
            for i := 0; i < count; i++ {
                presents = append(presents, Present{ShapeIdx: shapeIdx, Area: shapeAreas[shapeIdx]})
                totalArea += shapeAreas[shapeIdx]
            }
        }
        // Early prune: total present area > grid area
        if totalArea > region.W*region.H {
            return false
        }
        // Sort presents by area descending (place big/awkward first)
        sort.Slice(presents, func(i, j int) bool { return presents[i].Area > presents[j].Area })
        // Try all permutations recursively
        var place func(idx int) bool
        place = func(idx int) bool {
            if idx == len(presents) {
                return true
            }
            shapeIdx := presents[idx].ShapeIdx
            for orientIdx, shape := range shapeOrientations[shapeIdx] {
                h, w := len(shape), len(shape[0])
                for y := 0; y <= region.H-h; y++ {
                    for x := 0; x <= region.W-w; x++ {
                        // Check if can place
                        canPlace := true
                        for i := 0; i < h && canPlace; i++ {
                            for j := 0; j < w; j++ {
                                if shape[i][j] == '#' && grid[y+i][x+j] != 0 {
                                    canPlace = false
                                    break
                                }
                            }
                        }
                        if canPlace {
                            // Place
                            for i := 0; i < h; i++ {
                                for j := 0; j < w; j++ {
                                    if shape[i][j] == '#' {
                                        grid[y+i][x+j] = idx + 1
                                    }
                                }
                            }
                            presents[idx].OrientIdx = orientIdx
                            if place(idx+1) {
                                return true
                            }
                            // Unplace
                            for i := 0; i < h; i++ {
                                for j := 0; j < w; j++ {
                                    if shape[i][j] == '#' {
                                        grid[y+i][x+j] = 0
                                    }
                                }
                            }
                        }
                    }
                }
            }
            return false
        }
        return place(0)
    }

    fitCount := 0
    for i, region := range regions {
        ok := canFit(region)
        fmt.Printf("Region %d (%dx%d): %v => %v\n", i+1, region.W, region.H, region.ShapeCounts, ok)
        if ok {
            fitCount++
        }
    }
    fmt.Printf("\nTotal regions that can fit all presents: %d\n", fitCount)
}
