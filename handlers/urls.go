package handlers

import (
	"net/http"
	index "weixin_api/handlers/index"
	adv "weixin_api/handlers/adv"
)

func MyUrls() {
	http.HandleFunc("/", index.Index)
	http.HandleFunc("/manager/advert/add_advert", adv.Add_adver)
	http.HandleFunc("/add_file", index.Add_file)
	http.HandleFunc("/upload_file",index.Upload_file)
}