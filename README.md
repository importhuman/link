# Exercise #4: HTML Link Parser

[![topic: html-parsing](https://img.shields.io/badge/topic-html%20parsing-green.svg?style=flat-square)](https://github.com/search?q=topic%3Ahtml-parsing+org%3Agophercises&type=Repositories)
[![topic: io-pkg](https://img.shields.io/badge/topic-io%20pkg-green.svg?style=flat-square)](https://github.com/search?q=topic%3Aio-pkg+org%3Agophercises&type=Repositories)
[![topic: recursive-funcs](https://img.shields.io/badge/topic-recursive%20funcs-green.svg?style=flat-square)](https://github.com/search?q=topic%3Arecursive-funcs+org%3Agophercises&type=Repositories)
[![topic: testing](https://img.shields.io/badge/topic-testing-green.svg?style=flat-square)](https://github.com/search?q=topic%3Atesting+org%3Agophercises&type=Repositories)

![video status: unreleased](https://img.shields.io/badge/video%20status-unreleased-red.svg?style=flat-square)
![code status: unreleased](https://img.shields.io/badge/code%20status-unreleased-red.svg?style=flat-square)

## Exercise details

In this exercise your goal is create a package that makes it easy to parse an HTML file and extract all of the links (`<a href="">...</a>` tags). For each extracted link you should return a data structure that includes both the `href`, as well as the text inside the link. Any HTML inside of the link can be stripped out, along with any extra whitespace including newlines, back-to-back spaces, etc.

Links will be nested in different HTML elements, and it is very possible that you will have to deal with HTML similar to code below.

```html
<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>
```

In situations like these we want to get output that looks roughly like:

```go
Link{
  Href: "/dog",
  Text: "Something in a span Text not in a span Bold text!",
}
```

Once you have a working program, try to write some tests for it to practice using the testing package in go.


### Notes

**1. Use the x/net/html package**

I recommend checking out the [x/net/html](https://godoc.org/golang.org/x/net/html) package for this task, which you will need to `go get`. It is provided by the Go team, but isn't included in the standard library. This makes it a little easier to parse HTML files.


**2. Ignore nested links**

You can ignore any links nested inside of another link. Eg with following HTML:

```html
<a href="#">
  Something here <a href="/dog">nested dog link</a>
</a>
```

It is okay if your code returns only the outside link, but it should still get all of the text inside the link, including the text inside the nested link.

**3. Get something working before focusing on edge-cases**

Don't worry about having perfect code. Chances are there will be a lot of edge cases here that will be kinda tricky to handle. Just try to cover the most basic use cases first and then improve on that.

**4. A few HTML examples have been provided**

I created a few simpler HTML files and included them in this repo to help with testing. They won't cover all potential use cases, but should help you start testing out your code.

## Bonus

The only bonuses here are to improve your tests and edge-case coverage.