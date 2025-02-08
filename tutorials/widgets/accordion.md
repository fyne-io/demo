# Widgets : Accordion

Expand or collapse content panels.

## Basic usage

A basic accordion can be created with `NewAccordion` passing in items created using
the `NewAccordionItem` helper, as shown:

```
ac := widget.NewAccordion(
	widget.NewAccordionItem("A", widget.NewLabel("One")),
	widget.NewAccordionItem("B", widget.NewLabel("Two")),
)
```

## Appending

You can append to an existing accordion using the 

```
ac := widget.NewAccordion(
	widget.NewAccordionItem("1", widget.NewLabel("First")),
)

ac.Append(widget.NewAccordionItem("2", &widget.Entry{Text: "Second"}))
```

## Opening multiple items

By default an accordion will only allow one item to be open at a time - so opening
another will close the currently open item.

This can be changed using the `MultiOpen` field:

```
ac := widget.NewAccordion(
	widget.NewAccordionItem("A", widget.NewLabel("One")),
	widget.NewAccordionItem("B", widget.NewLabel("Two")),
)

ac.MultiOpen = true
```