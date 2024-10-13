package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	flag.IntVar(&opts.Paragraphs, "p", 0, "number of paragraphs to generate")
	flag.IntVar(&opts.ParagraphSize, "P", 7, "maximum number of sentences in a paragraph")
	flag.IntVar(&opts.SentenceSize, "S", 7, "maximum number of words in a sentence")
	flag.IntVar(&opts.Words, "W", 0, "number of words to generate")
	flag.StringVar(&opts.JoinWords, "J", " ", "Used to join words together, default is a space")

	flag.Parse()

	if opts.Paragraphs == 0 && opts.Words == 0 {
		fmt.Fprint(os.Stderr, "-p or -W must be provided\n")
		flag.Usage()
		os.Exit(4)
	}

	var getter func() string
	var boundary int

	if opts.Words > 0 {
		boundary = opts.Words
		getter = func() string {
			return words.GetWord() + opts.JoinWords
		}
	} else {
		boundary = opts.Paragraphs
		getter = func() string {
			return words.GetParagraph(opts.ParagraphSize) + "\n"
		}
	}

	for i := 0; i < boundary; i++ {
		fmt.Print(getter())
	}
}

type settings struct {
	Paragraphs, ParagraphSize, SentenceSize, Words int
	JoinWords                                      string
}

var defaults = settings{
	ParagraphSize: 5,
	SentenceSize:  7,
	JoinWords:     " ",
}

type corpus []string

func (c corpus) GetWord() string {
	rand.Shuffle(len(c), func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})

	return c[0]
}

func (c corpus) GetSentence(n int) string {
	var result strings.Builder
	var length int
	for length == 0 {
		length = rand.Intn(n)
	}

	for range length {
		result.WriteString(c.GetWord() + " ")
	}

	sentence := result.String()
	return strings.ToUpper(sentence[:1]) + sentence[1:len(sentence)-2] + ". "
}

func (c corpus) GetParagraph(sentences int) string {
	var result strings.Builder
	var length int
	for length == 0 {
		length = rand.Intn(sentences)
	}

	for range length {
		result.WriteString(c.GetSentence(opts.SentenceSize))
	}

	return result.String() + "\n"
}

var opts = settings{}

var words corpus = strings.Split("sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium totam rem aperiam eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt neque porro quisquam est qui dolorem ipsum quia dolor sit amet consectetur adipisci velit sed quia non numquam  eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem ut enim ad minima veniam quis nostrum exercitationem ullam corporis suscipit laboriosam nisi ut aliquid ex ea commodi consequatur quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur vel illum qui dolorem eum fugiat quo voluptas nulla pariatur at vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium voluptatum deleniti atque corrupti quos dolores et quas molestias excepturi sint obcaecati cupiditate non provident similique sunt in culpa qui officia deserunt mollitia animi id est laborum et dolorum fuga et harum quidem reruum facilis est et expedita distinctio nam libero tempore cum soluta nobis est eligendi optio cumque nihil impedit quo minus id quod maxime placeat facere possimus omnis voluptas assumenda est omnis dolor repellendus temporibus autem quibusdam et aut officiis debitis aut rerum necessitatibus saepe eveniet ut et voluptates repudiandae sint et molestiae non recusandae itaque earum rerum hic tenetur a sapiente delectus ut aut reiciendis voluptatibus maiores alias consequatur aut perferendis doloribus asperiores repellat", " ")
