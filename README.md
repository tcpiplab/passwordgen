# passwordgen

A command line password generator written in Golang. Works on Mac, Linux, Unix, or Windows. For now colors don't work on Windows.

[![Build Status](https://github.com/tcpiplab/passwordgen/actions/workflows/go.yml/badge.svg)](https://github.com/tcpiplab/passwordgen/actions)


## Build

```shell
$ git clone https://github.com/tcpiplab/passwordgen.git

$ cd passwordgen

$ go build passwordgen.go
```

## Or just download a binary for your platform

See the [latest release binaries](https://github.com/tcpiplab/passwordgen/releases/latest) for Mac (Intel or ARM), Linux, or Windows.

## Run

```shell
$ ./passwordgen 20
┌───┬──────────────────────┐
│ 0 │ ]Be6ITzc%*ZpZe{]XWvY │
├───┼──────────────────────┤
│ 1 │ 3E1zUVaSw4CtU13NE]H) │
├───┼──────────────────────┤
│ 2 │ %l2VuZSt8Iu96CEB#Phm │
├───┼──────────────────────┤
│ 3 │ mkEBj!t0oW5[xU[sDSad │
├───┼──────────────────────┤
│ 4 │ hF[tKgW!QXpzUvSDt8r( │
├───┼──────────────────────┤
│ 5 │ WWTh%6fEbSm5su^JDn6b │
├───┼──────────────────────┤
│ 6 │ }oHBNXgX5FKIVE@yU8uK │
├───┼──────────────────────┤
│ 7 │ 2Lw@Z8oXq@*JxV}mI9[k │
├───┼──────────────────────┤
│ 8 │ n5NQ35DJYr3QU7*J7q8z │
└───┴──────────────────────┘
Enter an integer: 5
Input has been copied to clipboard.
Waiting for 60 seconds before clearing the clipboard.
█████████████████████████████████████████████████████████████60
```

# Examples

## Usage

```shell
  -erase
        ./passwordgen --erase[=false]
         (default true)
  -help
        ./passwordgen n
        Where n is the length of the password.
        Length must be the last argument.
        
  -interactive
        ./passwordgen --interactive[=false]
        (default true)
        
  -mixed
        ./passwordgen --mixed
        
  -passphrases
        ./passwordgen --passphrases
        
  -random
        ./passwordgen --random
         (default true)
         
  -word-chains
        ./passwordgen --word-chains
```
## Password Chains

```shell
./passwordgen -word-chains 32
┌───┬──────────────────────────────────────────────┐
│ 0 │ washable_breeder_plexiglas_savage_fritter    │
├───┼──────────────────────────────────────────────┤
│ 1 │ panther^repossess^shifter^gopher^yield       │
├───┼──────────────────────────────────────────────┤
│ 2 │ bleep+footprint+culminate+cavalier+factsheet │
├───┼──────────────────────────────────────────────┤
│ 3 │ conform#powwow#flashback#acronym#reburial    │
├───┼──────────────────────────────────────────────┤
│ 4 │ unexpired|glutinous|grape|fructose|comic     │
├───┼──────────────────────────────────────────────┤
│ 5 │ aeration$finisher$unmade$naturist$paradox    │
├───┼──────────────────────────────────────────────┤
│ 6 │ wistful&translate&sherry&selection&engulf    │
├───┼──────────────────────────────────────────────┤
│ 7 │ anatomy/slicing/upfront/engraved/haunt       │
├───┼──────────────────────────────────────────────┤
│ 8 │ yield\grant\myspace\dusk\skylight\subsoil    │
└───┴──────────────────────────────────────────────┘
```

## Mixed Passwords

```shell
./passwordgen -mixed 25
┌───┬──────────────────────────────────┐
│ 0 │ @{*&^approach^-marshy&-voice*}@  │
├───┼──────────────────────────────────┤
│ 1 │ .{~#flagship#-pebbly-power~}.    │
├───┼──────────────────────────────────┤
│ 2 │ [-seldom--rendering-survival]    │
├───┼──────────────────────────────────┤
│ 3 │ [*@scored@-unpaved-phrasing*]    │
├───┼──────────────────────────────────┤
│ 4 │ ^<^-carport--flatbed-discover^>^ │
├───┼──────────────────────────────────┤
│ 5 │ <^boaster-engraving^-email>      │
├───┼──────────────────────────────────┤
│ 6 │ 1<e<endorphin-muck-grimy>        │
├───┼──────────────────────────────────┤
│ 7 │ (residual-duplicity-reprogram)   │
├───┼──────────────────────────────────┤
│ 8 │ (outback-unlighted-aflutter)     │
└───┴──────────────────────────────────┘
```

## Passphrases

The example below shows use of the `-passphrases` feature as well as the optional feature `-interactive=false` so that a list of passphrases is simply printed to the screen and the program exits. This example is run in Powershell on Windows 10 but works on all platforms.

```shell
PS C:\Users\somebody\Downloads> .\passwordgen-v1.3.0-windows-amd64.exe -passphrases -interactive=false
┌────┬───────────────────────────────────────────────┐
│  0 │ active stalling dubbed almighty entity        │
├────┼───────────────────────────────────────────────┤
│  1 │ tying wireless relish levitate outwit         │
├────┼───────────────────────────────────────────────┤
│  2 │ unreeling angelic camper augmented hardhat    │
├────┼───────────────────────────────────────────────┤
│  3 │ had humble polish legume external             │
├────┼───────────────────────────────────────────────┤
│  4 │ regime twilight risk outburst overarch        │
├────┼───────────────────────────────────────────────┤
│  5 │ hazing move uncle clustered dehydrate         │
├────┼───────────────────────────────────────────────┤
│  6 │ sizzling singer predict surplus debtor        │
├────┼───────────────────────────────────────────────┤
│  7 │ division imprudent tapeless unbounded console │
├────┼───────────────────────────────────────────────┤
│  8 │ replica cabbage regress detector purifier     │
├────┼───────────────────────────────────────────────┤
│  9 │ denote ancient customize tidal puppy          │
├────┼───────────────────────────────────────────────┤
│ 10 │ sworn scraggly sandstorm crayon untapped      │
├────┼───────────────────────────────────────────────┤
│ 11 │ deceiver bunny subpanel decathlon lifting     │
├────┼───────────────────────────────────────────────┤
│ 12 │ unvalued could easter polymer unlimited       │
├────┼───────────────────────────────────────────────┤
│ 13 │ legroom caretaker reference frantic genetics  │
└────┴───────────────────────────────────────────────┘
```

# Building releases for multiple platforms

```shell
GOOS=darwin GOARCH=arm64 go build -o Release-Binaries/v1.4.x/passwordgen-v1.4.0-darwin-arm64
GOOS=darwin GOARCH=amd64 go build -o Release-Binaries/v1.4.x/passwordgen-v1.4.0-darwin-amd64
GOOS=windows GOARCH=amd64 go build -o Release-Binaries/v1.4.x/passwordgen-v1.4.0-windows-amd64.exe
GOOS=linux GOARCH=amd64 go build -o Release-Binaries/v1.4.x/passwordgen-v1.4.0-linux-amd64
```