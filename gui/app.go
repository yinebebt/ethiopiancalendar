package gui

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Run starts the GUI application.
func Run() {
	a := app.New()
	a.Settings().SetTheme(&ethioTheme{})

	title := "Ethiocal — Ethiopian Calendar"
	w := a.NewWindow(title)
	w.Resize(fyne.NewSize(600, 700))

	homeTab := newHomeTab()
	converterTab := newConverterTab()
	bahirTab := newBahirTab()

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), homeTab),
		container.NewTabItemWithIcon("Date Converter", theme.HistoryIcon(), converterTab),
		container.NewTabItemWithIcon("Bahire-Hasab", theme.ListIcon(), bahirTab),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	w.SetContent(container.NewBorder(nil, newFooter(), nil, nil, tabs))
	w.ShowAndRun()
}

// centeredColumn lays a single child out as a centered column capped at maxWidth,
// with a uniform margin: breathing room on wide desktop windows while still
// filling narrow mobile screens.
type centeredColumn struct {
	maxWidth, margin float32
}

func (c *centeredColumn) MinSize(objs []fyne.CanvasObject) fyne.Size {
	if len(objs) == 0 {
		return fyne.Size{}
	}
	m := objs[0].MinSize()
	return fyne.NewSize(m.Width+2*c.margin, m.Height+2*c.margin)
}

func (c *centeredColumn) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	if len(objs) == 0 {
		return
	}
	w := min(size.Width-2*c.margin, c.maxWidth)
	objs[0].Resize(fyne.NewSize(w, size.Height-2*c.margin))
	objs[0].Move(fyne.NewPos((size.Width-w)/2, c.margin))
}

// centered wraps content in a margined, max-width centered column.
func centered(o fyne.CanvasObject) fyne.CanvasObject {
	return container.New(&centeredColumn{maxWidth: 560, margin: theme.Padding() * 3}, o)
}

// accentHeading returns a bold, accent-colored heading. It uses RichText with a
// themed color name, so the bundled Ethiopic font still renders Amharic.
func accentHeading(text string) *widget.RichText {
	return widget.NewRichText(&widget.TextSegment{
		Text: text,
		Style: widget.RichTextStyle{
			Alignment: fyne.TextAlignCenter,
			ColorName: theme.ColorNamePrimary,
			SizeName:  theme.SizeNameSubHeadingText,
			TextStyle: fyne.TextStyle{Bold: true},
		},
	})
}

// newFooter creates a compact footer with source and author links.
func newFooter() fyne.CanvasObject {
	ghURL, _ := url.Parse("https://github.com/yinebebt/ethiocal")
	ghLink := widget.NewHyperlink("Source", ghURL)

	authorURL, _ := url.Parse("https://yinebebt.com")
	authorLink := widget.NewHyperlink("yinebebt.com", authorURL)

	sep := widget.NewLabel("·")

	return container.NewVBox(
		widget.NewSeparator(),
		container.NewCenter(
			container.NewHBox(ghLink, sep, authorLink),
		),
	)
}
