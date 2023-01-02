#!/bin/bash

echo "########### Loading data to Mongo DB ###########"
mongoimport --jsonArray --db dealls --collection user --file /tmp/data/data.json