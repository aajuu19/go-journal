# Small CLI Journal App written in go

As I am learning go, I wanted to play around with the language and the file system. I decided to create a small CLI journal app. This is actually the first go app that i am writing. The app is very simple and has the following features:

- Add a journal entry as an md file using vim
- Edit a journal entry using vim
- Feature which allows you to select the files from a list in the cli
- preview a journal entry in the CLI using glow

## Usage

Run the following command to add a journal entry

From the root folder of the project run the following command

```bash
    go run .
```

## Required Dependencies

Since I wanted to preview the journal entries in the CLI, I used the following dependencies to make it look better. This dependencies need to be installed in order to preview the journal entries visually appealing. If not installed, the journal entries will be displayed as plain text.

- [Glow](https://github.com/charmbracelet/glow)
