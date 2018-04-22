#!/usr/bin/env bash

set -e

deployment_dir="${PWD}/deployment"
cpi_release_tarball="bosh-oracle-cpi.tgz"
stemcell_tarball="stemcell.tgz"

# State used by teardown tasks as well
manifest_filename="director-manifest.yml"
state_filename="director-state.json"

pwd=`pwd`

echo "Setting up artifacts..."
cp ./candidate/*.tgz ${deployment_dir}/${cpi_release_tarball}
cp ./stemcell/*.tgz ${deployment_dir}/${stemcell_tarball}
cp ./bosh-deployment/bosh.yml ${deployment_dir}/

vars_file=${deployment_dir}/director-env-vars.yml
jq '. + {}' < terraform-oci/metadata > ${vars_file}

ops_files=" -o ${pwd}/bosh-deployment/oracle/cpi.yml"
ops_files+=" -o ${pwd}/bosh-deployment/oracle/replace-default-network-with-public.yml"
ops_files+=" -o ${pwd}/bosh-deployment/oracle/use-public-ip-for-ssh.yml"
ops_files+=" -o ${pwd}/cpi-release-src/bosh-deployment/remove-hm.yml"

vars=" -v director_name=ci-bosh-director"
vars+=" -v local_oracle_cpi_release=${cpi_release_tarball}"
vars+=" -v local_oracle_light_stemcell=${stemcell_tarball}"


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
  bosh int ${ops_files}  --vars-store ./creds.yml --vars-file ${vars_file} ${vars} bosh.yml > ${manifest_filename}
  bosh create-env --state ${state_filename} ${manifest_filename}

  trap - ERR
  finish
popd
