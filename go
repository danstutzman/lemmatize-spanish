#!/bin/bash
PSQL=/Applications/Postgres.app/Contents/MacOS/bin/psql
mkdir -p songs
for SONG_ID in `$PSQL -t -U postgres -c "select song_id from songs"`; do
  # -s means not zero size
  if [ -s songs/$SONG_ID.out ]; then
      $PSQL -t -U postgres -c "select line from lines where song_id = $SONG_ID order by line_id" > songs/$SONG_ID.txt
    cat songs/$SONG_ID.txt
    cat songs/$SONG_ID.txt | iconv -f utf-8 -t ISO8859-1 | \
      FREELINGSHARE=/usr/local/share/freeling \
      ./myfreeling/src/main/analyzer -f ./myfreeling/data/config/es.cfg \
      | iconv -f ISO8859-1 -t utf-8 | tee songs/$SONG_ID.out
    if [ "$?" == 130 ]; then exit 130; fi # if interrupted with Ctrl-C
  fi
done
