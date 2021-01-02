# procman

`procman`, the process manager, helps manage the execution of long running processes.

# Installation

Download your system's binary from the [latest release](https://github.com/bantmen/procman/releases/latest). Releases are done automatically on every merge.

# Usage

```shell 
procman -h
```

The simplest invocation ensures that your command runs forever:

```shell 
procman ls -lh
```

You can enforce restarts based on the process status:

```shell 
procman -mem 50 python some_file.py
```

# Supported systems

Linux, macOS, and Raspberry Pi.

Windows is currently not supported.
