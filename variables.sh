#!/bin/bash

export AppVersion=${AppVersion:-$(git tag --sort=-version:refname | head -n 1)}