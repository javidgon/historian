package main

import (
  "os"
  "time"
  "log"
  "bufio"
  "fmt"
  "strings"
  "strconv"
  "github.com/codegangsta/cli"
)

type Commit struct {
  Hash        string
  AuthorEmail string
  Date        string
  Message     string
}

func getHash(s string) string {
  splitted := strings.Split(s, " ")
  // Hash is always in the second position
  return splitted[1]
}

func getAuthorEmail(s string) string {
  open := strings.Index(s, "<") + 1
  end := strings.Index(s, ">")
  return s[open:end]
}

func getDate(s string) string {
  open := strings.Index(s, ">") + 2
  end := strings.Index(s, "+") + 5
  return s[open:end]
}

func getMessage(s string) string {
  open := strings.Index(s, ":") + 2
  return s[open:]
}

func extractCommit(s string) *Commit {
  commit := new(Commit)
  commit.Hash = getHash(s)
  commit.AuthorEmail = getAuthorEmail(s)
  commit.Date = getDate(s)
  commit.Message = getMessage(s)
  return commit
}

func main()  {
  app := cli.NewApp()
  app.Name = "Historian"
  app.Usage = "Fetch all the commit messages and create a release notes document"
  app.Action = func (c *cli.Context)  {
    // Select file with the commit's messages
    file, err := os.Open(".git/logs/HEAD")
    if err != nil {
      log.Fatal(err)
    }
    // Don't forget to close the file at the end
    defer file.Close()

    scanner := bufio.NewScanner(file)
    fmt.Println("- Release notes (" + time.Now().Format(time.RFC850) + "):")
    // Iterate over each line
    counter := 1
    for scanner.Scan() {
      commit := extractCommit(scanner.Text())
      fmt.Println(
        strconv.Itoa(counter) + ") " + commit.Hash + " <" + commit.AuthorEmail +
        "> (" + commit.Date + ") " + commit.Message)
      counter = counter + 1
    }

    if err := scanner.Err(); err != nil {
      log.Fatal(err)
    }
  }
  app.Run(os.Args)
}
