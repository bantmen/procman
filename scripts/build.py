#!/usr/bin/python3

import os

BINS_FOLDER = "bins"

os.mkdir(BINS_FOLDER)


def build(goos, goarch, goarm=None):
    assert goarm is None or goarch == "arm"
    output_name = "procman_{}_{}".format(goos, goarch)
    if goarm is not None:
        output_name += "v{}".format(goarm)
    if goos == "windows":
        output_name += ".exe"

    env = "GOOS={} GOARCH={}".format(goos, goarch)
    if goarm:
        env += " GOARM={}".format(goarm)

    cmd = "{} go build -ldflags '-s -w' -o {}".format(
        env, os.path.join(BINS_FOLDER, output_name)
    )
    print("Running", cmd)
    os.system(cmd)


# raspberry pi
build(goos="linux", goarch="arm", goarm="5")

# linux
build(goos="linux", goarch="amd64")

# macOS
build(goos="darwin", goarch="amd64")
