#!/bin/bash
#check if go is installed
if ! [ -x "$(command -v go)" ]; then
  echo 'Error: go is not installed.' >&2
  exit 1
fi

#check if git is installed
if ! [ -x "$(command -v git)" ]; then
  echo 'Error: git is not installed.' >&2
  exit 1
fi

git clone https://github.com/ProductionPanic/imshow.git ./imshow
sleep 1
cd imshow
sleep 1
go build -o imshow
sleep 1
sudo mv ./imshow /usr/local/bin
sleep 1
cd ..
sleep 1
rm -rf imshow
sleep 1
