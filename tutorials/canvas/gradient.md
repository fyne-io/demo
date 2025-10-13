# Canvas : Gradient

A gradient canvas object fills an area with a gradient from one color to another.
There are two types of gradient, LinearGradient and RadialGradient.

## LinearGradient

The simplest linear gradient is created using the `NewHorizontalGradient` or `NewVerticalGradient` constructor function,
passing in the start and end colors it should use to draw.

```
grad := canvas.NewHorizontalGradient(color.NRGBA{0x80, 0, 0, 0xff}, color.NRGBA{0, 0x80, 0, 0xff})
```

### Specifying an angle

You can specify the angle of a gradient in 45 degree increments to control the angle manually.

```
grad := canvas.NewLinearGradient(color.NRGBA{0x80, 0, 0, 0xff}, color.NRGBA{0, 0x80, 0, 0xff}, 45)
```

## RadialGradient

A radial gradient draws the colour gradient from the center out to the perimeters with a circular draw.
For example:

```
grad := canvas.NewRadialGradient(color.NRGBA{0x80, 0, 0, 0xff}, color.NRGBA{0, 0x80, 0x80, 0xff})
```