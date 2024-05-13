package main

import (
    "fmt"
    "bufio"
    "os"
    // "strings"
)

type Task struct {
    Description string
    Completed   bool
}

var tasks []Task

func AddTask(description string) {
    tasks = append(tasks, Task{Description: description, Completed: false})
}

func CompleteTask(index int) {
    if index >= 0 && index < len(tasks) {
        tasks[index].Completed = true
    } else {
        fmt.Println("Invalid task number.")
    }
}

func ListTasks() {
    for i, task := range tasks {
        status := "Pending"
        if task.Completed {
            status = "Completed"
        }
        fmt.Printf("%d. %s [%s]\n", i+1, task.Description, status)
    }
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("\nChoose an option:")
        fmt.Println("1. Add Task")
        fmt.Println("2. Complete Task")
        fmt.Println("3. List Tasks")
        fmt.Print("Enter option: ")

        scanner.Scan()
        option := scanner.Text()

        switch option {
        case "1":
            fmt.Print("Enter task description: ")
            scanner.Scan()
            description := scanner.Text()
            AddTask(description)
        case "2":
            fmt.Print("Enter task number to complete: ")
            scanner.Scan()
            index := 0
            fmt.Sscanf(scanner.Text(), "%d", &index)
            CompleteTask(index - 1)  // Adjusting index for zero-based slice index
        case "3":
            ListTasks()
        default:
            fmt.Println("Invalid option, please choose again.")
        }
    }
}