package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func response(rw http.ResponseWriter, request *http.Request) {
	links := []string{
		`https://upload.wikimedia.org/wikipedia/commons/a/a4/Misc_pollen.jpg`,
		`http://www.wired.com/wp-content/uploads/images_blogs/wiredscience/2009/10/nikon2003_1st_wittmann.jpg`,
		`http://www.scientificcomputing.com/sites/scientificcomputing.com/files/21166%20-%20Michael%20Shribak_Entry_21166_Rotifer-VEP%20polscope.jpg`,
		`http://wallpoper.com/images/00/42/67/18/canada-microscopic_00426718.jpg`,
		`http://easyscienceforkids.com/wp-content/uploads/2013/06/virus-under-microscope.jpg`,
		`https://upload.wikimedia.org/wikipedia/commons/d/dd/Stem_of_first_year_Pinus_taiwanensis_(Taiwan_Red_Pine)_-_Cross_section_microscopic_image_(detail).jpg`,
		`http://40.media.tumblr.com/tumblr_lhnrlpBp7l1qbpwkro1_1280.jpg`,
		`https://d1w5usc88actyi.cloudfront.net/wp-content/uploads/2012/08/Oliver-Meckes-4.jpeg`,
	}

	returnUrl := links[rand.Intn(len(links))]
	json, err := json.Marshal(returnUrl)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	rw.Write([]byte(json))
}

func main() {
	log.Println("Look reeeeallly close on :8000/")
	http.HandleFunc("/", response)
	http.ListenAndServe(":8000", nil)
}
