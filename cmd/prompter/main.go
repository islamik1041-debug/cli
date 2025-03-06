package main

import (
	"fmt"
	"log"

	"github.com/cli/cli/v2/internal/prompter"
	"github.com/cli/cli/v2/pkg/iostreams"
)

func main() {
	io := iostreams.System()
	p := prompter.New("vim", io.In, io.Out, io.ErrOut)

	// Demonstrating single-option select / dropdown prompts
	cuisines := []string{"Italian", "Greek", "Indian", "Japanese", "American"}
	favorite, err := p.Select("Favorite cuisine?", "Italian", cuisines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Favorite cuisine: %s\n", cuisines[favorite])

	// Demonstrating multi-option select / dropdown prompts
	favorites, err := p.MultiSelect("Favorite cuisines?", []string{}, cuisines)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range favorites {
		fmt.Printf("Favorite cuisine: %s\n", cuisines[f])
	}

	// Demonstrating text input prompts
	text, err := p.Input("Favorite meal?", "Breakfast")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Favorite meal: %s\n", text)

	// Demonstrating password input prompts
	safeword, err := p.Password("Safe word?")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Safe word: %s\n", safeword)

	// Demonstrating confirmation prompts
	confirmation, err := p.Confirm("Are you sure?", false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Confirmation: %t\n", confirmation)

	// Demonstrating auth token prompt
	token, err := p.AuthToken()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Auth token: %s\n", token)

	// Demonstrating deletion confirmation prompt
	err = p.ConfirmDeletion("delete-me")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Item deleted")

	// Demonstrating hostname input prompt
	hostname, err := p.InputHostname()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hostname: %s\n", hostname)

	// Demonstrating editor prompt
	editorText, err := p.MarkdownEditor("Edit your text:", "Initial text", true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Edited text: %s\n", editorText)
}
