# passwordgen

## Build

```shell
$ git clone https://github.com/tcpiplab/passwordgen.git

$ cd passwordgen

$ go build passwordgen.go
```

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

# Usage

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
         Note that `false` behavior is not yet implemented.
```