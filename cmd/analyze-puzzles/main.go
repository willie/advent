package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var (
	puzzleDir = flag.String("dir", "puzzles", "directory containing downloaded puzzles")
	verbose   = flag.Bool("v", false, "verbose output - show matching puzzles for each pattern")
)

// Pattern represents an algorithmic/conceptual pattern to detect
type Pattern struct {
	Name     string
	Keywords []string
	Matches  []string // puzzles that match
}

var patterns = []Pattern{
	// Graph algorithms
	{Name: "Graph/Network", Keywords: []string{"graph", "node", "edge", "vertex", "vertices", "connected", "path", "network", "route"}},
	{Name: "Shortest Path", Keywords: []string{"shortest", "fewest steps", "minimum steps", "quickest", "fastest route", "dijkstra"}},
	{Name: "BFS/DFS", Keywords: []string{"breadth", "depth", "traverse", "explore", "visit", "reachable", "flood fill"}},
	{Name: "A* / Heuristic Search", Keywords: []string{"heuristic", "a-star", "a*", "best path", "optimal path"}},

	// Dynamic Programming
	{Name: "Dynamic Programming", Keywords: []string{"memoiz", "cache", "previous result", "subproblem", "optimal substructure"}},
	{Name: "Counting/Combinations", Keywords: []string{"how many ways", "how many different", "number of ways", "combinations", "permutations", "arrangements"}},

	// Data Structures
	{Name: "Grid/2D Array", Keywords: []string{"grid", "row", "column", "2d", "map", "floor", "wall", "tile", "coordinate"}},
	{Name: "3D Space", Keywords: []string{"3d", "cube", "x,y,z", "three-dimensional", "3-dimensional"}},
	{Name: "Tree Structure", Keywords: []string{"tree", "parent", "child", "root", "leaf", "branch", "ancestor", "descendant"}},
	{Name: "Stack/Queue", Keywords: []string{"stack", "queue", "push", "pop", "lifo", "fifo", "bracket", "parenthes"}},
	{Name: "Linked List", Keywords: []string{"linked", "next", "previous", "chain", "circular"}},
	{Name: "Hash/Map", Keywords: []string{"lookup", "dictionary", "mapping", "associate", "key-value"}},
	{Name: "Set Operations", Keywords: []string{"unique", "distinct", "duplicate", "intersection", "union", "overlap"}},
	{Name: "Priority Queue/Heap", Keywords: []string{"priority", "heap", "minimum", "maximum", "smallest first", "largest first"}},

	// String/Parsing
	{Name: "String Parsing", Keywords: []string{"parse", "extract", "pattern", "format", "syntax", "token"}},
	{Name: "Regular Expressions", Keywords: []string{"regex", "regexp", "match", "pattern matching"}},
	{Name: "State Machine", Keywords: []string{"state", "transition", "automaton", "machine", "mode"}},

	// Math
	{Name: "Modular Arithmetic", Keywords: []string{"modulo", "remainder", "mod ", "divisible", "chinese remainder"}},
	{Name: "Number Theory", Keywords: []string{"prime", "factor", "gcd", "lcm", "greatest common", "least common"}},
	{Name: "Binary/Bitwise", Keywords: []string{"binary", "bit", "bitwise", "xor", "and", "or", "mask", "shift"}},
	{Name: "Geometry", Keywords: []string{"distance", "manhattan", "euclidean", "area", "perimeter", "polygon", "angle"}},
	{Name: "Linear Algebra", Keywords: []string{"matrix", "matrices", "vector", "transform", "rotation"}},
	{Name: "Range/Interval", Keywords: []string{"range", "interval", "overlap", "between", "from x to y"}},

	// Simulation
	{Name: "Simulation", Keywords: []string{"simulate", "step", "tick", "round", "turn", "after n", "iteration"}},
	{Name: "Cellular Automaton", Keywords: []string{"neighbor", "adjacent", "surrounding", "conway", "game of life", "rules"}},
	{Name: "Physics/Movement", Keywords: []string{"velocity", "acceleration", "position", "move", "direction", "north", "south", "east", "west"}},

	// Optimization
	{Name: "Optimization/Search", Keywords: []string{"maximize", "minimize", "optimal", "best", "most", "least", "fewest"}},
	{Name: "Brute Force", Keywords: []string{"all possible", "every combination", "exhaustive", "try each"}},
	{Name: "Binary Search", Keywords: []string{"binary search", "narrow down", "half", "bisect"}},
	{Name: "Greedy", Keywords: []string{"greedy", "always choose", "immediate", "local optimal"}},

	// Special Techniques
	{Name: "Recursion", Keywords: []string{"recursive", "self-similar", "fractal", "nested"}},
	{Name: "Cycle Detection", Keywords: []string{"cycle", "repeat", "period", "loop", "pattern repeats"}},
	{Name: "Compression/Encoding", Keywords: []string{"compress", "encode", "decode", "checksum", "hash"}},
	{Name: "Assembly/VM", Keywords: []string{"instruction", "register", "opcode", "program", "execute", "assembly", "intcode"}},
	{Name: "Reverse Engineering", Keywords: []string{"reverse", "figure out", "deduce", "work backwards"}},

	// Input Types
	{Name: "Large Numbers", Keywords: []string{"trillion", "billion", "million", "1000000", "very large", "huge number"}},
	{Name: "ASCII Art", Keywords: []string{"ascii", "letters", "display", "render", "print", "pixels"}},
}

func main() {
	flag.Parse()

	files, err := filepath.Glob(filepath.Join(*puzzleDir, "*", "*.md"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding puzzles: %v\n", err)
		os.Exit(1)
	}

	if len(files) == 0 {
		fmt.Println("No puzzle files found in", *puzzleDir)
		fmt.Println("Run download-puzzles first to fetch the puzzles.")
		os.Exit(1)
	}

	fmt.Printf("Analyzing %d puzzles...\n\n", len(files))

	// Analyze each puzzle
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			continue
		}
		text := strings.ToLower(string(content))

		// Extract year/day from path
		rel, _ := filepath.Rel(*puzzleDir, file)
		puzzleID := strings.TrimSuffix(rel, ".md")

		// Check each pattern
		for i := range patterns {
			for _, kw := range patterns[i].Keywords {
				if strings.Contains(text, strings.ToLower(kw)) {
					patterns[i].Matches = append(patterns[i].Matches, puzzleID)
					break
				}
			}
		}
	}

	// Sort patterns by frequency
	sort.Slice(patterns, func(i, j int) bool {
		return len(patterns[i].Matches) > len(patterns[j].Matches)
	})

	// Print results
	fmt.Println("=== Pattern Frequency Analysis ===")
	fmt.Println()

	maxNameLen := 0
	for _, p := range patterns {
		if len(p.Name) > maxNameLen {
			maxNameLen = len(p.Name)
		}
	}

	for _, p := range patterns {
		count := len(p.Matches)
		if count == 0 {
			continue
		}

		// Visual bar
		bar := strings.Repeat("█", count/2)
		if count%2 == 1 {
			bar += "▌"
		}

		fmt.Printf("%-*s %3d %s\n", maxNameLen, p.Name, count, bar)

		if *verbose && count > 0 {
			// Group by year
			years := make(map[string][]string)
			for _, m := range p.Matches {
				parts := strings.Split(m, string(filepath.Separator))
				if len(parts) >= 2 {
					year := parts[0]
					day := parts[1]
					years[year] = append(years[year], day)
				}
			}

			var yearKeys []string
			for y := range years {
				yearKeys = append(yearKeys, y)
			}
			sort.Strings(yearKeys)

			for _, year := range yearKeys {
				days := years[year]
				sort.Strings(days)
				fmt.Printf("%*s  %s: %s\n", maxNameLen, "", year, strings.Join(days, ", "))
			}
			fmt.Println()
		}
	}

	// Summary stats
	fmt.Println()
	fmt.Println("=== Recommended Libraries/Tools ===")
	fmt.Println()

	recommendations := analyzeRecommendations(patterns)
	for _, rec := range recommendations {
		fmt.Println(rec)
	}
}

func analyzeRecommendations(patterns []Pattern) []string {
	var recs []string

	// Check what's common and recommend accordingly
	counts := make(map[string]int)
	for _, p := range patterns {
		counts[p.Name] = len(p.Matches)
	}

	if counts["Grid/2D Array"] > 10 {
		recs = append(recs, "• Grid utilities: 2D array helpers, neighbor iteration, boundary checking")
	}
	if counts["Shortest Path"] > 5 || counts["BFS/DFS"] > 5 {
		recs = append(recs, "• Graph library: BFS, DFS, Dijkstra, A* implementations")
	}
	if counts["Parsing/String"] > 5 || counts["String Parsing"] > 5 {
		recs = append(recs, "• Parsing utilities: regex helpers, string splitting, number extraction")
	}
	if counts["Counting/Combinations"] > 5 || counts["Dynamic Programming"] > 3 {
		recs = append(recs, "• Memoization: generic memoization wrapper for recursive functions")
	}
	if counts["Set Operations"] > 5 {
		recs = append(recs, "• Set data structure: with union, intersection, difference operations")
	}
	if counts["Number Theory"] > 3 {
		recs = append(recs, "• Math utilities: GCD, LCM, prime factorization, modular arithmetic")
	}
	if counts["Cycle Detection"] > 2 {
		recs = append(recs, "• Cycle detection: Floyd's algorithm, Brent's algorithm")
	}
	if counts["Priority Queue/Heap"] > 2 {
		recs = append(recs, "• Priority queue: min-heap and max-heap implementations")
	}
	if counts["Assembly/VM"] > 3 {
		recs = append(recs, "• Simple VM: register-based interpreter for Intcode-style problems")
	}
	if counts["Geometry"] > 3 {
		recs = append(recs, "• Geometry: Point/Vector types, distance functions, area calculations")
	}
	if counts["Range/Interval"] > 3 {
		recs = append(recs, "• Interval/Range: range merging, overlap detection, interval trees")
	}
	if counts["Binary/Bitwise"] > 3 {
		recs = append(recs, "• Bitwise utilities: bit manipulation helpers, binary string conversion")
	}

	// Add some universal recommendations
	recs = append(recs, "")
	recs = append(recs, "Universal recommendations:")
	recs = append(recs, "• Input parsing: automatic number/grid/list detection")
	recs = append(recs, "• Test harness: easy example validation before real input")
	recs = append(recs, "• Visualization: ASCII grid printing for debugging")

	return recs
}

// extractTopics looks for specific topic mentions using regex
func extractTopics(content string) []string {
	var topics []string

	topicPatterns := map[string]*regexp.Regexp{
		"elves":     regexp.MustCompile(`(?i)\belf|elves\b`),
		"santa":     regexp.MustCompile(`(?i)\bsanta\b`),
		"reindeer":  regexp.MustCompile(`(?i)\breindeer\b`),
		"submarine": regexp.MustCompile(`(?i)\bsubmarine\b`),
		"rocket":    regexp.MustCompile(`(?i)\brocket|spacecraft\b`),
	}

	for topic, re := range topicPatterns {
		if re.MatchString(content) {
			topics = append(topics, topic)
		}
	}

	return topics
}
