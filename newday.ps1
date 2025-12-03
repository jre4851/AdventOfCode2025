# Find all decN folders and get the next day number
$existing = Get-ChildItem -Directory | Where-Object { $_.Name -match '^dec\d+$' }
$maxDay = ($existing | ForEach-Object { [int]($_.Name -replace 'dec', '') } | Measure-Object -Maximum).Maximum
if (-not $maxDay) { $maxDay = 0 }
$Day = $maxDay + 1

$base = "dec$Day"
$importHelpers = "`"dec$Day/helpers`""

New-Item -ItemType Directory -Path $base -Force
New-Item -ItemType Directory -Path "$base\puzzles" -Force
New-Item -ItemType Directory -Path "$base\helpers" -Force
New-Item -ItemType Directory -Path "$base\input" -Force

Set-Content -Path "$base\puzzles\part1.go" -Value @"
package puzzles

import (
    "log"
    //$importHelpers
)

func Part1() {
    log.Println("Part 1 placeholder")
}
"@

Set-Content -Path "$base\puzzles\part2.go" -Value @"
package puzzles

import (
    "log"
    //$importHelpers
)

func Part2() {
    log.Println("Part 2 placeholder")
}
"@

Set-Content -Path "$base\helpers\validateInput.go" -Value @"
package helpers

func ValidateInput() {
    // TODO: Implement input validation
}
"@

Set-Content -Path "$base\input\aocInput.txt" -Value ''
Set-Content -Path "$base\input\testInput.txt" -Value ''

Set-Content -Path "$base\main.go" -Value @"
package main

import "dec$Day/puzzles"

func main() {
    puzzles.Part1()
    puzzles.Part2()
}
"@

Set-Content -Path "$base\go.mod" -Value "module dec$Day`r`n`r`ngo 1.25.4`r`n"