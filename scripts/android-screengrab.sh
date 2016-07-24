#!/bin/bash

#capture screenshot, copy to local then delete from device
adb shell screencap -p /sdcard/screen.png
adb pull /sdcard/screen.png
adb shell rm /sdcard/screen.png

#frame the screenshot
device-art --input-image=screen.png --device=nexus_6