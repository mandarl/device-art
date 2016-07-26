#!/bin/bash

if [ -z "$1" ]
  then
    echo "ERR! No version number supplied"
    exit
fi

version=$1

#clean up github deploy files
#rm -rf $GOPATH/bin/device-art-xc

#run the default task to publish to github
goxc -d=tmp -c=, -pv=$version

#generate binaries
goxc -c=binaries -pv=$version


#remove any .gz files from previous runs; go-selfupdate processes these files
#   as binary and creates .gz.gz file
rm $GOPATH/bin/device-art-xc/device-art/$version/*.gz*
rm $GOPATH/bin/device-art-xc/device-art/*.json
#generate selfupdate archives and json
go-selfupdate -o $GOPATH/bin/device-art-xc/device-art/ $GOPATH/bin/device-art-xc/$version/device-art/ $version


#publish to bintray
#goxc -c=bintray -pv=$version

#upload to server
lftp -vv -f upload.x

#clean up github deploy files
#rm -rf $GOPATH/bin/device-art-xc
