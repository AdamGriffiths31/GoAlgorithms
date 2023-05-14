# GoAlgorithms

GoAlgorithms is a project that generates blank algorithm templates with test cases in Go, organized in folders based on the current date.

## Program Overview

The GoAlgorithms program generates a folder structure with the current date (e.g., "2023-05-13") as the root folder. Inside this folder, it creates a set of algorithm templates, each consisting of a Go file and a test file.

The generated template files are named based on the algorithm or problem they represent. You can modify these templates to implement your own algorithm solutions and run the corresponding tests.

```
go test .\2023-05-13/... -v
```

## Running the Program

Clone the GoAlgorithms repository to your local machine.

```bash
go build
```

```bash
./GoAlgorithms
```

After running the program, you will find a new folder named with today's date (e.g., "2023-05-13") in the project's root directory. Inside this folder, you'll find the generated algorithm templates.
