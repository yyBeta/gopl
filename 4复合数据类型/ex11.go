// copy from github.com/suzuken/gopl
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const API = "https://api.github.com"

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func CreateIssue(owner, repo, title, body string) (bool, error) {
	issue := &Issue{
		Title: title,
		Body:  body,
	}
	b, err := json.Marshal(issue)
	if err != nil {
		return false, err
	}
	resp, err := http.Post(API+fmt.Sprintf("/repos/%s/%s/issues", owner, repo), "application/json", bytes.NewReader(b))
	defer resp.Body.Close()
	if err != nil {
		return false, err
	}
	if resp.StatusCode != http.StatusCreated {
		return false, fmt.Errorf("create issue failed: %s", resp.Status)
	}
	return true, nil
}

func ReadIssue(owner, repo, number string) (*Issue, error) {
	resp, err := http.Get(API + fmt.Sprintf("/repos/%s/%s/issues/%s", owner, repo, number))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get issue failed: %s", resp.Status)
	}
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func patch(urlStr string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("PATCH", urlStr, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return http.DefaultClient.Do(req)
}

func UpdateIssue(owner, repo, number, title, body string) (bool, error) {
	issue := &Issue{
		Title: title,
		Body:  body,
	}
	b, err := json.Marshal(issue)
	if err != nil {
		return false, err
	}
	resp, err := patch(API+fmt.Sprintf("/repos/%s/%s/issues/%s", owner, repo, number), bytes.NewReader(b))
	defer resp.Body.Close()
	if err != nil {
		return false, err
	}
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("update issue failed: %s", resp.Status)
	}
	return true, nil
}

func CloseIssue(owner, repo, number string) (bool, error) {
	issue := &Issue{
		State: "closed",
	}
	b, err := json.Marshal(issue)
	if err != nil {
		return false, err
	}
	resp, err := patch(API+fmt.Sprintf("/repos/%s/%s/issues/%s", owner, repo, number), bytes.NewReader(b))
	defer resp.Body.Close()
	if err != nil {
		return false, err
	}
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("close issue failed: %s", resp.Status)
	}
	return true, nil
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `GitHub issue manager (for exercise 4-11 of gopl.io)
Open Issue: create new issue for https://github.com/suzuken/gopl
	create suzuken gopl mytitle testbody
Read Issue: read issue https://github.com/suzuken/gopl/10
	read suzuken gopl 10
Update Issue: update title and body of https://github.com/suzuken/gopl/5.
	update suzuken gopl 5 mytitle-updated testbody-updated
Close Issue: close https://github.com/suzuken/gopl/3
	close suzuken gopl 3
`)
	os.Exit(0)
}

// tool for managing GitHub issues
func main() {
	if len(os.Args) < 2 {
		printUsage()
	}
	switch action := os.Args[1]; action {
	case "create":
		if len(os.Args) != 6 {
			fmt.Fprintf(os.Stderr, "Usage: %s create owner repo title body", os.Args[0])
			os.Exit(1)
		}
		if ok, err := CreateIssue(os.Args[2], os.Args[3], os.Args[4], os.Args[5]); err != nil {
			fmt.Fprintf(os.Stderr, "create issue failed: %s", err)
			os.Exit(1)
		} else if !ok {
			fmt.Fprintf(os.Stderr, "create issue failed: %s", err)
			os.Exit(1)
		} else {
			fmt.Fprint(os.Stdout, "issue created.")
		}
	case "read":
		if len(os.Args) != 5 {
			fmt.Fprintf(os.Stderr, "Usage: %s read owner repo issue-number", os.Args[0])
			os.Exit(1)
		}
		issue, err := ReadIssue(os.Args[2], os.Args[3], os.Args[4])
		if err != nil {
			fmt.Fprintf(os.Stderr, "read issue failed: %s", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%#v", issue)
	case "update":
		if len(os.Args) != 7 {
			fmt.Fprintf(os.Stderr, "Usage: %s update owner repo issue-number title body", os.Args[0])
			os.Exit(1)
		}
		if ok, err := UpdateIssue(os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6]); err != nil {
			fmt.Fprintf(os.Stderr, "update issue failed: %s", err)
			os.Exit(1)
		} else if !ok {
			fmt.Fprintf(os.Stderr, "update issue failed: %s", err)
			os.Exit(1)
		} else {
			fmt.Fprint(os.Stdout, "issue updated.")
		}
	// exercise says "delete", but GitHub issue does not support delete.
	// Instead, close here.
	case "close":
		if len(os.Args) != 5 {
			fmt.Fprintf(os.Stderr, "Usage: %s close owner repo issue-number", os.Args[0])
			os.Exit(1)
		}
		if ok, err := CloseIssue(os.Args[2], os.Args[3], os.Args[4]); err != nil {
			fmt.Fprintf(os.Stderr, "close issue failed: %s", err)
			os.Exit(1)
		} else if !ok {
			fmt.Fprintf(os.Stderr, "close issue failed: %s", err)
			os.Exit(1)
		} else {
			fmt.Fprint(os.Stdout, "issue closed.")
		}
	default:
		fmt.Fprintf(os.Stderr, "1st arguments should be selected from create/read/update/delete.")
		os.Exit(1)
	}
}
