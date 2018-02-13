#!/bin/bash

for datafile in $( ls ./data/. | grep .csv$ ); do
   ./main -file data/${datafile}
done