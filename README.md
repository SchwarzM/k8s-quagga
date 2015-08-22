# k8s-quagga

Alpine linux running quagga zebra and ospfd.
Config done by small go binary that translates environment to files

## Usage ##
### Global flags ###
```K8S_QUAGGA_OUTPUT``` - Directory to put files to defaults to ```/etc/quagga```

```K8S_QUAGGA_PASSWORD``` - Password to put in config defaults to ```changeme```

### Subcommands ###
```zebra``` - Print out zebra.conf

```ospfd``` - Print out ospfd.conf

### Ospfd Flags ###
```K8S_QUAGGA_INTERFACE``` - Interface to use

```K8S_QUAGGA_ROUTERID``` - RouterId to use

```K8S_QUAGGA_PORTALNET``` - Net 1 to announce

```K8S_QUAGGA_CONTAINERNET``` - Net 2 to announce
