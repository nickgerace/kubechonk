# kubechonk

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Build Status](https://github.com/nickgerace/kubechonk/workflows/Build/badge.svg)](https://github.com/nickgerace/kubechonk/actions?query=workflow%3ABuild)
[![Go Report Card](https://goreportcard.com/badge/github.com/nickgerace/kubechonk)](https://goreportcard.com/report/github.com/nickgerace/kubechon://goreportcard.com/report/github.com/nickgerace/kubechonk)

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

You can change the install destination in the Makefile, or use the default.

```bash
git clone https://github.com/nickgerace/kubechonk.git
cd kubechonk
make build
sudo make install
```

Note: the default location is ```/usr/local/bin/kubectl-chonk```.

## Uninstallation

Remove the binary to fully uninstall the plugin.
If you installed to destination different the default, you may have to manually remove the binary.

```bash
sudo make uninstall
```

Note: the default location is ```/usr/local/bin/kubectl-chonk```.

## Credits

- [Ahmet Alp Balkan](https://ahmet.im/)'s YouTube series on [kubectl plugins](https://www.youtube.com/watch?v=_W2qZvQT6XY)
- The official [sample-cli-plugin](https://github.com/kubernetes/sample-cli-plugin) repository as the basis for this repository
