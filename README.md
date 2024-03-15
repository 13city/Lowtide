
# Lowtide 🌊🔍 Dive into the Digital Depths

<p align="center">
<img src="https://your-logo-url-here.png" alt="Lowtide Logo" width="150"/>
</p>

**Embrace the cyberpunk spirit and navigate through the web's shadowy corridors.** 🛡️✨ Lowtide is not just a tool; it's your cybernetic augment for the digital age—a finely honed Go-powered engine designed for cybersecurity aficionados and digital defenders. With Lowtide, launch into cyberspace, penetrate the silence of the shadows, and illuminate the secrets that lie beneath.

## 🛠️ Features

- **⚡ Lightning Scans**: Unleash Go's concurrent might to slice through networks like a katana through silk.
- **🚩 Banner Reconnaissance**: Cast your net to catch elusive service banners, unveiling whispers of vulnerabilities.
- **🔓 Port Infiltration**: Pry open the digital doors across IP edifices and hidden subnets, charting your path in the neon grid.
- **📓 Chronicles of the Wire**: Every scan etches its tale into logs, crafting a saga of your digital conquests.
- **🌐 Subnet Expeditions**: Traverse vast subnets with stealth and precision, mapping the contours of digital domains.
- **🔧 Modularity Matrix**: Lowtide is engineered for evolution, inviting you to embed your code into its core.

## 🚀 Quick Start

Gear up, Runner! Before you jack in, let's refine your rig. Make sure Go is coursing through your system. Navigate to the `setup` directory in our repository for scripts that will streamline your environment setup, no matter your distro.

### Environment Setup

Equipping for the dive is straightforward with our tailored scripts:

- **Ubuntu Runners**:
  
  ```bash
  ./setup/setup_ubuntu.sh
  ```

- **Manjaro Nomads**:
  
  ```bash
  ./setup/setup_manjaro.sh
  ```

### Compiling Lowtide

The scripts will compile Lowtide, but should you ever need to invoke the compilation runes manually:

```bash
go build -o lowtide
```

Ensure your current directory is the root of Lowtide's domain, then summon the build.

## 🗺️ Usage

Lowtide accepts your commands both as whispers in the shell and through the arcane scripts of `config.json`.

### CLI Interface

Direct your scans with precision through the command-line interface:

```bash
./lowtide --startIP "192.0.2.1" --endIP "192.0.2.10" --ports "22,80,443" --timeout 1000
```

Channel the power to scan IP ranges, seeking out open ports with the patience of a shadow.

### Configuration Script (`config.json`)

For operations demanding detailed specifications, `config.json` is your grimoire:

```json
{
    "timeout": 1000,
    "ports": "80,443",
    "subnet": "192.168.1.0/24"
}
```

Invoke Lowtide with your configurations:

```bash
./lowtide -useConfig
```

### Configuration Parameters

- **timeout**: The span (in ms) Lowtide waits, listening for echoes from the void.
- **ports**: Designate your targets, be it a lone sentinel, a list, or an army ranged.
- **startIP**/**endIP** (CLI): Define the bounds of your digital odyssey.
- **subnet** (Config file): Chart your course through the subnet's expanse in CIDR.

## 📁 Logs

Lowtide engraves the stories of each expedition into the `./logging` directory, ensuring no detail of your digital dominion is lost to the ether.

## 🤝 Contribute

Aid in the evolution of Lowtide. Craft your enhancements, and through the rite of a pull request, merge your vision with ours.

## 🆘 Support

Lost among the wires or seeking knowledge? Open a beacon in our [GitHub repository](https://github.com/13city/Lowtide/issues), and guidance will find you.

---

<p align="center">
<i>With Lowtide, the digital frontier is yours to explore—beyond the veil of neon lies untold power.</i>
</p>

---
