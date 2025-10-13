package tutorials

import (
	"image"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var tutorials = map[string]GeneratedTutorial{
	"canvas/circle.md": {
		title:   "Canvas : Circle",
		content: []string{"The Circle canvas object draws a styled circle in the center of the available space.\nIts diameter will be the smaller of the available width and height, so it remains circular whilst filling the space.\n\n## Basic usage\n\nA circle is created using the `NewCircle` constructor function, passing in the color it should use to draw.\nA basic circle is a simple fill, though it can also have a stroke (border) as you will see below.", "circle := canvas.NewCircle(color.NRGBA{0, 0, 0x80, 0xff})", "It is recommended to use the NRGBA color type to avoid complications with transparency.\n\n## Setting styles\n\nBy default, a circle has no stroke (border), but you can set it up using the\n`StrokeWidth` and `StrokeColor` fields.", "circle := canvas.NewCircle(color.NRGBA{0x30, 0x30, 0x30, 0xb0})\ncircle.StrokeColor = color.NRGBA{0x80, 0, 0, 0xff}\ncircle.StrokeWidth = 2", ""},

		code: []func() fyne.CanvasObject{
			func() fyne.CanvasObject {
				circle := canvas.NewCircle(color.NRGBA{0, 0, 0x80, 0xff})
				return circle
			},
			func() fyne.CanvasObject {
				circle := canvas.NewCircle(color.NRGBA{0x30, 0x30, 0x30, 0xb0})
				circle.StrokeColor = color.NRGBA{0x80, 0, 0, 0xff}
				circle.StrokeWidth = 2
				return circle
			},
		},
	},

	"canvas/gradient.md": {
		title:   "Canvas : Gradient",
		content: []string{"A gradient canvas object fills an area with a gradient from one color to another.\nThere are two types of gradient, LinearGradient and RadialGradient.\n\n## LinearGradient\n\nThe simplest linear gradient is created using the `NewHorizontalGradient` or `NewVerticalGradient` constructor function,\npassing in the start and end colors it should use to draw.", "grad := canvas.NewHorizontalGradient(color.NRGBA{0x80, 0, 0, 0xff}, color.NRGBA{0, 0x80, 0, 0xff})", "### Specifying an angle\n\nYou can specify the angle of a gradient in 45 degree increments to control the angle manually.", "grad := canvas.NewLinearGradient(color.NRGBA{0x80, 0, 0, 0xff}, color.NRGBA{0, 0x80, 0, 0xff}, 45)", "## RadialGradient\n\nA radial gradient draws the colour gradient from the center out to the perimiters with a circular draw.\nFor example:", "grad := canvas.NewRadialGradient(color.NRGBA{0x80, 0, 0, 0xff}, color.NRGBA{0, 0x80, 0x80, 0xff})", ""},

		code: []func() fyne.CanvasObject{
			func() fyne.CanvasObject {
				grad := canvas.NewHorizontalGradient(color.NRGBA{0x80, 0, 0, 0xff}, color.NRGBA{0, 0x80, 0, 0xff})
				return grad
			},
			func() fyne.CanvasObject {
				grad := canvas.NewLinearGradient(color.NRGBA{0x80, 0, 0, 0xff}, color.NRGBA{0, 0x80, 0, 0xff}, 45)
				return grad
			},
			func() fyne.CanvasObject {
				grad := canvas.NewRadialGradient(color.NRGBA{0x80, 0, 0, 0xff}, color.NRGBA{0, 0x80, 0x80, 0xff})
				return grad
			},
		},
	},

	"canvas/image.md": {
		title:   "Canvas : Image",
		content: []string{"The Image canvas object is used to draw a loaded image to the screen.\nBy default, the image will fill available space, be sure to place it in container where space is allocated appropriately.\n\n## Basic usage\n\nAn image is typically created using the `NewImageFromFile`, `NewImgeFromResource`, or `NewImageFromImage` constructor functions,\npassing the appropriate parameter to specify the location of the image to display.", "img := canvas.NewImageFromResource(theme.HomeIcon())", "Note that the default minimum size for an image is just 1x1 because a Fyne app does not use pixel sizes for layout.\n\n## Setting fill and scale modes\n\nYou can control scale and fill modes to adjust how your iamge fills the space.\nIt is also possible to make the image blend into the background using `Translucency`.", "img := canvas.NewImageFromResource(theme.ComputerIcon())\nimg.FillMode = canvas.ImageFillContain\nimg.ScaleMode = canvas.ImageScaleFastest\nimg.Translucency = 0.5", ""},

		code: []func() fyne.CanvasObject{
			func() fyne.CanvasObject {
				img := canvas.NewImageFromResource(theme.HomeIcon())
				return img
			},
			func() fyne.CanvasObject {
				img := canvas.NewImageFromResource(theme.ComputerIcon())
				img.FillMode = canvas.ImageFillContain
				img.ScaleMode = canvas.ImageScaleFastest
				img.Translucency = 0.5
				return img
			},
		},
	},

	"canvas/line.md": {
		title:   "Canvas : Line",
		content: []string{"The Line canvas object draws a styled line from one position to another.\nThe default positions correspond with the top, left and bottom, right of its bounding box, \nthough they can be controlled using Position1 and Position2 fields.\nTherefor a line should normally be drawn in a container without layout, or a custom widget.\n\n## Basic usage\n\nA line is created using the `NewLine` constructor function, passing in the color it should use to draw.", "line := canvas.NewLine(color.NRGBA{0, 0, 0x80, 0xff})", "It is recommended to use the NRGBA color type to avoid complications with transparency.\n\n## Setting styles\n\nBy default, a line is 1 unit across (device independent pixel).\nYou can specify different values using the `StrokeWidth` field.", "line := canvas.NewLine(color.NRGBA{0, 0x80, 0, 0xdf})\nline.StrokeWidth = 5", ""},

		code: []func() fyne.CanvasObject{
			func() fyne.CanvasObject {
				line := canvas.NewLine(color.NRGBA{0, 0, 0x80, 0xff})
				return line
			},
			func() fyne.CanvasObject {
				line := canvas.NewLine(color.NRGBA{0, 0x80, 0, 0xdf})
				line.StrokeWidth = 5
				return line
			},
		},
	},

	"canvas/raster.md": {
		title:   "Canvas : Raster",
		content: []string{"The Raster canvas object is used to draw image content to screen programmatically.\nA raster object is used to fill pixels on screen, where the number of pixels will vary based on the\ndensity of the output device and the space available.\n\n## Basic usage\n\nA raster is typically created using the `NewRaster` constructor which passes in a generator function.\nThe generator function is called when the raster refreshes or the number of pixels required changes.", "rast := canvas.NewRaster(func(w, h int) image.Image {\n\tret := image.NewNRGBA(image.Rect(0, 0, w, h))\n\tfor y := 0; y < h; y++ {\n\t\tfor x := 0; x < w; x++ {\n\t\t\tif (x+y)/4%2 == 0 {\n\t\t\t\tret.Set(x, y, color.Black)\n\t\t\t} else {\n\t\t\t\tret.Set(x, y, color.White)\n\t\t\t}\n\t\t}\n\t}\n\treturn ret\n})", ""},

		code: []func() fyne.CanvasObject{
			func() fyne.CanvasObject {
				rast := canvas.NewRaster(func(w, h int) image.Image {
					ret := image.NewNRGBA(image.Rect(0, 0, w, h))
					for y := 0; y < h; y++ {
						for x := 0; x < w; x++ {
							if (x+y)/4%2 == 0 {
								ret.Set(x, y, color.Black)
							} else {
								ret.Set(x, y, color.White)
							}
						}
					}
					return ret
				})
				return rast
			},
		},
	},

	"canvas/rectangle.md": {
		title:   "Canvas : Line",
		content: []string{"The Rectang;e canvas object draws a styled rectangle that fills available space.\nA basic rectangle is a simple fill, though it can have a stroke and also rounded corners, as you will see below.\n\n## Basic usage\n\nA rectangle is created using the `NewRectangle` constructor function, passing in the color it should use to draw for fill.\nBy default, a rectangle will have no stroke (border) and just a solit fill:", "rect := canvas.NewRectangle(color.NRGBA{0x80, 0, 0, 0xff})", "It is recommended to use the NRGBA color type to avoid complications with transparency.\n\n## Setting styles\n\nBy default, a rectangle has no stroke (border), but you can set it up using the\n`StrokeWidth` and `StrokeColor` fields.", "rect := canvas.NewRectangle(color.NRGBA{0x80, 0, 0, 0xff})\nrect.StrokeColor = color.NRGBA{0xff, 0x80, 0, 0xff}\nrect.StrokeWidth = 3", "## Rounded Corners\n\nA rectangle (and its stroke) can also have rounded corners, controlled by the `CornerRadius` field:", "rect := canvas.NewRectangle(color.NRGBA{0x80, 0, 0, 0xff})\nrect.StrokeColor = color.NRGBA{0xff, 0, 0x80, 0xff}\nrect.StrokeWidth = 2.0\nrect.CornerRadius = 5", ""},

		code: []func() fyne.CanvasObject{
			func() fyne.CanvasObject {
				rect := canvas.NewRectangle(color.NRGBA{0x80, 0, 0, 0xff})
				return rect
			},
			func() fyne.CanvasObject {
				rect := canvas.NewRectangle(color.NRGBA{0x80, 0, 0, 0xff})
				rect.StrokeColor = color.NRGBA{0xff, 0x80, 0, 0xff}
				rect.StrokeWidth = 3
				return rect
			},
			func() fyne.CanvasObject {
				rect := canvas.NewRectangle(color.NRGBA{0x80, 0, 0, 0xff})
				rect.StrokeColor = color.NRGBA{0xff, 0, 0x80, 0xff}
				rect.StrokeWidth = 2.0
				rect.CornerRadius = 5
				return rect
			},
		},
	},

	"canvas/text.md": {
		title:   "Canvas : Text",
		content: []string{"The Text canvas object is used to draw text to screen with control over its output.\n\n## Basic usage\n\nText is created using the `NewText` constructor function, passing in the string to\ndisplay along with the color it should use to draw.", "txt := canvas.NewText(\"Hello\", color.NRGBA{0, 0x80, 0, 0xff})", "It is recommended to use the NRGBA color type to avoid complications with transparency.\n\n## Setting styles\n\nBy default, text is the normal style (not italic or bold or monospace) and\nthe font size matches the operating system normal text (approx 14pt).\n\nYou can specify different values for these and other fields as shown:", "txt := canvas.NewText(\"Styled\", color.NRGBA{0x80, 0, 0, 0xb0})\ntxt.TextSize = 24\ntxt.TextStyle.Bold = true\ntxt.Alignment = fyne.TextAlignTrailing", ""},

		code: []func() fyne.CanvasObject{
			func() fyne.CanvasObject {
				txt := canvas.NewText("Hello", color.NRGBA{0, 0x80, 0, 0xff})
				return txt
			},
			func() fyne.CanvasObject {
				txt := canvas.NewText("Styled", color.NRGBA{0x80, 0, 0, 0xb0})
				txt.TextSize = 24
				txt.TextStyle.Bold = true
				txt.Alignment = fyne.TextAlignTrailing
				return txt
			},
		},
	},

	"widgets/accordion.md": {
		title:   "Widgets : Accordion",
		content: []string{"Expand or collapse content panels.\n\n## Basic usage\n\nA basic accordion can be created with `NewAccordion` passing in items created using\nthe `NewAccordionItem` helper, as shown:", "ac := widget.NewAccordion(\n\twidget.NewAccordionItem(\"A\", widget.NewLabel(\"One\")),\n\twidget.NewAccordionItem(\"B\", widget.NewLabel(\"Two\")),\n)", "## Appending\n\nYou can append to an existing accordion using the", "ac := widget.NewAccordion(\n\twidget.NewAccordionItem(\"1\", widget.NewLabel(\"First\")),\n)\n\nac.Append(widget.NewAccordionItem(\"2\", &widget.Entry{Text: \"Second\"}))", "## Opening multiple items\n\nBy default an accordion will only allow one item to be open at a time - so opening\nanother will close the currently open item.\n\nThis can be changed using the `MultiOpen` field:", "ac := widget.NewAccordion(\n\twidget.NewAccordionItem(\"A\", widget.NewLabel(\"One\")),\n\twidget.NewAccordionItem(\"B\", widget.NewLabel(\"Two\")),\n)\n\nac.MultiOpen = true", ""},

		code: []func() fyne.CanvasObject{
			func() fyne.CanvasObject {
				ac := widget.NewAccordion(
					widget.NewAccordionItem("A", widget.NewLabel("One")),
					widget.NewAccordionItem("B", widget.NewLabel("Two")),
				)
				return ac
			},
			func() fyne.CanvasObject {
				ac := widget.NewAccordion(
					widget.NewAccordionItem("1", widget.NewLabel("First")),
				)

				ac.Append(widget.NewAccordionItem("2", &widget.Entry{Text: "Second"}))
				return ac
			},
			func() fyne.CanvasObject {
				ac := widget.NewAccordion(
					widget.NewAccordionItem("A", widget.NewLabel("One")),
					widget.NewAccordionItem("B", widget.NewLabel("Two")),
				)

				ac.MultiOpen = true
				return ac
			},
		},
	},

	"widgets/activity.md": {
		title:   "Widgets : Activity",
		content: []string{"A spinner indicating activity used in buttons etc.\n\n## Basic usage\n\nYou can create an activity \"spinner\" with the `NewActivity` constructor function.\nBefore you show it (directly or indirectly) call `Start()` to begin the animation.\n\nYou should be sure to call `Stop()` and optionally hide the object when the activity is finished.", "act := widget.NewActivity()\n\nact.Start()\n\ntime.AfterFunc(10*time.Second, func() {\n\tfyne.Do(act.Stop)\n})", ""},

		code: []func() fyne.CanvasObject{
			func() fyne.CanvasObject {
				act := widget.NewActivity()

				act.Start()

				time.AfterFunc(10*time.Second, func() {
					fyne.Do(act.Stop)
				})
				return act
			},
		},
	},

	"widgets/button.md": {
		title:   "Widgets : Button",
		content: []string{"The button widget is the basic tappable interaction for an app.\nA user tapping this will run the function passed to the constructor function. \n\n## Basic usage\n\nSimply create a button using the `NewButton` constructor function, passing in a function\nthat should run when the button is tapped.", "btn := widget.NewButton(\"Tap me\", func() {})", "If you want to use an icon in your button that is possible.\nYou can also set the label to \"\" if you want icon only!", "btn := widget.NewButtonWithIcon(\"Home\",\n    theme.HomeIcon(), func() {})", "## Disabled\n\nA button can also be disabled so that it cannot be tapped:", "btn := widget.NewButton(\"Tap me\", func() {})\nbtn.Disable()", "## Importance\n\nYou can change the colour / style of the button by setting its `Importance` value, like this:", "btn := widget.NewButton(\"Danger!\", func() {})\nbtn.Importance = widget.DangerImportance", ""},

		code: []func() fyne.CanvasObject{
			func() fyne.CanvasObject {
				btn := widget.NewButton("Tap me", func() {})
				return btn
			},
			func() fyne.CanvasObject {
				btn := widget.NewButtonWithIcon("Home",
					theme.HomeIcon(), func() {})
				return btn
			},
			func() fyne.CanvasObject {
				btn := widget.NewButton("Tap me", func() {})
				btn.Disable()
				return btn
			},
			func() fyne.CanvasObject {
				btn := widget.NewButton("Danger!", func() {})
				btn.Importance = widget.DangerImportance
				return btn
			},
		},
	},
}
