package tutorials

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/cmd/fyne_demo/data"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

func welcomeScreen(_ fyne.Window) fyne.CanvasObject {
	logo := canvas.NewImageFromResource(data.FyneScene)
	//logo_flowshield := canvas.NewImageFromFile("/Users/libiao/DATA/ICODE/FlowShield/fyne-gui/cmd/fyne_demo/logo.png")
	logo.FillMode = canvas.ImageFillContain
	if fyne.CurrentDevice().IsMobile() {
		logo.SetMinSize(fyne.NewSize(192, 192))
	} else {
		logo.SetMinSize(fyne.NewSize(256, 256))
	}

	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("Welcome to FlowShield Security Network", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		logo,
		container.NewHBox(
			widget.NewHyperlink("flowshield.xyz", parseURL("https://flowshield.xyz")),
			widget.NewLabel("-"),
			widget.NewHyperlink("documentation", parseURL("https://www.flowshield.xyz/flowshield_docs")),
			widget.NewLabel("-"),
			widget.NewHyperlink("OpenSource", parseURL("https://github.com/flowshield")),
		),
		widget.NewLabel(""), // balance the header on the tutorial screen we leave blank on this content
	))
}
