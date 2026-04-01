// --- Free-and-Fast-Packaging Turbo Logic ---
const REPO = "SparkShardMC/Free-and-Fast-Packaging";

async function downloadLatest(osType) {
    const statusBox = document.getElementById('status-box');
    if (!statusBox) return; // Safety check

    statusBox.style.display = "inline-block";
    statusBox.style.color = "#39ff14"; // Neon Green
    statusBox.innerText = `⚡ Requesting ${osType} build from Cloud Engine...`;

    try {
        // 1. Fetch the latest release data
        const response = await fetch(`https://api.github.com/repos/${REPO}/releases/latest`);
        
        if (!response.ok) {
            throw new Error("Release not found");
        }

        const data = await response.json();
        
        // 2. Flexible Asset Matching
        // Looks for "Win", "Mac", or "Linux" regardless of casing or file extension
        const asset = data.assets.find(a => 
            a.name.toLowerCase().includes(osType.toLowerCase().substring(0, 3))
        );

        if (asset) {
            statusBox.innerText = "🚀 Connection verified! Downloading...";
            
            // Trigger download
            const link = document.createElement('a');
            link.href = asset.browser_download_url;
            link.download = asset.name;
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);
            
            setTimeout(() => {
                statusBox.style.display = "none";
            }, 5000);
        } else {
            // This triggers if the GitHub Action is still spinning
            statusBox.innerText = `⚙️ Cloud Engine is still packaging ${osType}. Refresh in 30s.`;
            statusBox.style.color = "#00ffff"; // Cyan for "Processing"
        }
    } catch (error) {
        console.error(error);
        statusBox.innerText = "⚠️ Server Busy: Try again or check GitHub Releases.";
        statusBox.style.color = "#ff4444";
    }
}

// --- Auto-Detection System ---
window.addEventListener('load', () => {
    const platform = navigator.userAgent.toLowerCase();
    const btns = {
        win: document.getElementById('win-btn'),
        mac: document.getElementById('mac-btn'),
        linux: document.getElementById('linux-btn')
    };
    
    let key = "";
    if (platform.includes("win")) key = "win";
    else if (platform.includes("mac")) key = "mac";
    else if (platform.includes("linux")) key = "linux";

    if (key && btns[key]) {
        btns[key].classList.add('recommended');
        const badge = document.createElement('div');
        badge.style.cssText = "font-size: 0.65rem; color: #39ff14; margin-top: 5px; letter-spacing: 1px; font-weight: bold;";
        badge.innerText = "BEST FOR YOUR SYSTEM";
        btns[key].appendChild(badge);
    }
});
