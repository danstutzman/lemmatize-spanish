#!/bin/bash -ex

#brew install icu4c
#brew install boost --with-icu4c

#svn checkout http://devel.cpl.upc.edu/freeling/svn/trunk myfreeling
#svn checkout http://devel.cpl.upc.edu/freeling/svn/versions/freeling-3.1

cd myfreeling

autoreconf --install

#cd /opt/local/include/boost
#sudo ln -s property_map/property_map.hpp 

env LDFLAGS="-L/opt/local/lib -L/opt/local/lib/db46 -L/usr/local/opt/icu4c/lib" CPPFLAGS="-I/opt/local/include -I/opt/local/include/boost -I/opt/local/include/db46 -I/usr/local/Cellar/boost/1.59.0/include -I/usr/local/opt/icu4c/include" ./configure --enable-boost-locale
make 
make install

if [ ! -e myfreeling/data/config/es.cfg.bak ]; then
  cp myfreeling/data/config/es.cfg myfreeling/data/config/es.cfg.bak
fi
cat myfreeling/data/config/es.cfg.bak | sed "s/NERecognition=yes/NERecognition=no/" > myfreeling/data/config/es.cfg
