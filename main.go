package main

func main() {
    // 1. Initialize the UI (from ui.go)
    // We pass the version and a title
    ShowUpdateUI("1.0.0", "Initial Release", func() {
        // This is what happens when the user clicks 'Update/Start'
        LaunchScanner() 
    })
}

// LaunchScanner bridges the UI to the scanner.go logic
func LaunchScanner() {
    files, err := ScanForJavaFiles("./")
    if err != nil {
        println("Error scanning:", err.Error())
        return
    }
    println("Found files to package:", len(files))
}
