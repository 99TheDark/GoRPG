#!/bin/bash
go build
cp -R ./res ./build/GoRPG.app/Contents/Resources
cp ./game ./build/GoRPG.app/Contents/MacOS
chmod +x ./build/GoRPG.app/Contents/MacOS/game