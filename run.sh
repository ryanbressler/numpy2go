#!/bin/bash

go build -buildmode=c-shared -o numpy2go.so
python numpy2go.py
