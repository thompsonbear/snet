# :eight_spoked_asterisk: snet
Simple, easy to use, CLI based subnetting calculator/tool built in Go

![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/thompsonbear/snet?include_prereleases)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/thompsonbear/snet)

## Installation
### Linux
```Bash
curl -LJO https://github.com/thompsonbear/snet/releases/download/v0.1.0-alpha/snet
chmod +x snet
sudo mv snet /usr/bin
```
### Windows
1. [Download snet.exe](https://github.com/thompsonbear/snet/releases/download/v0.1.0-alpha/snet.exe)
2. Add %USERPROFILE%\go\bin to the path environment variable **(Can skip if you have Go installed)**
    * Run sysdm.cpl
    * Advanced Tab > Environment Variables...
    * Select variable "path" and select Edit
    * Add "%USERPROFILE%\go\bin" to the list
3. Move snet.exe to %USERPROFILE%\go\bin **(You may need to create this directory)**

## Usage
#### Address with CIDR
```Bash
snet -c 192.168.0.1/24
```
#### Address with Subnet Mask
```Bash
snet -a 192.168.0.1 -m 255.255.255.0
```
#### **Output**
```Bash
NETWORK      USEABLE          BROADCAST      MASK           CIDR  
192.168.0.0  192.168.0.1-254  192.168.0.255  255.255.255.0  /24
```
## Planned Features
:green_circle:  Basic IPv4 Subnet Calculator

:black_circle: List all neighboring networks with the specified CIDR

:black_circle: Optional configuration file to add and reorder output columns

:black_circle: Split the specified address block into a specified amount of subnets

:black_circle: IPv6 Support :neckbeard:


