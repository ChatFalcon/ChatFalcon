#!/bin/sh
set -e
git submodule update --init
cd setup
npm i
npm run build
cd ..
cd frontend
npm i
npm run build
cd ..
