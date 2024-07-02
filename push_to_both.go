package main

import (
	"fmt"
	"log"
	"os/exec"	
)

func runGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()
	return cmd.Run()
} 

func main() {
	// // Define the remote repository URLs
	// githubRepo := "git@github.com:github.com/garveyshah/guess-it-1-gh.git"
	// giteaRepo := "git@gitea.learn.zone01kisumu.ke:oumouma/guess-it-1.git"

	// // Add remote repositories if they don't exist.
	// err := runGitCommand("remote","add","github", githubRepo)
	// if err != nil

	// Fetch updates from both repositories.
	fmt.Println("Fetching updates from Github and Gitea...")
	err := runGitCommand("fetch", "github")
	if err != nil {
		log.Fatalf("Failed to fetch from Github: %v", err)
	}
	err = runGitCommand("fetch", "gitea")
	if err != nil {
		log.Fatalf("Failed to fetch from Gitea: %v", err)
	}

	// Push changes to both repositories
	fmt.Println("Pushing changes to Github and Gitea...")
	err = runGitCommand("push","github","main")
	if err != nil {
		log.Fatalf("Failed to push to Github: %v", err)
	}
	err = runGitCommand("push","gitea","master")
	if err != nil {
		log.Fatalf("Failed to push to Gitea: %v", err)
	}
	fmt.Println("Successfully pushed changes to both Github and Gitea.")
}