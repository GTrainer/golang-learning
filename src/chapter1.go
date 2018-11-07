package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{ color.White, color.Black }

const (
	whiteIndex = 0
	blackIndex = 1
)

func createGIF(out io.Writer) {
	const (
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{ LoopCount: nframes }
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles * 2 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}

func curl() {
	for _, url := range os.Args[1:] {
		rsp, err := http.Get(url)
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(rsp.Body)
		rsp.Body.Close()

		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		fmt.Printf("%s", body)
	}
}

func channelCurl(url string, barrier chan string) {
	start := time.Now()

	rsp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	_, err = ioutil.ReadAll(rsp.Body)
	rsp.Body.Close()

	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	barrier<-fmt.Sprintf("%.2fs %s", time.Since(start).Seconds(), url)
} 

func mcurl() {
	start := time.Now()
	
	barrier := make(chan string)
	for _, url := range os.Args[1:] {
		go channelCurl(url, barrier)
	}

	for range os.Args[1:] {
		fmt.Println(<-barrier)
	}

	fmt.Printf("%.2fs elapsed, finish\n", time.Since(start).Seconds())
}

func server() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	createGIF(w)
}

