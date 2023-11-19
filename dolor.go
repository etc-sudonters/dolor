package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	flag.IntVar(&opts.Paragraphs, "p", 6, "number of paragraphs to generate")
	flag.IntVar(&opts.ParagraphSize, "P", 7, "maximum number of sentences in a paragraph")
	flag.IntVar(&opts.SentenceSize, "S", 7, "maximum number of words in a sentence")

	flag.Parse()

	for i := 0; i < opts.Paragraphs; i++ {
		fmt.Println(words.GetParagraph(opts.ParagraphSize))
	}
}

type settings struct {
	Paragraphs    int
	ParagraphSize int
	SentenceSize  int
}

var defaults = settings{
	Paragraphs:    6,
	ParagraphSize: 5,
	SentenceSize:  7,
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

type corpus []string

func (c corpus) GetWord() string {
	rand.Shuffle(len(c), func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})

	return c[0]
}

func (c corpus) GetSentence(min, max int) string {
	var result strings.Builder
	length := randInt(min, max)

	for i := 0; i < length; i++ {
		result.WriteString(c.GetWord() + " ")
	}

	sentence := result.String()
	return strings.ToUpper(sentence[:1]) + sentence[1:len(sentence)-2] + ". "
}

func (c corpus) GetParagraph(sentences int) string {
	var result strings.Builder
	length := randInt(3, sentences)

	for i := 0; i < length; i++ {
		result.WriteString(c.GetSentence(4, opts.SentenceSize))
	}

	return result.String() + "\n"
}

var opts = settings{}

var words corpus = strings.Split("sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium totam rem aperiam eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt neque porro quisquam est qui dolorem ipsum quia dolor sit amet consectetur adipisci velit sed quia non numquam  eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem ut enim ad minima veniam quis nostrum exercitationem ullam corporis suscipit laboriosam nisi ut aliquid ex ea commodi consequatur quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur vel illum qui dolorem eum fugiat quo voluptas nulla pariatur at vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium voluptatum deleniti atque corrupti quos dolores et quas molestias excepturi sint obcaecati cupiditate non provident similique sunt in culpa qui officia deserunt mollitia animi id est laborum et dolorum fuga et harum quidem reruum facilis est et expedita distinctio nam libero tempore cum soluta nobis est eligendi optio cumque nihil impedit quo minus id quod maxime placeat facere possimus omnis voluptas assumenda est omnis dolor repellendus temporibus autem quibusdam et aut officiis debitis aut rerum necessitatibus saepe eveniet ut et voluptates repudiandae sint et molestiae non recusandae itaque earum rerum hic tenetur a sapiente delectus ut aut reiciendis voluptatibus maiores alias consequatur aut perferendis doloribus asperiores repellat", " ")
