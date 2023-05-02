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

## Usage

```shell
$ ./passwordgen -h
Usage of ./passwordgen:
  -erase
        ./passwordgen --erase[=false]
         (default true)
  -examples
        ./passwordgen --examples
        
  -grammatical
        ./passwordgen --grammatical
        
  -grammatical-ai
        ./passwordgen --grammatical-ai
        (Requires an openai.com GPT-4 API key)
  -help
        ./passwordgen n
        Where n is the length of the password.
        Length must be the last argument.
        
  -hex
        ./passwordgen --hex
        
  -interactive
        ./passwordgen --interactive[=false]
         (default true)
  -memorable
        ./passwordgen --memorable
        
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

## Examples 

```shell
$ ./passwordgen --examples 20
┌──────────────────┬──────────────────────────────────────┐
│ --random         │ tRrLcl}Y{Mv#Tn4wmiDi                 │
├──────────────────┼──────────────────────────────────────┤
│ --hex            │ 67F2B27CE6126C82C55E                 │
├──────────────────┼──────────────────────────────────────┤
│ --word-chains    │ mullets\anatomist\sadness            │
├──────────────────┼──────────────────────────────────────┤
│ --mixed          │ Necessary/Tone2                      │
├──────────────────┼──────────────────────────────────────┤
│ --passphrases    │ dial chamomile trailing sheath straw │
├──────────────────┼──────────────────────────────────────┤
│ --memorable      │ {Rearrange|Adventurously|1940}       │
├──────────────────┼──────────────────────────────────────┤
│ --grammatical    │ Don't slug the crazy.                │
├──────────────────┼──────────────────────────────────────┤
│ --grammatical-ai │ He is like his grandfather.          │
└──────────────────┴──────────────────────────────────────┘


```

## Grammatical and Grammatical-AI

Both `--grammatical` and `grammatical-ai` will generate grammatically correct sentences for use as passphrases. But the AI sentences will make a lot more sense to you. Use of the AI option requires that you have a valid GPT-4 API key in an environment variable named `GPT_API_KEY`.

Also note the usual trailing length integer on the command line. This is required but ignored. It is a bug that needs fixing.

```shell
/passwordgen --grammatical-ai --interactive=false 20
┌───┬─────────────────────────────────────────┐
│ 0 │ Then, they didn't deduce.               │
├───┼─────────────────────────────────────────┤
│ 1 │ Is it time to sail?                     │
├───┼─────────────────────────────────────────┤
│ 2 │ Help the willing man.                   │
├───┼─────────────────────────────────────────┤
│ 3 │ Their impossible series didn't succeed. │
├───┼─────────────────────────────────────────┤
│ 4 │ That is through the window.             │
├───┼─────────────────────────────────────────┤
│ 5 │ Undermine my assumption.                │
├───┼─────────────────────────────────────────┤
│ 6 │ Don't leave their stupidity.            │
├───┼─────────────────────────────────────────┤
│ 7 │ He patted her well-worn bottom.         │
├───┼─────────────────────────────────────────┤
│ 8 │ Grill a united dish.                    │
└───┴─────────────────────────────────────────┘
```


## Memorable Passwords

```shell
/passwordgen --memorable 12
┌───┬───────────────┐
│ 0 │ [856]Unglue   │
├───┼───────────────┤
│ 1 │ 1423(Unfazed) │
├───┼───────────────┤
│ 2 │ {Landowner}4  │
├───┼───────────────┤
│ 3 │ [552]Juror    │
├───┼───────────────┤
│ 4 │ Errand{1796}  │
├───┼───────────────┤
│ 5 │ 138[Revenge]  │
├───┼───────────────┤
│ 6 │ Groove{1768}  │
├───┼───────────────┤
│ 7 │ (Getup)1691   │
├───┼───────────────┤
│ 8 │ Cadillac{780} │
└───┴───────────────┘
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
./passwordgen --mixed --interactive=false 20
┌───┬────────────────────────┐
│ 0 │ 3Avaricious|Cousin     │
├───┼────────────────────────┤
│ 1 │ Exalted+Level8         │
├───┼────────────────────────┤
│ 2 │ 7Vigorous$Permission   │
├───┼────────────────────────┤
│ 3 │ Elementary%Set2        │
├───┼────────────────────────┤
│ 4 │ 2Vibrant@Relative      │
├───┼────────────────────────┤
│ 5 │ 6Meek:Agency           │
├───┼────────────────────────┤
│ 6 │ %8Well-lit/Chain%      │
├───┼────────────────────────┤
│ 7 │ Affectionate`Designer6 │
├───┼────────────────────────┤
│ 8 │ Smooth=Arm9            │
└───┴────────────────────────┘
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

## Hexadecimal PINs

Hex pins may be 4 or more characters long.

```shell
passwordgen --hex 4
┌───┬──────┐
│ 0 │ FC70 │
├───┼──────┤
│ 1 │ DA10 │
├───┼──────┤
│ 2 │ 2DB6 │
├───┼──────┤
│ 3 │ C314 │
├───┼──────┤
│ 4 │ D186 │
├───┼──────┤
│ 5 │ 5139 │
├───┼──────┤
│ 6 │ D760 │
├───┼──────┤
│ 7 │ 5B32 │
├───┼──────┤
│ 8 │ 48F4 │
└───┴──────┘
```

# Building releases for multiple platforms

```shell
GOOS=darwin GOARCH=arm64 go build -o Release-Binaries/v1.7.x/passwordgen-v1.7.0-darwin-arm64
GOOS=darwin GOARCH=amd64 go build -o Release-Binaries/v1.7.x/passwordgen-v1.7.0-darwin-amd64
GOOS=windows GOARCH=amd64 go build -o Release-Binaries/v1.7.x/passwordgen-v1.7.0-windows-amd64.exe
GOOS=linux GOARCH=amd64 go build -o Release-Binaries/v1.7.x/passwordgen-v1.7.0-linux-amd64
```