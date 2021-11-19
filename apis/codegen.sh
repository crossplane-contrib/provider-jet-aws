#!/bin/bash
set -eE

function command () {
  cd $1
  go run -tags generate sigs.k8s.io/controller-tools/cmd/controller-gen object:headerFile=../../hack/boilerplate.go.txt paths=./... crd:trivialVersions=true,allowDangerousTypes=true,crdVersions=v1 output:artifacts:config=../package/crds
  cd ..
}

for PKG in $(find * -maxdepth 0 -type d | xargs echo)
do
  echo "running for ${PKG}"
	if ! command $PKG; then
    echo "failed $PKG"
    exit 1
  fi
done