#!/usr/bin/env bash

set -e

deployment_dir="${PWD}/deployment"
manifest_filename="director-manifest.yml"
state_filename="director-state.json"

pushd ${deployment_dir}

  echo "Using BOSH CLI version..."
  bosh -v

  echo "Deleting BOSH Director..."
  bosh delete-env --state ${state_filename} ${manifest_filename}
popd
