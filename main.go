package main

import (
//    "fmt"
    "time"
    "math/rand"
    "image/color"
    "github.com/hajimehoshi/ebiten"
    "github.com/mitsuji/kawakudari_ebiten/std15"
)

type Game struct {
  std15 * std15.Std15
  frame uint32
  x int32
  running bool
}

func main() {
    rand.Seed(time.Now().UnixNano())
    ebiten.SetWindowSize(512, 384)
    ebiten.SetWindowTitle("kawakudari")
    ebiten.SetVsyncEnabled(false)
    game := &Game{
      std15 : std15.New(512, 384, 32, 24),
      frame : 0,
      x : 15,
      running : true,
    }
    if err := ebiten.RunGame(game); err != nil {
        panic(err)
    }
}

func (g *Game) Layout(outW, outH int)(screenW, screenH int) {
    return 512, 384
}

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
      g.std15.Scroll()
      g.std15.Locate(rand.Int31n(32),23)
      g.std15.Putc('*')
      if g.std15.Scr(g.x,5) != 0 {
        g.running = false
      }
    }
    g.frame ++
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
   screen.Fill(color.Black)
   g.std15.PAppletDraw(screen)
}

