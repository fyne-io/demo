package tutorials

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var tutorials = map[string]GeneratedTutorial{
	"widgets/accordion.md": GeneratedTutorial{
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

	"widgets/activity.md": GeneratedTutorial{
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

	"widgets/button.md": GeneratedTutorial{
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
