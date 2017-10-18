package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"github.com/golang/glog"
	"strings"
	"path/filepath"
)

var (
	modeldir string
	imgfile  string
	imgdir string
)

func setFlags() error {
	flag.StringVar(&modeldir, "modeldir", "./model-data/inception/", "model directory")
	flag.StringVar(&imgfile, "imgfile", "./imgs/cat.jpg", "path to the image file")
	flag.StringVar(&imgdir, "imgdir", "./imgs/", "path to the image files")
	flag.Parse()
	if modeldir == "" {
		fmt.Println("modeldir must be provided.")
		flag.Usage()
		return fmt.Errorf("wrong parameter")
	}

	return nil
}

func predictDir(model *TfModel, dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		glog.Errorf("Failed to readDir %v: %v", dir, err)
		return
	}

	for _, file := range files {
		fname := file.Name()
		if !strings.HasSuffix(fname, "jpg") {
			continue
		}

		fname = filepath.Join(dir, file.Name())
		model.PredictTopkFile(fname, 5)
	}
}

func main() {
	if nil != setFlags() {
		return
	}

	fmt.Printf("begin to load model from: %v\n", modeldir)
	model := NewModel(modeldir)
	model.Init()
	glog.V(2).Infof("model init success.")

	if len(imgdir) > 0 {
		predictDir(model, imgdir)
	}

	if len(imgdir) == 0 && len(imgfile) > 0 {
		model.PredictTopkFile(imgfile, 5)
	}

	fmt.Println("bye")
	return
}
