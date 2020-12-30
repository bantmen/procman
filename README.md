# procman

`procman`, the process manager, helps manage the execution of long running processes.

# Usage

The simplest invocation ensures that your command runs forever:

```shell 
procman ls -lh
```

You can enforce restarts based on the process status:

```shell 
procman -mem 50 python some_file.py
```

# Supported platforms

Linux, macOS.

Windows is currently not supported.
