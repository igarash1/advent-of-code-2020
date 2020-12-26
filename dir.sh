for y in {2015..2020}
do
  mkdir $y
  for d in {1..27}
  do
    if [ $d -le 9 ]
    then
      mkdir "$y/0$d"
    else
      mkdir "$y/$d"
    fi
  done
done
