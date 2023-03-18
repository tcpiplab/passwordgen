# passwordgen

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
+────+──────────────────────+
│ 00 │ l[TB(15k&6H3RU55iC11 │
+────+──────────────────────+
│ 01 │ Mt6Zy0Q)s@&V3w2o#0aA │
+────+──────────────────────+
│ 02 │ xkTKxfRpGAfHUTCxEqZB │
+────+──────────────────────+
│ 03 │ @DJvj!jGgG5w8uwwGtS1 │
+────+──────────────────────+
│ 04 │ {EpREgvQguP[i]7!x2OV │
+────+──────────────────────+
│ 05 │ S%cKxii2@01r1cHt8^(k │
+────+──────────────────────+
│ 06 │ ^dAK33@0NS(OWepX*#u7 │
+────+──────────────────────+
│ 07 │ 8H^8m&ifDl9KghtaymKx │
+────+──────────────────────+
│ 08 │ [dkq2#Qu6g44FdBt8f@D │
+────+──────────────────────+
Enter an integer: 5
Input has been copied to clipboard.
Waiting for 60 seconds before clearing the clipboard.
█████████████████████████████████████████████████████████████60
```

# Examples

## Usage

```shell
  -erase
        ./passwordgen -erase[=false]
         (default true)
  -help
        ./passwordgen n
        Where n is the length of the password.
        Length must be the last argument.
  -interactive
        ./passwordgen -interactive[=false]
         (default true)
  -mixed
        ./passwordgen -mixed        
  -random
        ./passwordgen -random
         (default true)
  -word-chains
        ./passwordgen -word-chains
```
## Password Chains

```shell
./passwordgen -word-chains 32
+────+──────────────────────────────────+
│ 00 │ afterlife_roundworm_winnings_div │
+────+──────────────────────────────────+
│ 01 │ posture*routing*browse*turbofan* │
+────+──────────────────────────────────+
│ 02 │ earflap=action=take=coil=constru │
+────+──────────────────────────────────+
│ 03 │ cognitive&nutty&cork&renewal&sit │
+────+──────────────────────────────────+
│ 04 │ viscous\showman\finale\abrasion\ │
+────+──────────────────────────────────+
│ 05 │ snowplow#launch#unaltered#refill │
+────+──────────────────────────────────+
│ 06 │ tamper/unloved/quickly/spoilage/ │
+────+──────────────────────────────────+
│ 07 │ shifty=regular=unscrew=confused= │
+────+──────────────────────────────────+
```

## Mixed Passwords

```shell
./passwordgen -mixed 25
+────+───────────────────────────+
│ 00 │ (rocky-outer-gliding)g)02 │
+────+───────────────────────────+
│ 01 │ {&&reunion-payee&-playpen │
+────+───────────────────────────+
│ 02 │ <<!batboy-atlas-unmoved!> │
+────+───────────────────────────+
│ 03 │ &{.uncloak-pond.-finch}&8 │
+────+───────────────────────────+
│ 04 │ 12{b{bluff-stand-trance}4 │
+────+───────────────────────────+
│ 05 │ &!?coma?-waged!-urgency&} │
+────+───────────────────────────+
│ 06 │ [quit-faction-bubble]e]12 │
+────+───────────────────────────+
│ 07 │ (?sizable-trial-squeeze?) │
+────+───────────────────────────+
```

# Building releases for multiple platforms

```shell
GOOS=darwin GOARCH=arm64 go build -o Release-Binaries/v1.2.x/passwordgen-v1.2.0-darwin-arm64
GOOS=darwin GOARCH=amd64 go build -o Release-Binaries/v1.2.x/passwordgen-v1.2.0-darwin-amd64
GOOS=windows GOARCH=amd64 go build -o Release-Binaries/v1.2.x/passwordgen-v1.2.0-windows-amd64.exe
GOOS=linux GOARCH=amd64 go build -o Release-Binaries/v1.2.x/passwordgen-v1.2.0-linux-amd64
```