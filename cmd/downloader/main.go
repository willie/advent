package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/willie/advent/aoc"
	"golang.org/x/net/html"
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

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	var md strings.Builder
	var f func(*html.Node)
	foundArticle := false

	// Recursive function to traverse the HTML tree
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "article" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "day-desc" {
					foundArticle = true
					md.WriteString(convertNodeToMarkdown(n))
					md.WriteString("\n\n")
					return // Don't traverse inside article again with the main loop
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if !foundArticle {
		return fmt.Errorf("no description found in %s", url)
	}

	return os.WriteFile(filepath, []byte(strings.TrimSpace(md.String())), 0644)
}

func convertNodeToMarkdown(n *html.Node) string {
	var sb strings.Builder
	var walk func(*html.Node)

	walk = func(node *html.Node) {
		if node.Type == html.TextNode {
			sb.WriteString(node.Data)
			return
		}

		if node.Type == html.ElementNode {
			switch node.Data {
			case "h2":
				sb.WriteString("## ")
				for c := node.FirstChild; c != nil; c = c.NextSibling {
					walk(c)
				}
				sb.WriteString("\n\n")
			case "p":
				for c := node.FirstChild; c != nil; c = c.NextSibling {
					walk(c)
				}
				sb.WriteString("\n\n")
			case "pre":
				sb.WriteString("```\n")
				// Usually <code> is inside <pre>, let's just grab the text content
				// but sometimes it's <pre><code>...</code></pre>
				// Advent of code puts <code> inside <pre>.
				// Let's rely on the inner code tag handling, or handle here?
				// Actually, AoC uses <pre><code>...</code></pre>.
				// If we encounter <code> inside <pre>, we don't want backticks.
				for c := node.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.ElementNode && c.Data == "code" {
						for cc := c.FirstChild; cc != nil; cc = cc.NextSibling {
							sb.WriteString(cc.Data) // Raw text inside code block
						}
					} else {
						walk(c)
					}
				}
				sb.WriteString("\n```\n\n")
			case "code":
				// Inline code, unless we are already in a pre (handled above)
				// But since we handle pre above by iterating children, this case
				// will be hit if code is used outside pre.
				// However, pre handler doesn't call walk on code node, it manually iterates code children.
				// So this case is only for inline code.
				sb.WriteString("`")
				for c := node.FirstChild; c != nil; c = c.NextSibling {
					walk(c)
				}
				sb.WriteString("`")
			case "em":
				sb.WriteString("**")
				for c := node.FirstChild; c != nil; c = c.NextSibling {
					walk(c)
				}
				sb.WriteString("**")
			case "ul":
				for c := node.FirstChild; c != nil; c = c.NextSibling {
					walk(c)
				}
				sb.WriteString("\n")
			case "li":
				sb.WriteString("- ")
				for c := node.FirstChild; c != nil; c = c.NextSibling {
					walk(c)
				}
				sb.WriteString("\n")
			case "a":
				href := ""
				for _, a := range node.Attr {
					if a.Key == "href" {
						href = a.Val
						break
					}
				}
				sb.WriteString("[")
				for c := node.FirstChild; c != nil; c = c.NextSibling {
					walk(c)
				}
				sb.WriteString("](")
				sb.WriteString(href)
				sb.WriteString(")")
			case "span":
				// Ignore span tags, just print children
				for c := node.FirstChild; c != nil; c = c.NextSibling {
					walk(c)
				}
			default:
				// For other tags (article, div, etc if they exist inside), just walk children
				for c := node.FirstChild; c != nil; c = c.NextSibling {
					walk(c)
				}
			}
		}
	}

	// For the root node (article) passed in, we process its children
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walk(c)
	}

	return sb.String()
}
