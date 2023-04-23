package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/caicloud/nirvana/log"
	"github.com/gorilla/mux"
	"github.com/iFaceless/godub"
)

func main() {
	handleRequests()
}

func overlay(pubkeyHex string) error {
	segment, _ := godub.NewLoader().Load("test.mp3")
	fmt.Printf("Loaded: %s\n", segment)

	t0, err := godub.NewLoader().Load("tonesq2/t-00.mp3")
	t1, err := godub.NewLoader().Load("tonesq2/t-01.mp3")
	t2, err := godub.NewLoader().Load("tonesq2/t-02.mp3")
	t3, err := godub.NewLoader().Load("tonesq2/t-03.mp3")
	t4, err := godub.NewLoader().Load("tonesq2/t-04.mp3")
	t5, err := godub.NewLoader().Load("tonesq2/t-05.mp3")
	t6, err := godub.NewLoader().Load("tonesq2/t-06.mp3")
	t7, err := godub.NewLoader().Load("tonesq2/t-07.mp3")
	t8, err := godub.NewLoader().Load("tonesq2/t-08.mp3")
	t9, err := godub.NewLoader().Load("tonesq2/t-09.mp3")
	ta, err := godub.NewLoader().Load("tonesq2/t-10.mp3")
	tb, err := godub.NewLoader().Load("tonesq2/t-11.mp3")
	tc, err := godub.NewLoader().Load("tonesq2/t-12.mp3")
	td, err := godub.NewLoader().Load("tonesq2/t-13.mp3")
	te, err := godub.NewLoader().Load("tonesq2/t-14.mp3")
	tf, err := godub.NewLoader().Load("tonesq2/t-15.mp3")
	if err != nil {
		fmt.Println(err)
	}

	location := segment.Duration() / 5
	pubkeyHex = strings.ToUpper(pubkeyHex)
	for _, b := range pubkeyHex {
		switch string(b) {
		case "0":
			l, s, err := overlayTone(*segment, location, t0)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				location = l
				segment = &s
			}
		case "1":
			l, s, err := overlayTone(*segment, location, t1)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				location = l
				segment = &s
			}
		case "2":
			l, s, err := overlayTone(*segment, location, t2)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				location = l
				segment = &s
			}
		case "3":
			l, s, err := overlayTone(*segment, location, t3)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				location = l
				segment = &s
			}
		case "4":
			l, s, err := overlayTone(*segment, location, t4)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				location = l
				segment = &s
			}
		case "5":
			l, s, err := overlayTone(*segment, location, t5)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				location = l
				segment = &s
			}
		case "6":
			l, s, err := overlayTone(*segment, location, t6)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				location = l
				segment = &s
			}
		case "7":
			l, s, err := overlayTone(*segment, location, t7)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				location = l
				segment = &s
			}
		case "8":
			l, s, err := overlayTone(*segment, location, t8)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				location = l
				segment = &s
			}
		case "9":
			l, s, err := overlayTone(*segment, location, t9)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				location = l
				segment = &s
			}
		case "A":
			l, s, err := overlayTone(*segment, location, ta)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				location = l
				segment = &s
			}
		case "B":
			l, s, err := overlayTone(*segment, location, tb)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				location = l
				segment = &s
			}
		case "C":
			l, s, err := overlayTone(*segment, location, tc)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				location = l
				segment = &s
			}
		case "D":
			l, s, err := overlayTone(*segment, location, td)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				location = l
				segment = &s
			}
		case "E":
			l, s, err := overlayTone(*segment, location, te)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				location = l
				segment = &s
			}
		case "F":
			l, s, err := overlayTone(*segment, location, tf)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				location = l
				segment = &s
			}
		}
	}
	destPath := path.Join(tmpDataDirectory(), pubkeyHex+".mp3")
	err = godub.NewExporter(destPath).WithDstFormat("mp3").Export(segment)
	if err != nil {
		return err
	}
	return nil
}

func overlayTone(segment godub.AudioSegment, location time.Duration, tone *godub.AudioSegment) (time.Duration, godub.AudioSegment, error) {
	s, err := segment.Overlay(tone, &godub.OverlayConfig{Position: location, LoopToEnd: false, LoopCount: 1})
	if err != nil {
		return location, segment, err
	}

	location = location + tone.Duration()
	fmt.Println(location)
	return location, *s, nil
}

func tmpDataDirectory() string {
	return "temp"
}

func hex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println(key)
	if overlay(key) == nil {
		fileBytes, err := os.ReadFile("temp/" + key + ".mp3")
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(fileBytes)
	} else {
		fmt.Fprint(w, "something went wrong oops")
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/nostrovia/{id}", hex)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}