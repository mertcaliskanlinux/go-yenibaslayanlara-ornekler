package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"

	"github.com/vdobler/chart"
	"github.com/vdobler/chart/imgg"
)

//main blogu içide grafik oluşturuyoruz
func main() {
	const (
		GrafikEni  = 800 //bir grafik için
		GrafikBoyu = 500 //bir grafik için
		Sutun      = 2
		Satir      = 1
	)
	//PASTA GRAFİĞİ
	PastaGrafigi := chart.PieChart{Title: "Birşeylerin Dağılımı"}
	PastaGrafigi.AddDataPair("Veriler",
		[]string{"Aaa", "Bbb", "Ccc", "Ddd", "Eee"},
		[]float64{10, 20, 30, 35, 15, 25})
	PastaGrafigi.Data[0].Samples[3].Flag = true
	PastaGrafigi.Inner = 0.5
	PastaGrafigi.Key.Border = -1
	PastaGrafigi.FmtVal = chart.AbsoluteValue

	//Sinüs Eğrisi

	SinusGrafigi := chart.ScatterChart{Title: "Sinüs Grafiği"}
	SinusGrafigi.XRange.Fixed(0, 4*math.Pi, math.Pi)
	SinusGrafigi.YRange.Fixed(-1.25, 1.25, 0.5)
	SinusGrafigi.XRange.TicSetting.Format = func(f float64) string {
		w := int(180*f/math.Pi + 0.5)
		return fmt.Sprintf("%d", w)
	}

	SinusGrafigi.AddFunc("Sin(x)",
		func(f float64) float64 { return math.Sin(f) }, chart.PlotStyleLines, chart.Style{Symbol: '@', LineWidth: 2, LineColor: color.NRGBA{0x00, 0x00, 0xcc, 0xff}, LineStyle: 0})

	SinusGrafigi.XRange.TicSetting.Tics, SinusGrafigi.YRange.TicSetting.Tics = 1, 1
	SinusGrafigi.XRange.TicSetting.Mirror, SinusGrafigi.YRange.TicSetting.Mirror = 2, 2
	SinusGrafigi.XRange.TicSetting.Grid, SinusGrafigi.YRange.TicSetting.Grid = 2, 1
	SinusGrafigi.XRange.ShowZero = true

	//Grafik Dosyalarını Kaydet

	pasta := GrafikKayitcisi("PastaGrafikDosyası", 1, 1, GrafikEni, GrafikBoyu)
	pasta.GrafikCiz(&PastaGrafigi)
	sinus := GrafikKayitcisi("Sinus Grafiği", 1, 1, GrafikEni, GrafikBoyu)
	sinus.GrafikCiz(&SinusGrafigi)

}

type GrafikYapısı struct {
	N, M, W, H, Cnt int
	I               *image.RGBA
	imgfile         *os.File
}

func GrafikKayitcisi(name string, n, m, w, h int) *GrafikYapısı {
	dumper := GrafikYapısı{N: n, M: m, W: w, H: h}
	dumper.imgfile, _ = os.Create(name + ".png")
	dumper.I = image.NewRGBA(image.Rect(0, 0, n*w, m*h))
	bg := image.NewUniform(color.RGBA{0xff, 0xff, 0xff, 0xff})
	draw.Draw(dumper.I, dumper.I.Bounds(), bg, image.ZP, draw.Src)
	return &dumper
}

func (d *GrafikYapısı) GrafikCiz(c chart.Chart) {
	row, col := d.Cnt/d.N, d.Cnt/d.N
	igr := imgg.AddTo(d.I, col*d.W, row*d.H, d.W, d.H, color.RGBA{0xff, 0xff, 0xff, 0xff}, nil, nil)
	c.Plot(igr)
	d.Cnt++
	png.Encode(d.imgfile, d.I)
	d.imgfile.Close()
}
