package services

import (
	"log"

	"github.com/marvinarlt/goshort.link/internal/consts"
	"github.com/thanhpk/randstr"
)

func UniqueId(length int) string {
	var exists bool

	id := randstr.String(length)
	err := DB.Raw("SELECT EXISTS(SELECT * FROM links WHERE id = ?)", id).Scan(&exists).Error

	if nil != err {
		log.Fatalln(consts.PREFIX, "Something went wrong while creating id")
	}

	if exists {
		UniqueId(length)
		return ""
	}

	return id
}
