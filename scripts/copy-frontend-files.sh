#!/bin/bash

mkdir -p ../services/frontend/internal/frontend/assets
cp -r ../frontend/dist/frontend/browser/* ../services/frontend/internal/frontend/assets

mkdir -p ../services/frontend/internal/docs/assets
cp -r ../docs/.retype/* ../services/frontend/internal/docs/assets