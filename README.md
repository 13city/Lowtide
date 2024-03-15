# Lowtide 🌌🔒 - Digital Recon Tool

<p align="center">
<img src="./images/lowtide.jpg" alt="Lowtide Logo" width="200"/>
</p>

**👾 Greetings, Hacker!** You've just discovered **Lowtide**, a cyber reconnaissance tool that dives deep into the digital abyss. Crafted with the precision of Go, this is your arsenal for uncovering hidden digital landscapes, exposing vulnerabilities, and navigating through networks with stealth and speed. 🚀💻

## 🛠️ Features

- **⚙️ Ultra-Fast Scans**: With Go's concurrency, scan thousands of ports in the blink of an eye.
- **🎯 Targeted Banner Sniping**: Extract banners to identify services and their potential weak spots.
- **🔑 Port Decoding**: Unlock the secrets behind every port with laser precision.
- **📖 Digital Chronicles**: Log every detail of your cyber exploration in real-time.
- **🌍 Subnet Safari**: Embark on subnet expeditions to map out uncharted digital territories.
- **🔧 Hackable Core**: Modular design invites you to tweak, twist, and transform Lowtide to your will.

## 🚀 Setup

Prep your rig for an epic hacking session. Ensure Go is installed, then unleash our setup script for a hassle-free setup.

### Configuration Script

Initiate your cyber toolkit with ease:

```bash
./setup.sh
```

This magic script preps your environment, readying Lowtide for its maiden voyage.

### Compiling Lowtide Manually

Need to compile by hand? No sweat:

```bash
go build -o Lowtide
```

Ensure you're in Lowtide's root directory and let the compilation begin.

## 🕹️ Deployment

Engage Lowtide either through the shadows of the terminal or with the `config.json`.

### CLI Invocation

Command Lowtide's powers directly from the terminal for instant action:

```bash
./Lowtide --startIP "10.10.0.1" --endIP "10.10.0.254" --ports "22,80,443" --timeout 1000
```

### The `config.json` Codex

Or, let `config.json` be your guide to specifying your digital raid parameters:

```json
{
    "timeout": 1000,
    "ports": "22,80,443",
    "subnet": "10.10.0.0/24"
}
```

Activate Lowtide with your bespoke config:

```bash
./Lowtide -useConfig
```

## 📚 Log Scrolls

Dive into `./logs` to witness the saga of your digital dominance, where every conquest is meticulously chronicled.

## 🛠 Contribute

Join the Lowtide legion. Enhance its capabilities, share your prowess through pull requests, and etch your code into the annals of cyber history.

## 🆘 Support

Encountered a cyber specter or in need of arcane knowledge? Cast a signal flare to our [GitHub issues](https://github.com/yourgithub/Lowtide/issues).

---

<p align="center">
<i>Embrace Lowtide. Your journey into the depths of cyberspace begins now.</i> 🌐👤
</p>

