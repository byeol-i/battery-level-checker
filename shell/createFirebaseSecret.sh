#!/bin/bash
cd ..

docker secret create firebase-key ./conf/firebase/key.json

cd shell