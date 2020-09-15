#!/bin/sh
# check if GO ise installed
which go
success=`echo $?`
if [[ $success != 0 ]]
then
echo "GO is not installed !"
exit 1
fi

# Build the code
go build
# Remove the DB file
# rm resource/organization-board.db
# Run the binary
./organization-board
