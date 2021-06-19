# tcpshark
A simple tcp visualizer for sniffing tcp packets in cli

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

# Installation
```
git clone https://github.com/krishpranav/tcpshark
cd tcpshark
go get
go build main.go
./main
```

# Usage
```
$ tcpshark -h
NAME:
   tcpterm - tcpdump for human

USAGE:
   shark [global options] command [command options] [arguments...]

VERSION:
     1.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --interface value, -i value  If unspecified, use lowest numbered interface.
   --read value, -r value       Read packets from pcap file.
   --filter value, -f value     BPF Filter
   --debug                      debug mode.
   --help, -h                   show help
   --version, -v                print the version
```
