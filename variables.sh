#!/bin/bash

export AppVersion=`git tag --sort=-version:refname | head -n 2`