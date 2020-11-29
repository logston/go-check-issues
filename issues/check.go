package issues

import (
	"context"
	"fmt"

	"github.com/google/go-github/v32/github"
)

func FetchIssues() error {
	client := github.NewClient(nil)

	ctx := context.Background()

	il, r, err := client.Issues.ListByRepo(ctx, "kubernetes", "kubernetes", &github.IssueListByRepoOptions{
		Assignee: "none",
		Labels:   []string{"good first issue"},
	})
	if err != nil {
		return err
	}

	fmt.Println(r)

	for _, i := range il {
		fmt.Println(i)
	}

	return nil
}
