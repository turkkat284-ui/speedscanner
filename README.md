SpeedScanner 🚀
A high-performance, cross-platform TCP port scanner written in Go (Golang). Leveraging the power of goroutines and sync.WaitGroup, it scans network ports concurrently with extreme speed.

✨ Features
Platform Independent: Works seamlessly on Windows, Linux, and macOS thanks to Go's runtime.

High Concurrency: Scans multiple ports simultaneously using Go's lightweight threads (Goroutines).

Smart Timeout: Implements net.DialTimeout (2 seconds) to prevent the application from hanging on unresponsive ports.

Zero Dependencies: Built entirely with Go's standard library. No external packages required.

🚀 Getting Started
Prerequisites
Make sure you have Go installed on your system.

Running the App
Clone or download the source code.

Open your terminal and navigate to the project folder.

Run the following command:

Bash

go run main.go
Enter the target IP address when prompted (e.g., 127.0.0.1 or scanme.nmap.org).

📦 Compilation (Build for any OS)
You can compile this code into a standalone executable for any operating system:

For Windows:

Bash

GOOS=windows GOARCH=amd64 go build -o scanner.exe
For Linux:

Bash

GOOS=linux GOARCH=amd64 go build -o scanner
For macOS:

Bash

GOOS=darwin GOARCH=amd64 go build -o scanner
🛠 How It Works
The scanner iterates through ports 1 to 1024. For each port, it spawns a goroutine that attempts a TCP connection. If the connection is established within the timeout period, the port is identified as Open.

⚖️ Legal Disclaimer
This tool is for educational and ethical testing purposes only. Scanning targets without prior authorization is illegal. The developer is not responsible for any misuse or damage caused by this program.
