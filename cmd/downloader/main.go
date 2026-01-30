package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/willie/advent/aoc"
)

func main() {
	log.SetFlags(0)
	cookie := aoc.SessionCookie()

	now := time.Now()
	// AoC unlocks at midnight EST (UTC-5)
	est := time.FixedZone("EST", -5*60*60)
	nowInEst := now.In(est)
	currentYear := nowInEst.Year()

	fmt.Printf("Downloading puzzles up to %d/%d/%d (EST)\n", currentYear, nowInEst.Month(), nowInEst.Day())

	for y := 2015; y <= currentYear; y++ {
		for d := 1; d <= 25; d++ {
			if y == currentYear {
				if nowInEst.Month() < time.December {
					// It's before December of the current year, so no puzzles for this year yet.
					break
				}
				if nowInEst.Month() == time.December && d > nowInEst.Day() {
					// It's December, but this day hasn't unlocked yet.
					break
				}
			}

			yearDir := fmt.Sprintf("%d", y)
			fileBase := fmt.Sprintf("%d-%02d", y, d)
			mdPath := filepath.Join(yearDir, fileBase+".md")
			txtPath := filepath.Join(yearDir, fileBase+".txt")

			if err := os.MkdirAll(yearDir, 0755); err != nil {
				log.Fatalf("failed to create dir %s: %v", yearDir, err)
			}

			// Download Input
			if _, err := os.Stat(txtPath); os.IsNotExist(err) {
				fmt.Printf("Downloading input %d/%d...\n", y, d)
				url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", y, d)
				if err := downloadFile(url, txtPath, cookie); err != nil {
					log.Printf("Failed to download input %s: %v", url, err)
				}
				time.Sleep(1 * time.Second) // Respect AoC rate limit
			}

			// Download Description
			if _, err := os.Stat(mdPath); os.IsNotExist(err) {
				fmt.Printf("Downloading description %d/%d...\n", y, d)
				url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", y, d)
				if err := downloadDescription(url, mdPath, cookie); err != nil {
					log.Printf("Failed to download description %s: %v", url, err)
				}
				time.Sleep(1 * time.Second) // Respect AoC rate limit
			}
		}
	}
}

func downloadFile(url, filepath string, cookie string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
	req.Header.Set("User-Agent", "github.com/willie/advent downloader")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("status code %d", resp.StatusCode)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func downloadDescription(url, filepath string, cookie string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
	req.Header.Set("User-Agent", "github.com/willie/advent downloader")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("status code %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	body := string(bodyBytes)

	re := regexp.MustCompile(`(?s)<article class="day-desc">(.*?)</article>`)
	matches := re.FindAllStringSubmatch(body, -1)

	if len(matches) == 0 {
		return fmt.Errorf("no description found in %s", url)
	}

	var md strings.Builder
	for _, m := range matches {
		md.WriteString(convertHTMLToMarkdown(m[1]))
		md.WriteString("\n\n")
	}

	return os.WriteFile(filepath, []byte(strings.TrimSpace(md.String())), 0644)
}

func convertHTMLToMarkdown(htmlContent string) string {
	s := htmlContent

	// Remove <span> tags (keep content)
	s = regexp.MustCompile(`</?span[^>]*>`).ReplaceAllString(s, "")

	// Headers <h2> -> ##
	s = regexp.MustCompile(`(?s)<h2[^>]*>(.*?)</h2>`).ReplaceAllString(s, "## $1\n")

	// Paragraphs <p> -> text\n\n
	s = regexp.MustCompile(`(?s)<p[^>]*>(.*?)</p>`).ReplaceAllString(s, "$1\n\n")

	// Code blocks <pre><code>...</code></pre> -> ```...```
	s = regexp.MustCompile(`(?s)<pre><code>(.*?)</code></pre>`).ReplaceAllString(s, "```\n$1\n```\n")

	// Inline code <code>...</code>
	s = regexp.MustCompile(`(?s)<code>(.*?)</code>`).ReplaceAllString(s, "`$1`")

	// Bold <em>...</em>
	s = regexp.MustCompile(`(?s)<em>(.*?)</em>`).ReplaceAllString(s, "**$1**")

	// Links <a href="...">...</a>
	s = regexp.MustCompile(`(?s)<a href="([^"]+)">(.*?)</a>`).ReplaceAllString(s, "[$2]($1)")

	// Lists <ul><li>...</li></ul>
	s = regexp.MustCompile(`(?s)<ul>\s*`).ReplaceAllString(s, "")
	s = regexp.MustCompile(`(?s)</ul>`).ReplaceAllString(s, "")
	s = regexp.MustCompile(`(?s)<li>(.*?)</li>`).ReplaceAllString(s, "- $1\n")

	// Decode HTML entities
	s = html.UnescapeString(s)

	return strings.TrimSpace(s)
}
