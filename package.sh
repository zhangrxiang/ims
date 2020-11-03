#!/usr/bin/env bash

NAME=simple-ims

echo "=============run=====start=========="

starttime=$(date '+%Y年%m月%d日 %H时%M分%S秒')
starttime2=$(date +'%Y-%m-%d %H:%M:%S')
echo "开始时间：$starttime"

echo "build....start..."
if [ -f ./${NAME}.exe ]; then
  ./${NAME}.exe build
  else
    go build
  ./${NAME}.exe init
fi

go build
echo "----------------------------------------------"
./${NAME}.exe info
echo "----------------------------------------------"

echo "build....over..."

version=$(./${NAME}.exe version | awk '{print $2}')

if [ ! -d "./production" ]; then
  mkdir ./production
fi

cd ./www-dev && yarn build
cd ..
if [ ! -d "./www" ]; then
  mkdir ./www
fi
rm -rf ./www/* && mv ./www-dev/dist/* www/

tar -czvf ./production/${NAME}."${version}".tar.gz --exclude=www/js/*.map ${NAME}.exe \
uninstall.bat \
install.bat \
www

endtime=$(date '+%Y年%m月%d日 %H时%M分%S秒')
endtime2=$(date +'%Y-%m-%d %H:%M:%S')
echo "结束时间：$endtime"
start_seconds=$(date --date="$starttime2" +%s)
end_seconds=$(date --date="$endtime2" +%s)
echo "本次运行时间： "$((end_seconds - start_seconds))"s"

echo "=============run=====end=========="
