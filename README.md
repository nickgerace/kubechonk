# kubechonk

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Build Status](https://github.com/nickgerace/kubechonk/workflows/Build/badge.svg)](https://github.com/nickgerace/kubechonk/actions?query=workflow%3ABuild)
[![Go Report Card](https://goreportcard.com/badge/github.com/nickgerace/kubechonk)](https://goreportcard.com/report/github.com/nickgerace/kubechonk)

Find the "chonkiest" nodes with this kubectl plugin.

## Usage

This plugin provides a single command to interact with your Kubernetes cluster.
The ```chonk``` command returns all the node(s) with the highest number of CPU cores, and all the node(s) with the largest amount of memory.

```bash
[user@hostname:~]
% kubectl chonk
NODE                    RESOURCE        VALUE
kind-control-plane      cpu             8
kind-control-plane      mem             12911996Ki
```

## Installation

Download the [latest GitHub release](https://github.com/nickgerace/kubechonk/releases/latest) for your operating system and architecture.
Move the binary to your ```PATH``` to get started.

## Uninstallation

Delete the binary from your ```PATH``` to completely uninstall the plugin.

## Credits

- [Ahmet Alp Balkan](https://ahmet.im/)'s YouTube series on [kubectl plugins](https://www.youtube.com/watch?v=_W2qZvQT6XY)
- The official [sample-cli-plugin](https://github.com/kubernetes/sample-cli-plugin) repository as the basis for this repository
