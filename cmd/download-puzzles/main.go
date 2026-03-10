package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	outputDir  = flag.String("out", "puzzles", "output directory for downloaded puzzles")
	startYear  = flag.Int("start", 2015, "first year to download")
	endYear    = flag.Int("end", 2024, "last year to download")
	delay      = flag.Duration("delay", 500*time.Millisecond, "delay between requests (be nice to the server)")
	forceAll   = flag.Bool("force", false, "re-download even if file exists")
	outputHTML = flag.Bool("html", false, "also save raw HTML files")
)

func main() {
	flag.Parse()

	cookie := os.Getenv("AOC_SESSION_COOKIE")
	if cookie == "" {
		log.Fatal("AOC_SESSION_COOKIE environment variable not set")
	}

	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatal(err)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	totalDownloaded := 0

	for year := *startYear; year <= *endYear; year++ {
		yearDir := filepath.Join(*outputDir, fmt.Sprintf("%d", year))
		if err := os.MkdirAll(yearDir, 0755); err != nil {
			log.Fatal(err)
		}

		for day := 1; day <= 25; day++ {
			mdPath := filepath.Join(yearDir, fmt.Sprintf("day%02d.md", day))

			// Skip if already exists (unless force flag)
			if !*forceAll {
				if _, err := os.Stat(mdPath); err == nil {
					fmt.Printf("  [skip] %d day %d (already exists)\n", year, day)
					continue
				}
			}

			url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
			fmt.Printf("Fetching %d day %d... ", year, day)

			html, err := fetchPage(client, url, cookie)
			if err != nil {
				fmt.Printf("ERROR: %v\n", err)
				continue
			}

			// Check if puzzle exists (404 or "please don't repeatedly request this")
			if strings.Contains(html, "Please don't repeatedly request") ||
				strings.Contains(html, "404") ||
				len(html) < 1000 {
				fmt.Println("not available yet")
				continue
			}

			// Extract puzzle content
			content := extractPuzzle(html)
			if content == "" {
				fmt.Println("no puzzle content found")
				continue
			}

			// Convert to markdown
			md := htmlToMarkdown(content, year, day)

			// Save markdown
			if err := os.WriteFile(mdPath, []byte(md), 0644); err != nil {
				log.Printf("Error writing %s: %v", mdPath, err)
				continue
			}

			// Optionally save HTML
			if *outputHTML {
				htmlPath := filepath.Join(yearDir, fmt.Sprintf("day%02d.html", day))
				os.WriteFile(htmlPath, []byte(content), 0644)
			}

			totalDownloaded++
			fmt.Println("OK")

			time.Sleep(*delay)
		}
	}

	fmt.Printf("\nDownloaded %d puzzles to %s/\n", totalDownloaded, *outputDir)
}

func fetchPage(client *http.Client, url, cookie string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
	req.Header.Set("User-Agent", "github.com/willie/advent puzzle-downloader (contact: willie@pobox.com)")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// extractPuzzle pulls out the <article class="day-desc"> content
func extractPuzzle(html string) string {
	// Find all article tags with class="day-desc"
	re := regexp.MustCompile(`(?s)<article class="day-desc">(.*?)</article>`)
	matches := re.FindAllStringSubmatch(html, -1)

	if len(matches) == 0 {
		return ""
	}

	var parts []string
	for _, match := range matches {
		parts = append(parts, match[1])
	}

	return strings.Join(parts, "\n\n---\n\n")
}

// htmlToMarkdown converts AoC puzzle HTML to readable markdown
func htmlToMarkdown(html string, year, day int) string {
	var sb strings.Builder

	// Header
	sb.WriteString(fmt.Sprintf("# Advent of Code %d - Day %d\n\n", year, day))
	sb.WriteString(fmt.Sprintf("**Link:** https://adventofcode.com/%d/day/%d\n\n", year, day))
	sb.WriteString("---\n\n")

	content := html

	// Convert headers
	content = regexp.MustCompile(`<h2[^>]*>--- (.*?) ---</h2>`).ReplaceAllString(content, "## $1\n\n")
	content = regexp.MustCompile(`<h2[^>]*>(.*?)</h2>`).ReplaceAllString(content, "## $1\n\n")

	// Convert emphasis
	content = regexp.MustCompile(`<em class="star">([^<]*)</em>`).ReplaceAllString(content, "**$1**")
	content = regexp.MustCompile(`<em>([^<]*)</em>`).ReplaceAllString(content, "**$1**")
	content = regexp.MustCompile(`<strong>([^<]*)</strong>`).ReplaceAllString(content, "**$1**")

	// Convert code
	content = regexp.MustCompile(`<code>([^<]*)</code>`).ReplaceAllString(content, "`$1`")
	content = regexp.MustCompile(`(?s)<pre><code>(.*?)</code></pre>`).ReplaceAllString(content, "\n```\n$1\n```\n")

	// Convert lists
	content = regexp.MustCompile(`<li>(.*?)</li>`).ReplaceAllString(content, "- $1\n")
	content = regexp.MustCompile(`<ul[^>]*>`).ReplaceAllString(content, "\n")
	content = regexp.MustCompile(`</ul>`).ReplaceAllString(content, "\n")

	// Convert paragraphs
	content = regexp.MustCompile(`<p>(.*?)</p>`).ReplaceAllString(content, "$1\n\n")

	// Convert links
	content = regexp.MustCompile(`<a href="([^"]*)"[^>]*>([^<]*)</a>`).ReplaceAllString(content, "[$2]($1)")

	// Convert spans (often used for highlighting)
	content = regexp.MustCompile(`<span[^>]*title="([^"]*)"[^>]*>([^<]*)</span>`).ReplaceAllString(content, "$2 ($1)")
	content = regexp.MustCompile(`<span[^>]*>([^<]*)</span>`).ReplaceAllString(content, "$1")

	// Remove remaining HTML tags
	content = regexp.MustCompile(`<[^>]+>`).ReplaceAllString(content, "")

	// Decode HTML entities
	content = strings.ReplaceAll(content, "&lt;", "<")
	content = strings.ReplaceAll(content, "&gt;", ">")
	content = strings.ReplaceAll(content, "&amp;", "&")
	content = strings.ReplaceAll(content, "&quot;", "\"")
	content = strings.ReplaceAll(content, "&#39;", "'")
	content = strings.ReplaceAll(content, "&nbsp;", " ")

	// Clean up whitespace
	content = regexp.MustCompile(`\n{3,}`).ReplaceAllString(content, "\n\n")
	content = strings.TrimSpace(content)

	sb.WriteString(content)
	sb.WriteString("\n")

	return sb.String()
}
