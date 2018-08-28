# MikeBot [![Build Status](https://travis-ci.org/PapayaJuice/mikebot.svg?branch=master)](https://travis-ci.org/PapayaJuice/mikebot) [![codecov](https://codecov.io/gh/PapayaJuice/mikebot/branch/master/graph/badge.svg)](https://codecov.io/gh/PapayaJuice/mikebot)
A general Discord bot for my server

## Deploy
```bash
> cd cmd/mikebot && go build
> ./mikebot -token <token here>
```

## Commands

### !coinflip
Flips a coin and reveals the result.
```bash
> !coinflip
$ MikeBot flips a coin... It's heads!
```

### !love
Holds a user and let's them know everything is right in the world

```bash
> !love Yoshi
$ Mike holds Yoshi closely and kisses their cheek.
```

### !roll
Rolls a dice, IN PROGRESS

```bash
> !roll 1d20+2
$ Mike rolled ([13] + 2) = 15
```

### !slap
Slaps a user

```bash
> !slap a large lad
$ Mike slaps a large lad around with a large trout.
```