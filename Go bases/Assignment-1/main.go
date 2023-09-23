package main

import (
    "fmt"
    "yourcompany/employee/manager"
    "yourcompany/employee/developer"
    "yourcompany/employee/sales"
    "yourcompany/employee/marketing"
    "yourcompany/employee/hr"
)

func main() {
    // Create instances of different employees
    manager := manager.NewManager("Manager", 75000.0, "123 Main St")
    developer := developer.NewDeveloper("Developer", 60000.0, "456 Elm St")
    sales := sales.NewSales("Sales", 80000.0, "789 Oak St")
    marketing := marketing.NewMarketing("Marketing", 70000.0, "101 Pine St")
    hr := hr.NewHR("HR", 65000.0, "202 Cedar St")

    // Access and print employee details
    fmt.Println("Manager Position:", manager.GetPosition())
    fmt.Println("Developer Salary:", developer.GetSalary())
    fmt.Println("Sales Address:", sales.GetAddress())
    fmt.Println("Marketing Position:", marketing.GetPosition())
    fmt.Println("HR Salary:", hr.GetSalary())

    // Modify employee details
    manager.SetSalary(80000.0)
    developer.SetPosition("Senior Developer")
    sales.SetAddress("999 Maple St")
    marketing.SetSalary(75000.0)
    hr.SetPosition("HR Manager")

    // Print updated employee details
    fmt.Println("Updated Manager Salary:", manager.GetSalary())
    fmt.Println("Updated Developer Position:", developer.GetPosition())
    fmt.Println("Updated Sales Address:", sales.GetAddress())
    fmt.Println("Updated Marketing Salary:", marketing.GetSalary())
    fmt.Println("Updated HR Position:", hr.GetPosition())
}