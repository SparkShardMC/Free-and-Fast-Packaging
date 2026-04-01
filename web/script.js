// --- Free-and-Fast-Packaging Official Logic ---
const REPO = "SparkShardMC/Free-and-Fast-Packaging";

async function downloadLatest(osType) {
    const statusBox = document.getElementById('status-box');
    statusBox.style.display = "inline-block";
    statusBox.innerText = `Preparing ${osType} download...`;

    try {
        // Fetch the latest release data from GitHub API
        const response = await fetch(`https://api.github.com/repos/${REPO}/releases/latest`);
        
        if (!response.ok) {
            throw new Error("Release not found");
        }

        const data = await response.json();
        
        // Match the asset name based on the release.yml naming convention
        // (Windows.exe, Mac.dmg, or Linux)
        const asset = data.assets.find(a => 
            a.name.toLowerCase().includes(osType.toLowerCase())
        );

        if (asset) {
            statusBox.innerText = "Connection successful. Starting download...";
            // Trigger the native browser "Save As" interface
            window.location.href = asset.browser_download_url;
            
            // Hide status after a few seconds
            setTimeout(() => {
                statusBox.style.display = "none";
            }, 5000);
        } else {
            statusBox.innerText = `Error: Build for ${osType} is still compiling. Check back in 1 minute.`;
            statusBox.style.color = "#ff4444";
        }
    } catch (error) {
        console.error(error);
        statusBox.innerText = "Update Server Offline: Please check GitHub Releases manually.";
        statusBox.style.color = "#ff4444";
    }
}

// --- Auto-Detection & Recommendation System ---
window.addEventListener('load', () => {
    const platform = navigator.userAgent.toLowerCase();
    const winBtn = document.getElementById('win-btn');
    const macBtn = document.getElementById('mac-btn');
    const linuxBtn = document.getElementById('linux-btn');
    
    let detectedBtn = null;

    // Detect OS and assign the 'recommended' class for the CSS neon pulse
    if (platform.includes("win")) {
        detectedBtn = winBtn;
    } else if (platform.includes("mac")) {
        detectedBtn = macBtn;
    } else if (platform.includes("linux")) {
        detectedBtn = linuxBtn;
    }

    if (detectedBtn) {
        detectedBtn.classList.add('recommended');
        // Add a small "Recommended" badge text to the button
        const badge = document.createElement('span');
        badge.style.fontSize = "0.7rem";
        badge.style.display = "block";
        badge.style.color = "#39ff14";
        badge.innerText = "MATCHED FOR YOUR OS";
        detectedBtn.appendChild(badge);
    }
});
