package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	defaultTemplate = `<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="content-type" content="text/html; charset=utf-8" />
		<title>{{ .Title }}</title>
	</head>
	<body>
	{{ .Body }}
	</body>
</html>		
	`
)

type content struct {
	Title string
	Body  template.HTML
}

func main() {
	filename := flag.String("file", "", "Mardkdown file to preview")
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	tFname := flag.String("t", "", "Alternative template name")
	flag.Parse()

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename, *tFname, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string, tFname string, out io.Writer, skipPreview bool) error {
	input, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	htmlData, err := parseContent(input, tFname)

	if err != nil {
		return err
	}

	temp, err := ioutil.TempFile("", "mdp*.html")

	if err != nil {
		return err
	}

	if err := temp.Close(); err != nil {
		return err
	}

	outName := temp.Name()

	fmt.Fprintln(out, outName)

	if err := saveHtml(outName, htmlData); err != nil {
		return err
	}

	if skipPreview {
		return nil
	}

	defer os.Remove(outName)

	return preview(outName)
}

func parseContent(input []byte, tFname string) ([]byte, error) {
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)
	t, err := template.New("mdp").Parse(defaultTemplate)
	if err != nil {
		return nil, err
	}

	if tFname != "" {
		t, err = template.ParseFiles(tFname)
		if err != nil {
			return nil, err
		}
	}

	c := content{
		Title: "Markdown Preview Tool",
		Body:  template.HTML(body),
	}

	var buffer bytes.Buffer
	if err := t.Execute(&buffer, c); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func saveHtml(outFname string, data []byte) error {
	return ioutil.WriteFile(outFname, data, 0644)
}

func preview(fileName string) error {
	cName := ""
	cParams := []string{}
	switch runtime.GOOS {
	case "linux":
		cName = "xdg-open"
	case "windows":
		cName = "cmd.exe"
		cParams = []string{"/C", "start"}
	case "darwin":
		cName = "open"
	default:
		return fmt.Errorf("OS not supported")
	}

	cParams = append(cParams, fileName)
	cPath, err := exec.LookPath(cName)
	if err != nil {
		return err
	}

	// Delay de 2s para renderizar o html antes de deletar o arquivo.
	time.Sleep(2 * time.Second)
	return exec.Command(cPath, cParams...).Run()
}
