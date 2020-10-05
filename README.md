# kawakudari-ebiten

This project implements part of the [std15.h](https://github.com/IchigoJam/c4ij/blob/master/src/std15.h) API (from [c4ij](https://github.com/IchigoJam/c4ij)) with [ebiten](https://ebiten.org), and [Kawakudari Game](https://ichigojam.github.io/print/en/KAWAKUDARI.html) on top of it.

It will allow programming for [IchigoJam](https://ichigojam.net/index-en.html)-like targets that display [IchigoJam FONT](https://mitsuji.github.io/ichigojam-font.json/) on screen using a Go programming language.
```
func (g *Game) Update(screen *ebiten.Image) error {
    if !g.running {return nil}
    if g.frame % 5 == 0 {
      if ebiten.IsKeyPressed(ebiten.KeyLeft) {
        g.x--
      }
      if ebiten.IsKeyPressed(ebiten.KeyRight) {
        g.x++
      }
      g.std15.Locate(g.x,5)
      g.std15.Putc('0')
      g.std15.Locate(rand.Int31n(32),23)
      g.std15.Putc('*')
      g.std15.Scroll(ij.Up)
      if g.std15.Scr(g.x,5) != 0 {
        g.std15.Locate(0,23)
        g.std15.Putstr("Game Over...")
        g.std15.Putnum(int32(g.frame))
        g.running = false
      }
    }
    g.frame ++
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
   g.std15.DrawScreen(screen)
}
```

## Prerequisite

* This project using programming language Go, so you need [Go](https://golang.org/doc/install) build tool properly installd to run example code.


## How to use

To just run example
```
$ go run main.go
```

To build executeble and run example
```
$ go build
$ ./kawakudari_ebiten
```


## License
[![Creative Commons License](https://i.creativecommons.org/l/by/4.0/88x31.png)](http://creativecommons.org/licenses/by/4.0/)
[CC BY](https://creativecommons.org/licenses/by/4.0/) [mitsuji.org](https://mitsuji.org)

This work is licensed under a [Creative Commons Attribution 4.0 International License](http://creativecommons.org/licenses/by/4.0/).
