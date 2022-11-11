#!/usr/bin/env bash

wget https://github.com/nushell/nushell/releases/download/0.71.0/nu-0.71.0-x86_64-unknown-linux-gnu.tar.gz

tar -xzf nu-0.71.0-x86_64-unknown-linux-gnu.tar.gz

ln ./nu-0.71.0-x86_64-unknown-linux-gnu/nu /usr/local/bin/nu

nu