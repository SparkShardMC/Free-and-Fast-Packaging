package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const CurrentVersion = "1.0.0"
const RepoURL = "https://api.github.com/repos/SparkShardMC/Free-and-Fast-Packaging/releases/latest"

type GitHubRelease struct {
	TagName string `json:"tag_name"`
	Body    string `json:"body"`
	Assets  []struct {
		BrowserDownloadURL string `json:"browser_download_url"`
		Name               string `json:"name"`
	} `json:"assets"`
}

func CheckForUpdates() {
	fmt.Println("🔍 Checking SparkShardMC GitHub for updates...")

	resp, err := http.Get(RepoURL)
	if err != nil {
		fmt.Println("❌ Could not connect to update server.")
		return
	}
	defer resp.Body.Close()

	var release GitHubRelease
	json.NewDecoder(resp.Body).Decode(&release)

	if release.TagName != "v"+CurrentVersion {
		fmt.Printf("🚀 New Update Available: %s\n", release.TagName)
		fmt.Println("--- CHANGELOG ---")
		fmt.Println(release.Body)
		fmt.Println("-----------------")
		
		// This is where the Progress Bar triggers
		StartUpdateDownload(release.Assets.BrowserDownloadURL)
	} else {
		fmt.Println("✅ App is up to date. Ready to Package!")
	}
}

func StartUpdateDownload(url string) {
	fmt.Println("📦 Downloading Update...")
	// Logic for the "Filling Jar" Progress Bar
	// 1. Create a GET request to the download URL
	// 2. Track the 'Content-Length'
	// 3. Print [#####-----] 50% style or a custom GUI bar
}

func main() {
	CheckForUpdates()
	// After update check, launch the Dropper UI
	// LaunchScanner()
}
