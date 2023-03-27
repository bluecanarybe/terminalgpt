# TerminalGPT

## Description

ChatGPT from the terminal

### Installation

* `git clone https://github.com/bluecanarybe/terminalgpt.git`
* Add your openAI api-key
* Build with `go build -o chatgpt`
* Add the compiled binary as an alias: `echo "alias chatgpt='$PWD/chatgpt'" >> ~/.zshrc`
* Restart your shell: `source ~/.zshrc`

### Usage

```
$ chatgpt "question?"
```

#### Example:

```
$ chatgpt "how much is 1+1?"
```
