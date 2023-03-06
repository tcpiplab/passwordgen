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
│ 00 │ methaqualone:auction:malicious:d │
+────+──────────────────────────────────+
│ 01 │ supernatural\productivity\rein\s │
+────+──────────────────────────────────+
│ 02 │ bbeg:outrageous:quondam:kaph:wis │
+────+──────────────────────────────────+
│ 03 │ oceanic/mercantile/allah/jury_ri │
+────+──────────────────────────────────+
│ 04 │ oasthouse*gain*rainmeter*barbari │
+────+──────────────────────────────────+
│ 05 │ rein$jupiter$xizang$uvulopalatop │
+────+──────────────────────────────────+
│ 06 │ nonetheless|fidelity|adapt|lingo │
+────+──────────────────────────────────+
│ 07 │ a_lot-accord-beestings-audiomete │
+────+──────────────────────────────────+
```

## Mixed Passwords

Mixed passwords only work on Mac, Linux, and Unix.

```shell
./passwordgen -mixed 25
+────+───────────────────────────+
│ 00 │ *lemnad-oxysalt-choreus*) │
+────+───────────────────────────+
│ 01 │ recco_-fissive_-Giansar>~ │
+────+───────────────────────────+
│ 02 │ 0{@iffy-Marconi@-patent}3 │
+────+───────────────────────────+
│ 03 │ [aflush-hurty-worth]h]319 │
+────+───────────────────────────+
│ 04 │ $Striges$-cashew-dumple)% │
+────+───────────────────────────+
│ 05 │ $[-#$upcrane$-Nona#-sonic │
+────+───────────────────────────+
│ 06 │ #{@putois-gharial@-newel} │
+────+───────────────────────────+
│ 07 │ {Gnatho-murga-chuhra}a}27 │
+────+───────────────────────────+

```

# Building releases for multiple platforms

```shell
GOOS=darwin GOARCH=arm64 go build -o passwordgen-v1.1.0-darwin-arm64
GOOS=darwin GOARCH=amd64 go build -o passwordgen-v1.1.0-darwin-amd64
GOOS=windows GOARCH=amd64 go build -o passwordgen-v1.1.0-windows-amd64.exe
GOOS=linux GOARCH=amd64 go build -o passwordgen-v1.1.0-linux-amd64
```