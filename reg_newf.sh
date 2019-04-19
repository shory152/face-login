#!/bin/sh


for i in newface/*; do
  id=`basename $i | cut -c 1-4`;
  echo $id $i;

  python face_reg.py $id $i;
done
