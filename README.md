# Gonsole

[![Buy me a coffee](https://badgen.net/badge/icon/buymeacoffee?icon=buymeacoffee&label)](https://www.buymeacoffee.com/zamaldinov28)
![Last stable version](https://badgen.net/github/release/zamaldinov28/gonsole)
[![Go Reference](https://pkg.go.dev/badge/github.com/zamaldinov28/gonsole.svg)](https://pkg.go.dev/github.com/zamaldinov28/goncole)

Console tools for Golang. Mostly this package is a compilation of console features, that can be easily used in projects.

## Colors

There are 256 predefined colors that available as `COLOR_*` constants. They can be used with such functions:
- `Foreground(c color) string`
- `Background(c color) string`
- `Underline(c color) string`

![Predefined colors](https://drive.google.com/uc?export=view&id=107CQjv-ftYYrhyc3VFLvWmJKwQdiH4Mn)

Also RGB colors can be used. To apply RGB style use such functions:
- `RGBForeground(r, b, b color) (string, error)`
- `RGBBackground(r, b, b color) (string, error)`
- `RGBUnderline(r, b, b color) (string, error)`

![RGB colors](https://drive.google.com/uc?export=view&id=1QtEOC0VrxI_6vObEhLr00uGpBxHPfR73)

More featured on its way!

If you need some more feature, of want to improve this package, feel free to create an issue!
