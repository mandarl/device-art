#!/bin/bash

#capture screenshot using libimobiledevice, brew install libimobiledevice to install
idevicescreenshot screen.tiff

#convert tiff to png
sips -s format png screen.tiff --out screen.png

#frame the screenshot
device-art --input-image=screen.png --device=iphone_6_plus
