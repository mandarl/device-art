rem change to the system temp directory
cd /D %temp%

rem capture screenshot, copy to local then delete from device
adb shell screencap -p /sdcard/screen.png
adb pull /sdcard/screen.png
adb shell rm /sdcard/screen.png

rem frame the screenshot
device-art --input-image=screen.png --device=nexus_6

rem open explorer and select the output file
explorer /select,%temp%\output.png