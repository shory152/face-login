#!/bin/sh


for i in staffImage/*; do
  id=`basename $i | cut -c 4-6`;
  echo $id $i;

  python face_reg.py $id $i;
done
