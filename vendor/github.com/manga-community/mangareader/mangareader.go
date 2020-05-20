package mangareader

import (
	"fmt"
	"strings"

	"github.com/anaskhan96/soup"
)

type Mangareader struct{}

func (m *Mangareader) Latest() ([]string, error) {
	var links []string

	response, err := soup.Get("http://www.mangareader.net/latest")
	if err != nil {
		return nil, err
	}

	doc := soup.HTMLParse(response)
	chapters := doc.Find("div", "id", "latestchapters").FindAll("a")

	for _, chapter := range chapters {

		if chapter.Attrs()["class"] == "chaptersrec" {
			//fmt.Println(chapter.Text(), "| Link :", chapter.Attrs()["href"])
			url := "https://www.mangareader.net" + chapter.Attrs()["href"]
			if IsURLValid(url) {
				links = append(links, url)
			}
		}
	}

	return links, err
}

func (m *Mangareader) RetrieveImageLinks(comic *Comic) ([]string, error) {
	var links []string

	response, err := soup.Get(comic.URLSource)

	if err != nil {
		return nil, err
	}

	doc := soup.HTMLParse(response)
	// retrieve the <option>
	options := doc.FindAll("option")

	for i := 1; i <= len(options); i++ {
		pageLink := fmt.Sprintf("https://%s/%s/%s/%d", comic.Source, comic.Name, comic.IssueNumber, i)
		rsp, soupErr := soup.Get(pageLink)
		if soupErr != nil {
			return nil, soupErr
		}

		doc = soup.HTMLParse(rsp)
		// return the first `<img>`
		// e.g. <img src="http://example.com">
		imgTag := doc.Find("img")
		// doc.Find returns an html.Node
		// the line below will return the src value
		src := imgTag.Pointer.Attr[3].Val
		links = append(links, src)
	}

	return links, err
}

func (m *Mangareader) IsSingleIssue(url string) bool {
	return len(TrimAndSplitURL(url)) >= 5
}

func (m *Mangareader) RetrieveLastIssue(url string) (string, error) {
	url = strings.Join(TrimAndSplitURL(url)[:4], "/")

	response, err := soup.Get(url)
	if err != nil {
		return "", err
	}

	doc := soup.HTMLParse(response)
	chapters := doc.Find("div", "id", "chapterlist").FindAll("a")

	lastIssue := chapters[len(chapters)-1].Attrs()["href"]
	lastIssueUrl := "https://www.mangareader.net" + lastIssue

	return lastIssueUrl, nil
}

// RetrieveIssueLinks gets a slice of urls for all issues in a comic
func (m *Mangareader) RetrieveIssueLinks(url string, all, last bool) ([]string, error) {
	if last {
		lastIssue, err := m.RetrieveLastIssue(url)
		return []string{lastIssue}, err
	}

	if all && m.IsSingleIssue(url) {
		url = strings.Join(TrimAndSplitURL(url)[:4], "/")
	} else if m.IsSingleIssue(url) {
		return []string{url}, nil
	}

	var links []string

	response, err := soup.Get(url)
	if err != nil {
		return nil, err
	}

	doc := soup.HTMLParse(response)
	chapters := doc.Find("div", "id", "chapterlist").FindAll("a")

	for _, chapter := range chapters {
		url := "https://www.mangareader.net" + chapter.Attrs()["href"]
		if IsURLValid(url) {
			links = append(links, url)
		}
	}

	return links, err
}

func (m *Mangareader) GetInfo(url string) (string, string) {
	parts := TrimAndSplitURL(url)
	name := parts[3]
	issueNumber := parts[4]

	return name, issueNumber
}

// Initialize loads links and metadata from mangareader
func (m *Mangareader) Initialize(comic *Comic) error {
	name, issueNumber := m.GetInfo(comic.URLSource)
	comic.Name = name
	comic.IssueNumber = issueNumber

	links, err := m.RetrieveImageLinks(comic)
	comic.Links = links

	return err
}
