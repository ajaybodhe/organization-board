#!/bin/sh
# check if GO is installed
which go
success=`echo $?`
if [[ $success != 0 ]]
then
echo "GO is not installed !"
exit 1
fi

# run tests
go test ./...