![logo](./logo.jpg)

# cache

Why? For the glory of the God Emperor of Human Kind of course. But in all seriousness, I always liked the
caching feature in Python. You just put up a decorator and DONE. You don't have to care about
anything at all. I found that super convenient when working with various algorithms and Advent of Code
problems.

This is a nice little tool to try and imitate _SOME_ of that power. Let's see what the limitations are.

## Parameter Duplication

Turns out it's rather difficult to have a nice, user-friendly API for something that is essentially able
to call any function with any number of any typed parameters. And this was the most difficult part in this
endeavour.

In the end, I decided to duplicate the parameters in order to cache them.

## Only one return type

For now, the API only supports a single return value. This can be rather inconvenient, but if you have
multiple values, I suggest putting them into a struct. That will work nicely.

To create a cache, you have to define the type of that return value like this:

```go
    c := New[<Your Type Here>]()
```

Then, create a function that returns a `Cacheable` type with that return type. Let's say using `int`
this looks something like this:

```go
    c := New[int]()
    callCount := 0

    var f = func(a, b int) Cacheable[int] {
        return func() int {
            callCount++

            return a + b
        }
    }

    var result int
    for range 10 {
        result = c.WithCache(f(1, 2), 1, 2)
    }


    fmt.Println(result) // 3
    fmt.Println(callCount) // 1
```

And here is the second caveat.

## Duplicate Parameters

The cache needs to know about the parameters. I can't define a function with any number of parameters and any
number of types that we could than use in order for the user to pass in something like `func bla(a, b string)`.
Because the type of that function is not `func bla(args ...any)` sadly.

Therefor the way to get access to those parameters for now is to also pass them to the `WithCache` function.

If someone has a more usable idea, please don't hesitate to create an issue for it.

Which brings us to the third part.

## Hashing

_Note_: The hashing "algorithm" for generating the keys is super trivial. It wouldn't stand against millions of
values and could cause collisions quickly if values are only slightly different. The key generation depends on
`%#v` to use as a clutch.

Further, because of the string representation, the output of `Key()` could change if the struct order
is modified. Thus, it is advised to avoid serializing the output of `Key()`.
