package menu

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
)

type JournalAction string

const (
	JournalActionNew     JournalAction = "Neues Journal anlegen"
	JournalActionEdit    JournalAction = "Journal bearbeiten"
	JournalActionPreview JournalAction = "Journal anzeigen"
	JournalActionExit    JournalAction = "Journal-App beenden"
	JournalActionDelete  JournalAction = "Journal Eintrag löschen"
)

type Menu struct {
	prompt               string        // the prompt to display to the user
	currentJournalAction JournalAction // the choice the user made
}

func (m *Menu) chooseJournalAction() {

	prompt := promptui.Select{
		Label: "Was möchtest du tun?",
		Items: []JournalAction{JournalActionNew, JournalActionEdit, JournalActionPreview, JournalActionDelete, JournalActionExit},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Ein Fehler ist aufgetreten %v\n", err)
		return
	}

	m.currentJournalAction = JournalAction(result)
}

func createJournal(name string) {
	newJournal, err := os.Create("./journals/" + name + ".md")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer newJournal.Close()

	fmt.Println("Das Journal wurde erfolgreich angelegt. Bitte gib einen ersten Eintrag ein")

	editJournal(name + ".md")
}

func editJournal(journal string) {
	path := "./journals/" + journal

	fmt.Println("Öffne das Journal im Editor...")

	editor := "vim"
	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Ein Fehler ist aufgetreten: ", err)
		return
	}

	fmt.Println("Dein Journal wurde erfolgreich gespeichert.")
}

func (m *Menu) addNewJournal() {
	prompt := promptui.Prompt{
		Label: "Titel des neuen Journals",
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Ein Fehler ist aufgetreten %v\n", err)
		return
	}

	createJournal(result)
}

func (m *Menu) selectExistingJournals() string {
	prompt := promptui.Select{
		Label: m.prompt,
		Items: getExistingJournals(),
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return result
}

func previewJournal(selectedJournal string, markdownPreviewer string) {
	path := "./journals/" + selectedJournal

	cmd := exec.Command(markdownPreviewer, path)

	cmd.Stdout = os.Stdout // Redirect Glow's output to the standard output
	cmd.Stderr = os.Stderr // Redirect errors to the standard error

	err := cmd.Run()

	if err != nil {
		fmt.Println("Wenn du möchtest, dass dein Journal in einem schönen Markdown-Viewer angezeigt wird, installiere bitte Glow.")
		previewJournal(selectedJournal, "cat")
	}
}

func getExistingJournals() []string {
	dirs, dirErr := os.ReadDir("./journals")

	if dirErr != nil {
		fmt.Println(dirErr)
		return nil
	}

	var choices []string

	for _, dir := range dirs {
		choices = append(choices, dir.Name())
	}

	return choices
}

func (m *Menu) deleteJournal() {
	selectedJournal := m.selectExistingJournals()

	prompt := promptui.Select{
		Label: "Bist du sicher, dass du die Datei: " + selectedJournal + " löschen möchtest? (ja/nein)",
		Items: []string{"ja", "nein"},
	}

	_, shouldDelete, err := prompt.Run()

	if err != nil {
		fmt.Printf("Ein Fehler ist aufgetreten %v\n", err)
	}

	if shouldDelete == "nein" {
		return
	}

	path := "./journals/" + selectedJournal

	deleteError := os.Remove(path)

	if deleteError != nil {
		fmt.Println("Der Eintrag konnte nicht gelöscht werden", err)
		return
	}

	fmt.Println("Der Eintrag wurde erfolgreich gelöscht.")
}

func triggerMenuAction(m Menu) {
	m.chooseJournalAction()

	if m.currentJournalAction == JournalActionEdit {
		selectedJournal := m.selectExistingJournals()
		editJournal(selectedJournal)
	}

	if m.currentJournalAction == JournalActionNew {
		m.addNewJournal()
	}

	if m.currentJournalAction == JournalActionPreview {
		selectedJournal := m.selectExistingJournals()

		previewJournal(selectedJournal, "glow")
	}

	if m.currentJournalAction == JournalActionDelete {
		m.deleteJournal()
	}

	if m.currentJournalAction == JournalActionExit {
		exitApp()
	}

	triggerMenuAction(m)
}

func exitApp() {
	os.Exit(0)
}

func InitMenu() {
	// Create a new menu
	// kind of uneccessary, but it's a good example of how i could use a struct
	menu := Menu{
		prompt: "Bitte wähle einen Eintrag aus:",
	}

	// Trigger the menu action
	triggerMenuAction(menu)
}
