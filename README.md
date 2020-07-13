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
NODE                         RESOURCE   VALUE
gke-cluster-cpu2mem1-alpha   cpu        2
gke-cluster-cpu2mem1-beta    cpu        2
gke-cluster-cpu1mem3-alpha   memory     3785940Ki
gke-cluster-cpu1mem3-beta    memory     3785940Ki
```

## Getting Started

You can install (and uninstall) this plugin by using [krew](https://krew.sigs.k8s.io/), or by downloading the binary manually.

### Krew

Clone this repository and run the following.

```bash
kubectl krew install --manifest=.krew.yaml
```

You can uninstall the plugin with ```krew``` as well.

```bash
kubectl krew uninstall chonk
```

### Manual

Download the [latest GitHub release](https://github.com/nickgerace/kubechonk/releases/latest) for your operating system and architecture.
Move the binary to your ```PATH``` to get started.

To completely uninstall the plugin, delete the binary from your ```PATH```.

## Credits

- [Ahmet Alp Balkan](https://ahmet.im/)'s YouTube series on [kubectl plugins](https://www.youtube.com/watch?v=_W2qZvQT6XY)
- The official [sample-cli-plugin](https://github.com/kubernetes/sample-cli-plugin) repository as the basis for this repository
