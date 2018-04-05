#!/usr/bin/env bash

set -e

deployment_dir="${PWD}/deployment"
cpi_release_name="bosh-oracle-cpi"
state_filename="director-state.json"

pwd=`pwd`

echo "Setting up artifacts..."
cp ./candidate/*.tgz ${deployment_dir}/${cpi_release_name}.tgz
cp ./stemcell/*.tgz ${deployment_dir}/stemcell.tgz
cp ./bosh-deployment/bosh.yml ${deployment_dir}/

vars_file=${pwd}/oci-config/director-env-vars.yml

ops_files=" -o ${pwd}/bosh-deployment/oracle/cpi.yml"
ops_files+=" -o ${pwd}/bosh-deployment/oracle/replace-default-network-with-public.yml"
ops_files+=" -o ${pwd}/bosh-deployment/oracle/use-public-ip-for-ssh.yml"
ops_files+=" -o ${pwd}/cpi-release-src/bosh-deployment/remove-hm.yml"

# Use the candidate artifacts
local_yml="local.yml"
cat >"${deployment_dir}/${local_yml}"<<EOF
---
- type: replace
  path: /releases/name=bosh-oracle-cpi
  value:
    name: bosh-oracle-cpi
    url: file://${cpi_release_name}.tgz

- type: replace
  path: /resource_pools/name=vms/stemcell?
  value:
    url: file://stemcell.tgz
EOF

ops_files+=" -o ${deployment_dir}/${local_yml}"

pushd ${deployment_dir}
  function finish {
    echo "Final state of director deployment:"
    echo "=========================================="
    cat ${state_filename}
    echo "=========================================="
  }
  trap finish ERR

  echo "Using BOSH CLI version..."
  bosh -v

  ls -al 

  echo "Deploying BOSH Director..."
  bosh create-env ${ops_files}  --vars-store ./creds.yml --state ${state_filename} --vars-file ${vars_file} bosh.yml

  trap - ERR
  finish
popd
