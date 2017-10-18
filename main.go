package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
)

var (
	modeldir string
	imgfile  string
)

func setFlags() error {
	flag.StringVar(&modeldir, "modeldir", "./model-data/inception/", "model directory")
	flag.StringVar(&imgfile, "imgfile", "./imgs/cat.jpg", "path to the image file")
	flag.Parse()
	if modeldir == "" || imgfile == "" {
		flag.Usage()
		return fmt.Errorf("wrong parameter")
	}

	return nil
}

func main() {
	if nil != setFlags() {
		return
	}

	fmt.Printf("begin to load model from: %v\n", modeldir)
	model := NewModel(modeldir)
	model.Init()
	glog.V(2).Infof("model init success.")

	model.PredictTopkFile(imgfile, 5)
}
