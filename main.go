package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/russross/blackfriday.v2"
)

/**
 * Flag instructions:
 *   -i <Input Folder Path>
 *   -o <Output Folder Path>
 *   -a (Generate Atom)
 *   -r (Generate RSS)
 *   -s (Generate sitemap.xml)
 *   -d (Organize content by date)
 */
func main() {
	inputPtr := flag.String("i", "", "[required] Input Folder Path")
	outputPtr := flag.String("o", "", "[required] Output Folder Path")
	atomPtr := flag.Bool("a", false, "Generate Atom file")
	rssPtr := flag.Bool("r", false, "Generate RSS file")
	sitemapPtr := flag.Bool("s", true, "Generate Sitemap.xml file")
	datePtr := flag.Bool("d", true, "Order content by date")

	flag.Parse()

	if *inputPtr == "" || *outputPtr == "" {
		fmt.Fprintf(os.Stderr, "Input folder (-i) and output folder (-o) must be specified with a non-empty argument.\n")
		return
	}

	fmt.Println("Input:", *inputPtr)
	fmt.Println("Output:", *outputPtr)
	fmt.Println("Generate Atom:", *atomPtr)
	fmt.Println("Generate RSS:", *rssPtr)
	fmt.Println("Generate Sitemap:", *sitemapPtr)
	fmt.Println("Content ordered by date:", *datePtr)

	fmt.Println("\nFound JSON Files in input path:")

	containsIndex := false
	jsonFiles := getJsonFilesFromPath(*inputPtr)
	for _, file := range jsonFiles {
		if strings.Contains(file, "index.json") {
			if containsIndex {
				fmt.Fprintf(os.Stderr, "Only one 'index.json' file can exist.")
			}
			containsIndex = true
		}
		fmt.Println(file)
	}

	if !containsIndex {
		fmt.Fprintf(os.Stderr, "You must have one 'index.json' file.")
	}

	// Get articles
	articles := getArticlesFromJson(jsonFiles, *datePtr)
	for _, art := range articles {
		fmt.Println(art)
	}

	//Template all articles

	// Create Index JSON

	// Template index JSON
}

func getArticlesFromJson(jsonFiles []string, orderByDate bool) []Article {
	var articles []Article

	for _, file := range jsonFiles {
		if !strings.Contains(file, "index.json") {
			b, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}
			art := parseArticleFromJson(b)
			articles = append(articles, art)
		}
	}

	if orderByDate {
		sort.Slice(articles[:], func(i, j int) bool {
			return articles[i].UnixTime < articles[j].UnixTime
		})
	}
	return articles
}

func parseArticleFromJson(jsonContent []byte) Article {
	var article Article
	err := json.Unmarshal(jsonContent, &article)
	if err != nil {
		panic(err)
	}
	return article
}

func ensurePathIsValidFolder(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if fi.Mode().IsDir() {
		return true
	}
	return false
}

func ensurePathIsValidFile(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if fi.Mode().IsRegular() {
		return true
	}
	return false
}

func handleIndexJson(indexJsonPath string, inputPath string, outputPath string, articles []Article) {
	// Read IndexTemplateHtml
	// Read ElementTemplateHtml
	// For each article,
	// 		index the ElementTemplateHtml
	//		index the IndexTemplateHtml with {{elements}}, and leave a dangling {{elements}}
	// Remove the occurance of {{elements}}
	// Save IndexTemplateHtml in outputPath as index.html
}

func handleStandardArticles(articles []Article, inputPath string, outputPath string) {
	// For each article:
	// 		Read TemplateHtml
	// 		Read Markdown & convert to HTML
	// 		Template MarkdownHtml in TemplateHtml with {{content}}
	// 		Template TemplateHtml with article
	// 		Save TemplateHtml as article url
}

func getJsonFilesFromPath(root string) []string {
	var files []string
	err := filepath.Walk(root, func(fileIt string, info os.FileInfo, err error) error {
		if filepath.Ext(fileIt) == ".json" {
			files = append(files, fileIt)
		}
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get files from directory '%s'", root)
		panic(err)
	}
	return files
}

func template(template string, id string, content string) {
	//TODO
}

func mdToHtml(md []byte) []byte {
	output := blackfriday.Run(md, blackfriday.WithNoExtensions())
	return output
}
