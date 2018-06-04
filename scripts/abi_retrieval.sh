#!/bin/bash
#
# Simple script to generate and retrieve the ABIs/Bins from the Civil repo.
# Retrieves both the ABIs and the bytecode into separate files.
#
# Given a clone of the Civil main repo, run clean, install, and build there
# to generate the artifacts.  Then pull out all the abis from all the artifact
# json files and generate the abi and bin files in the destination directory.
#
# Stopgap solution until we come up with something automated and better.
#
# Usage: ./abi_retrieval.sh /full/path/to/main/civil/repo /full/path/to/dest/dir
#
# TODO(PN): Needs more validation and error checking.

set -e

CIVIL_REPO_DIR=$1
DEST_DIR=$2
ARTIFACTS_SUB_DIR=packages/contracts/build/artifacts
USAGE_STR="Usage: ./abi_retrieval.sh /full/path/to/main/civil/repo /full/path/to/dest/dir"

# Check for jq command to parse JSON
command -v jq >/dev/null 2>&1 || \
    { echo >&2 "jq needs to be installed. Go to: https://stedolan.github.io/jq/"; exit 1; }

# Check to see if directories are passed in
if [ -z $CIVIL_REPO_DIR ]
then
    echo $USAGE_STR
    echo "Civil repo directory must be 1st param"
    exit 1
fi

if [ -z $DEST_DIR ]
then
    echo $USAGE_STR
    echo "Destination directory must be 2nd param"
    exit 1
fi

echo "SOURCE:", $CIVIL_REPO_DIR
echo "DESTINATION:", $DEST_DIR

# Run a full clean and build of the Civil main repo
cd $CIVIL_REPO_DIR
yarn clean
yarn install
yarn build

# Ensure a fresh start with the ABIs
rm -rf $DEST_DIR
mkdir -p $DEST_DIR

cd $CIVIL_REPO_DIR/$ARTIFACTS_SUB_DIR
for filename in $CIVIL_REPO_DIR/$ARTIFACTS_SUB_DIR/*.json; do
    # Name with no file extension
    name=`echo $filename |cut -d "." -f 1 |rev | cut -d "/" -f 1 |rev`
    cat $filename |jq .abi > $DEST_DIR/$name.abi
    cat $filename |jq .bytecode | sed -e 's/^"//' -e 's/"$//' > $DEST_DIR/$name.bin
done

echo "Done with ABI/BIN retrieval"
