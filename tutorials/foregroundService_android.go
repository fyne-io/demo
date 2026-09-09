//go:build android

package tutorials

import (
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var timestamp time.Time
var stoppedTime time.Duration

func foregroundServiceScreen(_ fyne.Window) fyne.CanvasObject {
	label := widget.NewLabel("0s")
	currentlyRunning := isRunning()

	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			if currentlyRunning {
				elapsedTime := getElapsed()
				log.Println("elapsed: ", elapsedTime)
				fyne.Do(func() {
					label.SetText(elapsedTime)
				})

				// Repeated starts update the existing service notification;
				// the object churn stays on the Java side where the GC owns
				// it, so this is safe to call once a second.
				fyne.CurrentApp().StartForegroundService("Demo Clock", elapsedTime)
			}
			<-ticker.C
		}
	}()

	button := widget.NewButton("Start", nil)
	button.OnTapped = func() {
		fyne.CurrentApp().RequestNotificationPermission()

		if !currentlyRunning {
			fyne.CurrentApp().StartForegroundService("Demo Clock", "0s")

			ticker.Reset(time.Second)
			start(time.Now().Unix())

			fyne.Do(func() {
				label.SetText("0s")
				button.Text = "Stop"
				button.Refresh()
			})
		} else {
			fyne.CurrentApp().StopForegroundService()

			ticker.Stop()
			elapsedTime := stop()

			fyne.Do(func() {
				label.SetText(elapsedTime)
				button.Text = "Start"
				button.Refresh()
			})
		}

		currentlyRunning = !currentlyRunning
	}

	c := container.NewVBox(label, button)
	return c
}

func start(unixTimestamp int64) {
	timestamp = time.Unix(unixTimestamp, 0)
	stoppedTime = 0
}

func stop() string {
	if timestamp.IsZero() {
		stoppedTime = 0
		return "0s"
	}
	stoppedTime = time.Since(timestamp)
	timestamp = time.Time{}
	return stoppedTime.Truncate(time.Second).String()
}

func isRunning() bool {
	return !timestamp.IsZero()
}

func getElapsed() string {
	if timestamp.IsZero() {
		if stoppedTime != 0 {
			return stoppedTime.Truncate(time.Second).String()
		}

		return "0s"
	}

	return time.Since(timestamp).Truncate(time.Second).String()
}
