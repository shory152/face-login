#!/bin/sh


for i in /home/dbsql/projects/pycharm/face-login/photo/staffImage/*; do
  id=`basename $i | cut -c 4-6`;
  echo $id $i;
  if [ $id -lt 306 ]; then
    continue
  fi
  python face_reg.py $id $i;
done
