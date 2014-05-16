#!/bin/bash

### http://tecadmin.net/create-rpm-of-your-own-script-in-centosredhat/#

sudo yum install rpm-build rpmdevtools
rm -rf ~/rpmbuild
rpmdev-setuptree

mkdir ~/rpmbuild/SOURCES/hydra-worker-sort-by-number-1
cp hydra-worker-sort-by-number-init.d.sh ~/rpmbuild/SOURCES/hydra-worker-sort-by-number-1
cp ../../bin/hydra-worker-sort-by-number ~/rpmbuild/SOURCES/hydra-worker-sort-by-number-1

cp hydra-worker-sort-by-number.spec ~/rpmbuild/SPECS

pushd ~/rpmbuild/SOURCES/
tar czf hydra-worker-sort-by-number-1.0.tar.gz hydra-worker-sort-by-number-1/
cd ~/rpmbuild 
rpmbuild -ba SPECS/hydra-worker-sort-by-number.spec

popd
cp ~/rpmbuild/RPMS/x86_64/hydra-worker-sort-by-number-1-0.x86_64.rpm .
