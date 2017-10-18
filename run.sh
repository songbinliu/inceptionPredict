#!/bin/bash

# make build

input="./imgs/tiger.jpg"
options="--v=2 --logtostderr"
options="$options --modeldir=./model-data/inception/"
options="$options --imgfile=$input" 

bin=./_output/inception
$bin $options
