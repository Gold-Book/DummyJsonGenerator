#!/bin/bash
dep ensure -v -vendor-only=true

# https://gist.github.com/jmoiron/3241916

LOCATION=${1:-./}

echo "watching location ${LOCATION}"

inotifywait -mqr --timefmt '%d/%m/%y %H:%M' --format '%T %w %f' \
   -e modify ${LOCATION} | while read date time dir file; do
    ext="${file##*.}"
    binaryFile="${file%.*}"
    if [[ "$ext" = "go" ]]; then
        echo "${dir}${file} changed @ $time $date $binaryFile, rebuilding..."
        if [ ! -e ./dist ]; then mkdir -p ./dist ; fi
        go build -o ./dist/${binaryFile} ${dir}${file}
        echo "done $binaryFile"
    fi
done
