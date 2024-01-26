
# Note Taker

A simple note taking application written to help me learn Golang


## Features

- Adding a new note
- Listing notes
- Deleting notes
- [WIP] Searching through notes


## Installation

#### Pre-requisites ####

- Golang >= 1.21.5
- Make   >= 4.2.1
- Vim

#### Quick start to using the CLI ####

```bash
  # download dependancies
  go get -d ./...
  # install the app
  make install
  note-taker -h
```
#### Build the app to run locally without adding to path ####

```bash
    make build
```
    
## Usage/Examples

#### Quick example of using the CLI to create a note called "Note" using a short inline message and listing off all the notes you've created

```bash
  note-taker new --name "Note" --message "Just a quick note to myself"
  note-taker list --limit 40 
```


## Authors

- [@l-baston](https://www.github.com/l-baston)

