package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	// The Web Listener
	http.HandleFunc("/compile", handleCompileRequest)
	
	port := "8080"
	fmt.Printf("🚀 SparkShard Engine LIVE on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}

func handleCompileRequest(w http.ResponseWriter, r *http.Request) {
	// 1. Get OS from the URL (e.g., /compile?os=windows)
	targetOS := r.URL.Query().Get("os")
	if targetOS == "" {
		targetOS = runtime.GOOS // Default to server OS
	}

	// 2. Generate the "Awesome" Filename
	timestamp := time.Now().Format("01-02-1504") // e.g., 04-20-1337
	fileName := fmt.Sprintf("FreeFast_Packager_%s_%s", targetOS, timestamp)
	
	extension := ""
	if targetOS == "windows" {
		extension = ".exe"
	} else if targetOS == "darwin" {
		extension = ".app"
	}

	fullFileName := fileName + extension

	// 3. The "Lightning" Compile
	fmt.Printf("🛠️  Compiling %s for user...\n", fullFileName)
	
	cmd := exec.Command("go", "build", "-o", fullFileName, "main.go", "scanner.go", "packager.go", "ui.go")
	
	// Cross-compile magic
	cmd.Env = append(os.Environ(), "GOOS="+targetOS, "GOARCH=amd64")

	err := cmd.Run()
	if err != nil {
		http.Error(w, "Compile Error: "+err.Error(), 500)
		return
	}

	// 4. Send the file back to the browser immediately
	w.Header().Set("Content-Disposition", "attachment; filename="+fullFileName)
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, fullFileName)

	// 5. Cleanup (Delete the file after sending to keep the server fast)
	go func() {
		time.Sleep(10 * time.Second)
		os.Remove(fullFileName)
	}()
}
