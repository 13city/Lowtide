# Lowtide - Cyber Reconnaissance Arsenal

<p align="center">
  <img src="./images/lowtide.jpg" alt="Lowtide Logo" width="600"/>
</p>

## **👨‍💻 Greetings, Cyber Seeker!**

Welcome to **Lowtide**, your dark ally in the vast digital expanse. Forged in the depths of cyber obscurity using the power of Go, Lowtide is your weapon of choice for uncovering hidden digital landscapes, exploiting vulnerabilities, and mastering the art of network traversal with precision and stealth. 🚀💻

## 🛠️ Features

- **⚡️ Lightning-Fast Scans**: Harnessing Go's concurrency, Lowtide swiftly scans thousands of ports, leaving no digital stone unturned.
- **🔍 Targeted Banner Analysis**: Extracting banners to unveil services and exploit potential weaknesses is Lowtide's specialty.
- **🔓 Port Decryption**: Lowtide's unparalleled accuracy unlocks the secrets concealed within each port, empowering you with knowledge.
- **📝 Real-Time Digital Logs**: Every maneuver is meticulously chronicled in real-time, providing you with a detailed roadmap of your cyber conquests.
- **🌐 Subnet Exploration**: Embark on daring expeditions to map out uncharted digital territories and claim them as your own.
- **🔧 Hackable Architecture**: Lowtide's modular design beckons you to bend and mold its capabilities to your will, shaping it into the ultimate cyber weapon.

## 🚀 Setup

Prepare your arsenal for a hacking odyssey. Ensure Go is installed, then deploy our setup script to initiate Lowtide for its maiden voyage.

### Configuration Script

Empower your cyber toolkit effortlessly:

```bash
./setup.sh
```

This script primes your environment, preparing Lowtide for its inaugural journey into the dark depths of cyberspace.

### Manual Compilation

Prefer to tread the path manually? Fear not:

```bash
go build -o Lowtide
```

Navigate to Lowtide's lair and commence the compilation process.

## 🕹️ Deployment

Unleash Lowtide's dark powers either through the shadows of the terminal or via the `config.json`.

### CLI Invocation

Summon Lowtide's abilities directly from the terminal for immediate action:

```bash
./Lowtide --startIP "10.10.0.1" --endIP "10.10.0.254" --ports "22,80,443" --timeout 1000
```

### The `config.json` Codex

Alternatively, let `config.json` serve as your oracle in defining your digital raid parameters:

```json
{
    "timeout": 1000,
    "ports": "22,80,443",
    "subnet": "10.10.0.0/24"
}
```

Initiate Lowtide with your custom incantations:

```bash
./Lowtide -useConfig
```

## 📚 Logging

Delve into the abyss of `logging` to witness the saga of your digital conquests, where each victory is etched into the annals of cyber history.

## 🛠 Contribute

Join the ranks of the Lowtide legion. Enhance its dark arsenal, share your forbidden knowledge through pull requests, and solidify your legacy in the shadows of cyber infamy.

## 🆘 Support

Encountered a digital specter or in need of arcane wisdom? Signal your distress to our [GitHub issues](https://github.com/yourgithub/Lowtide/issues).

---

<p align="center">
<i>Embrace Lowtide. Your descent into the abyss of cyberspace begins now.</i> 🌐👤
</p>