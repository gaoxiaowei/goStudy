package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
)

var palette =[]color.Color{color.White,color.Black}
const (
	whiteIndex =0 //画板中的第一种颜色
	blackIndex =1 //画板中的下一种颜色
)


func main()  {
	fmt.Println("start")
	handler := func(w http.ResponseWriter,r *http.Request){
		lissajous_(w)
	}
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

//func handler(w http.ResponseWriter,r*http.Request){
//	fmt.Fprintf(w,"%s %s %s\n",r.Method,r.URL,r.Proto)
//	for k,v :=range  r.Header{
//		fmt.Fprint(w,"Header[%q] =%q\n",k,v)
//	}
//	fmt.Fprintf(w,"Host =%q\n",r.Host)
//	fmt.Fprintf(w,"RemoteAddr =%q\n",r.RemoteAddr)
//	if err :=r.ParseForm();err!=nil{
//		log.Print(err)
//	}
//	for k,v :=range r.Form{
//		fmt.Fprintf(w,"Form[%q] =%q \n",k,v)
//	}
//
//}
func lissajous_(out io.Writer){
	const (
		cycles =5
		res =0.001
		size =100
		nframes =64
		delay =8
	)
	freq :=rand.Float64()* 3.0
	anim :=gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i :=0;i<nframes;i++{
		rect := image.Rect(0,0,2*size+1,2*size+1)
		img := image.NewPaletted(rect,palette)
		for t:= 0.0;t<cycles *2*math.Pi;t +=res{
			x :=math.Sin(t)
			y :=math.Sin(t*freq+phase)
			img.SetColorIndex(size+int(x*size+0.5),size+int(y*size+0.5),blackIndex)
		}
		phase +=0.1
		anim.Delay =append(anim.Delay,delay)
		anim.Image =append(anim.Image,img)

	}
	gif.EncodeAll(out,&anim)
}