A good enough for me lorem ipsum CLI tool

## Usage
```
dolor -h
Usage of dolor:
  -J string
        Used to join words together, default is a space (default " ")
  -P int
        maximum number of sentences in a paragraph (default 7)
  -S int
        maximum number of words in a sentence (default 7)
  -W int
        number of words to generate
  -p int
        number of paragraphs to generate
```

## Example

Generate a hundred placeholder blog posts with 10 random tags each

```bash
for i in {1..99}; do 
tags=$(dolor -W 10 -J ", ")
content=$(dolor -p 10 -S 20 -P 19)
echo -e "Title: LI${i}\nDate: 20${i}-07-29 18:36\nTags: ${tags}\n\n${contnet}" > content/articles/lore${i}.md;
done
```
