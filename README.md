# LowTide 🌊 Go for shells

**Step into the digital shadows and unveil the secrets hidden within the web's vast expanse.** LowTide is the quintessential toolkit for cybersecurity enthusiasts and network defenders, engineered with Go's prowess for speed and efficiency. It's crafted to conduct deep scans, expose vulnerabilities, and map out the digital landscape, all while maintaining stealth.

## 🛠 Features

- **Rapid Reconnaissance**: Harness the power of Go's concurrency for blazing-fast network scans.
- **Banner Grabbing**: Probe services to reveal banners that hint at potential vulnerabilities.
- **Port Scanning**: Identify open ports across specific IP addresses or vast subnets, marking gateways to deeper insights.
- **Detailed Logging**: Capture comprehensive logs of each scan, aiding in meticulous analysis.
- **Subnet Support**: Execute thorough scans across entire subnets to evaluate network security postures.
- **Extensibility**: Architectured for seamless integration of additional scanning capabilities and features.

## 🚀 Quick Start

Before diving in, ensure Go is correctly installed and configured on your machine. Follow the setup scripts provided within the project's `setup` directory to streamline your environment setup.

### Setting Up Your Environment

Initiate your setup with dedicated scripts tailored for Ubuntu and Manjaro environments:

- **For Ubuntu Users**: Execute the setup script to install Go, clone LowTide, and build the project.

```bash
./setup/setup_ubuntu.sh
```

- **For Manjaro Users**: Similar to Ubuntu, this script prepares your system, installs Go, clones LowTide, and compiles the executable.

```bash
./setup/setup_manjaro.sh
```

### Building LowTide

The setup scripts automatically compile LowTide. However, should you need to manually build or rebuild:

```bash
go build -o lowtide
```

Navigate to your project's root directory and execute the above command to compile LowTide into a ready-to-run executable.

## 🗺 Usage

Deploy LowTide using either command-line arguments for quick scans or the `config.json` file for customized scanning operations.

### Command-Line Interface (CLI)

- **Conduct an IP Range and Port Scan**:

```bash
./lowtide --startIP "192.0.2.1" --endIP "192.0.2.10" --ports "22,80,443" --timeout 1000
```

This command scans the specified IP range for open ports 22, 80, and 443, each with a 1-second timeout.

### Configuration File (`config.json`)

For more detailed or specific scanning configurations, utilize `config.json`:

- **Sample `config.json` Configuration**:

```json
{
    "timeout": 1000,
    "ports": "80,443",
    "subnet": "192.168.1.0/24"
}
```

This setup directs LowTide to scan the 192.168.1.0/24 subnet for ports 80 and 443.

Activate LowTide with the custom configuration:

```bash
./lowtide -useConfig
```

### Configuration Parameters

- **timeout**: The maximum duration (milliseconds) to wait for a response from each target.
- **ports**: Target ports for scanning. Specify individual ports, comma-separated lists, or ranges.
- **startIP** and **endIP** (CLI): The IP range for the scan.
- **subnet** (Config file): The CIDR notation subnet for scanning.

## 📁 Logging

Logs from each scan are meticulously recorded, detailing open ports and banner information. Stored in the `./logging` directory, these logs are timestamped for post-operation analysis and review.

## 🤝 Join the Mission

Enhance LowTide's capabilities by contributing to its development. Fork [13city/lowtide](https://github.com/13city/lowtide), integrate your innovative features, and submit a pull request to share your advancements.

## 🆘 Support

Encountering challenges or have inquiries? Open an issue in the [GitHub repository](https://github.com/13city/lowtide/issues) for support and guidance.

---

**LowTide** equips you with the tools to navigate the digital frontier, revealing vulnerabilities and securing the cyber landscape. Begin your journey into the depths of cybersecurity.

---