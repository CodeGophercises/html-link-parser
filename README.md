To use the parser, import `github.com/CodeGophercises/html-link-parser/parser`

The package exports the API `Parse`, which inputs an HTML byte slice.

This package extracts all the links ( including nested ones ) and returns them in the form of 
```
{
  Href: <LinkHref>
  Text: <Text inside link>
}
```

Sample HTML files are included in the dir `html-samples`.
