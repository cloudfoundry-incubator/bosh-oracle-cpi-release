#!/bin/sh

set -e

pwd=`pwd`
env_dir=${pwd}/terraform-env
api_key_path=terraform-env/oci_api_key.pem
vars_file=${env_dir}/oci.vars

mkdir -p $env_dir

echo "Creating terraform variables file..."

cat > ${api_key_path} <<EOF
${oracle_apikey}
EOF
chmod 600 ${api_key_path}

cat > ${vars_file} <<EOF
oracle_private_key_path: ${api_key_path}
EOF

echo "Done. Created: " ${vars_file}
