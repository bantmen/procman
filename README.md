# procman

`procman`, the process manager, helps manage the execution of long running processes.

# Installation

Download the relevant binary from the [latest release](https://github.com/bantmen/procman/releases/latest).

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

# Supported platforms

Linux, macOS, and Raspberry Pi.

Windows is currently not supported.
