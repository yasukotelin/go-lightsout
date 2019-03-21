# go-lightsout

The Lights out (game) implemented with golang.

> What is Lights out? [wikipdia]("https://en.wikipedia.org/wiki/Lights_Out_(game)")

![demo1]("./public/go-lightsout-demo.gif")

## Install

```bash
go get github.com/yasukotelin/go-lightsout
```

## How to play

```bash
go-lightsout
```

If you could turn off all of lights, game clear!

|disp|mean|
|--|--|
|!|lights on|
|.|lights off|

### Options

It provides optoins `-w` `-h` that change field width and height.
Default width and height are setted 4.

```bash
go-lightsout -h 8 -w 8
```

![demo2]("./public/go-lightsout-demo2.gif")
