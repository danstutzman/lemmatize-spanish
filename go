#!/bin/bash
PSQL=/Applications/Postgres.app/Contents/MacOS/bin/psql
mkdir -p songs
for SOURCE_NUM in `$PSQL -t -U postgres -c "select source_num from songs"`; do
  # -s means not zero size
  if [ ! -s songs/$SOURCE_NUM.out ]; then
      $PSQL -t -U postgres -c "select line from lines where song_source_num = $SOURCE_NUM order by line_id" > songs/$SOURCE_NUM.txt
    cat songs/$SOURCE_NUM.txt
    cat songs/$SOURCE_NUM.txt | iconv -f utf-8 -t ISO8859-1 | \
      FREELINGSHARE=/usr/local/share/freeling \
      ./myfreeling/src/main/analyzer -f ./myfreeling/data/config/es.cfg \
      | iconv -f ISO8859-1 -t utf-8 | tee songs/$SOURCE_NUM.out
    if [ "$?" == 130 ]; then exit 130; fi # if interrupted with Ctrl-C
  fi
done
