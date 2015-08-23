#!/bin/sh

set -ex

# generate zebra conf
./k8s-quagga zebra --PortalNet ${K8S_QUAGGA_PORTALNET} --PortalGw ${K8S_QUAGGA_PORTALGW}

# generate ospfd conf
./k8s-quagga ospfd --Interface ${K8S_QUAGGA_INTERFACE} --RouterId ${K8S_QUAGGA_ROUTERID} --HomeNet ${K8S_QUAGGA_HOMENET} --PortalNet ${K8S_QUAGGA_PORTALNET} --ContainerNet ${K8S_QUAGGA_CONTAINERNET}

exec /usr/bin/supervisord -c /etc/supervisord.conf
