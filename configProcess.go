package main

import "github.com/guotie/config"

func configRead() {
	masterDirPath = config.GetString("masterDir")
}
