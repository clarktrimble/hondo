
# Hondo

Smaller, Copyable Random Strings for Golang

![hondocrop](https://github.com/clarktrimble/hondo/assets/5055161/fc7517bd-7124-4c83-b36e-23863716fdf3)

## Why?

Why not just UUID?

 - UUID is hard on the eyes, especially when they gang up
 - UUID can be a pain to cut and paste
 - Shorter strings are adequate for many id's
 - Where uniqueness is required, best for back end to enforce in any case

Why a module?

 -  Single-responsibility
 -  Nice random strings come in handy often!

## Like Totally Horked

Attribution and thanks to:

Method 8 from https://stackoverflow.com/questions/22892120

## Requires go1.20 or later

Randomized seed is irresistible :)

## Example

    sessionId := hondo.Rand(7)
    fmt.Printf("Your session id will be: %s\n", sessionId)

    // -> Your session id will be: YXZo2N0

## Test, Run

    hondo % go test --count=1 ./...
    hondo % go run github.com/clarktrimble/hondo/examples/ex001

## Golang (Anti) Idioms

I dig the Golang community, but I might be a touch rouge with:

  - multi-char variable names
  - named return parameters
  - BDD/DSL testing
  - liberal use of vertical space

All in the name of readability, which of course, tends towards the subjective.

## License

This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <http://unlicense.org/>

