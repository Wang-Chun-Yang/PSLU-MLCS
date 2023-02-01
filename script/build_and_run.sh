# /bin/bash


cd ..
rm -fr bin/*

cd src

go build main.go
mv main ../bin

cd ../bin
./main -flagfile="../config/mlcs.flags"
