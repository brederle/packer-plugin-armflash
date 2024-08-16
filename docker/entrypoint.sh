#!/usr/bin/env bash
mountpoint -q /proc/sys/fs/binfmt_misc || mount binfmt_misc -t binfmt_misc /proc/sys/fs/binfmt_misc

# meanwhile, on debian, the required magic identifier are already available after mount

# allow any command to be executed
exec "$@"
