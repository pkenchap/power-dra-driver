# Power Dynamic Resource Allocation Driver
This repository contains a Power Architecture resource driver for use with the Dynamic Resource Allocation (DRA) feature of Kubernetes.

The driver facilitates access to:
1. Nest Accelerator unit (NX)are non-core components of the POWER processor chip including compression and encryption co-processors. One feature is nx-gzip, which is a DEFLATE compliant (RFC1950, 1951, 1952) compression accelerator in NX. This feature is a device driver in userspace `/dev/crypto/nx-gzip` which is shared among many LPARs and many containers. 

As additional features are made available, the driver will be expanded.

This project is licensed under the Apache 2.0 License.