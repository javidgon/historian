package main

import (
	"flag"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/libgit2/git2go"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

type Blob struct {
	*git.Blob
	id *git.Oid
}

type Commit struct {
	Hash        string
	AuthorEmail string
	Date        time.Time
	Message     string
}

// Implement the sort.Sort interface so we can sort by Date
type ByDate []Commit

func (d ByDate) Len() int           { return len(d) }
func (d ByDate) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d ByDate) Less(i, j int) bool { return d[i].Date.Before(d[j].Date) }

// Create the Release notes file using all the commits between the last one and the one provided
// in the command's arguments (including this one as well)
func createReleaseNotesFile(commits []Commit, lastCommitWhen time.Time) {
	sort.Sort(ByDate(commits))
	f, err := os.Create("release_notes.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString("==========================================================\n")
	f.WriteString("** Release Notes (" + time.Now().Format(time.RFC850) + "):\n")
	for idx, el := range commits {
		if el.Date.After(lastCommitWhen) || el.Date.Equal(lastCommitWhen) {
			f.WriteString(strconv.Itoa(idx) + ") <" + el.AuthorEmail + "> (" + el.Date.String() + ") " + el.Message)
		}
	}
	f.WriteString("==========================================================\n")
	f.Sync()
}

func main() {
	app := cli.NewApp()
	app.Name = "Historian"
	app.Usage = "Fetch all the commit messages and create a release notes document"
	app.Action = func(c *cli.Context) {
		var path string
		var hash string
		if len(c.Args()) > 1 {
			path = c.Args()[0]
			hash = c.Args()[1]
		} else {
			log.Fatal("Sorry, but you need to provide both a Path and a Commit Hash")
		}

		// According to the `git2go` documentation this is the way of accessing to a repository
		repoPath := flag.String("repo", path, "path to the git repository")
		flag.Parse()

		repo, err := git.OpenRepository(*repoPath)
		if err != nil {
			log.Fatal(err)
		}
		odb, err := repo.Odb()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Processing...")
		commits := make([]Commit, 0)
		var lastCommitWhen time.Time

		odb.ForEach(func(oid *git.Oid) error {
			obj, _ := repo.Lookup(oid)
			switch obj := obj.(type) {
			// We only want to extract the `git.Commit` type
			case *git.Commit:
				commit := new(Commit)
				author := obj.Author()
				commit.Hash = obj.Id().String()
				commit.AuthorEmail = author.Email
				commit.Date = author.When
				commit.Message = obj.Message()
				commits = append(commits, *commit)
				// If the hash matches we store the date as this is the last commit we are going to show
				if commit.Hash == hash {
					lastCommitWhen = commit.Date
				}
			}

			return nil
		})
		createReleaseNotesFile(commits, lastCommitWhen)
		fmt.Println("********************************************")
		fmt.Println("release_notes.txt file successfully created!")
		fmt.Println("********************************************")
	}
	app.Run(os.Args)
}
