# device-art
device-art: a simple utility for framing your phone screenshots in a device skin.

Currently supported device skins are Nexus 5 (1080x1920), 6 (1440x2560), 5x (1080x1920), 6p (1440x2560).
The screenshot must match the dimensions of the device.

![enter image description here](https://raw.githubusercontent.com/mandarl/device-art/master/art/screenshot.jpg)

## Download
Download the latest versions for Windows, Mac and Linux on the GitHub releases page:
https://github.com/mandarl/device-art/releases

## Usage
device-art can be used on the desktop to quickly capture a screenshot and frame it, or it can be used on a build server to frame screenshots taken as part of automated tests.

Here are example scripts to capture and frame screenshots: https://github.com/mandarl/device-art/tree/master/scripts


Option | Description |
:--- | :--- |
-h, --help | display help information |
-i, --input-image | *input screenshot image |
-d, --device[=nexus_6] | the device skin to use |
-o, --orientation[=port] | device orientation; can be port or land |
-p, --output-file[=output.png] | output file path |


## ToDo
- Add more devices
- Cache device images locally, currently it is an http get every time
- Add batch processing for multiple screenshots
- Option to read device images from local directory 
