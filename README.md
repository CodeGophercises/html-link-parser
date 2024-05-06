To use the parser, import `github.com/CodeGophercises/html-link-parser/parser`

The package exports the API `Parse`, which inputs an HTML filename.

This package extracts all the links ( including nested ones ) and prints them out in the form of 
```
{
  href: <LinkHref>
  text: <Text inside link>
}
```

Sample HTML files are included in the dir `html-samples`.
