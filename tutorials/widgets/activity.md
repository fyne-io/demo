# Widgets : Activity

A spinner indicating activity used in buttons etc.

## Basic usage

You can create an activity "spinner" with the `NewActivity` constructor function.
Before you show it (directly or indirectly) call `Start()` to begin the animation.

You should be sure to call `Stop()` and optionally hide the object when the activity is finished.

```
act := widget.NewActivity()

act.Start()

time.AfterFunc(10*time.Second, func() {
	fyne.Do(act.Stop)
})
```
