# Lowtide - Reconnaissance Tool

<p align="center">
  <img src="./images/lowtide.jpg" alt="Lowtide Logo" width="600"/>
</p>

## **👨‍💻 Greetings, Cyber Seeker!**

Welcome to **Lowtide**, your trusty cyber reconnaissance companion designed to navigate the digital depths with precision and stealth. Crafted using Go, Lowtide empowers you to unveil hidden digital landscapes, expose vulnerabilities, and traverse networks with finesse and speed. 🚀💻

## 🛠️ Features

- **⚡️ Lightning-Fast Scans**: Utilizing Go's concurrency, scan thousands of ports in a fraction of the time.
- **🔍 Targeted Banner Analysis**: Extract banners to identify services and pinpoint potential weaknesses.
- **🔓 Port Decryption**: Uncover the secrets behind each port with unparalleled accuracy.
- **📝 Real-Time Digital Logs**: Chronicle every detail of your cyber exploration in real-time.
- **🌐 Subnet Exploration**: Embark on expeditions to map out uncharted digital territories.
- **🔧 Hackable Architecture**: Lowtide's modular design invites customization and adaptation to suit your needs.

## 🚀 Setup

Prepare your arsenal for a hacking extravaganza. Ensure Go is installed, then deploy our setup script for a seamless configuration.

### Configuration Script

Initiate your cyber toolkit effortlessly:

```bash
./setup.sh
```

This script primes your environment, preparing Lowtide for its inaugural voyage.

### Manual Compilation

Prefer a manual approach? Fear not:

```bash
go build -o Lowtide
```

Navigate to Lowtide's root directory and commence the compilation process.

## 🕹️ Deployment

Unleash Lowtide's prowess either through the shadows of the terminal or via the `config.json`.

### CLI Invocation

Command Lowtide's abilities directly from the terminal for immediate action:

```bash
./Lowtide --startIP "10.10.0.1" --endIP "10.10.0.254" --ports "22,80,443" --timeout 1000
```

### The `config.json` Codex

Alternatively, let `config.json` serve as your guide in defining your digital raid parameters:

```json
{
    "timeout": 1000,
    "ports": "22,80,443",
    "subnet": "10.10.0.0/24"
}
```

Activate Lowtide with your tailored configuration:

```bash
./Lowtide -useConfig
```

## 📚 Log Scrolls

Delve into `./logs` to witness the chronicles of your digital dominance, where every conquest is meticulously documented.

## 🛠 Contribute

Join the ranks of the Lowtide legion. Enhance its capabilities, share your expertise through pull requests, and leave your mark on the annals of cyber history.

## 🆘 Support

Encountered a digital phantom or in need of esoteric knowledge? Send a signal flare to our [GitHub issues](https://github.com/yourgithub/Lowtide/issues).

---

<p align="center">
<i>Embrace Lowtide. Your odyssey into the depths of cyberspace commences now.</i> 🌐👤
</p>
