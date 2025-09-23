# go-cli-to-do

Simple and lightweight command-line to-do list application written in Go.

Manage your daily tasks directly from the terminal, add, list, complete, and delete tasks quickly and efficiently.

## Features
- Add tasks
- List all tasks
- Mark tasks as done
- Delete tasks
- View task details
- SQLite storage
- Fast, simple and easy to use

## Installation
1. Clone the repository
```bash
git clone https://github.com/goretante/go-cli-to-do.git
cd go-cli-to-do
```
2. Build the binary:
```bash
go build -o go-cli-to-do
```
3. Run the CLI:
```bash
./go-cli-to-do
```

## Usage
```bash
# Add a new task
./go-cli-to-do add "Buy groceries"

# List tasks
./go-cli-to-do list

# Mark task as completed
./go-cli-to-do done 1

# View task details
./go-cli-to-do info 1

# Delete a task
./go-cli-to-do delete 2
```

## Improvements
- Add due dates.

## Built with
- Go
- Cobra
- SQLite
- GORM

## License
MIT license. See [LICENSE](https://github.com/goretante/go-cli-to-do/blob/main/LICENSE) for details.