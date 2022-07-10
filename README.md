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

## Usage
#### Address with CIDR
```Bash
snet -c 192.168.0.1/24
```

#### Address with Subnet Mask
```Bash
snet -a 192.168.0.1 -m 255.255.255.0
```


#### Output
```Bash
NETWORK      USEABLE          BROADCAST      MASK           CIDR  
192.168.0.0  192.168.0.1-254  192.168.0.255  255.255.255.0  /24
```


