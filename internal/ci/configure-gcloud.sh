#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <deployment>"
    exit 1
fi
deployment=$1

set -ex

export GOOGLE_APPLICATION_CREDENTIALS=/keyconfig.json

gcloud auth activate-service-account
gcloud config set compute/zone us-central1-a
gcloud container clusters get-credentials "$deployment"
