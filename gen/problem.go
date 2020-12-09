package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/heldeen/aoc2020/util"
)

const (
	glueTemplate = `package cmd


import (
	"fmt"
	
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	
	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/challenge/day{{ .N }}"
)

func init() {
    day := &cobra.Command{
		Use:   "{{ .N }}",
		Short: "Problems for Day {{ .N }}",
	}

	day.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "Day {{ .N }}, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", day{{ .N }}.A(challenge.FromFile()))
		},
	})
	day.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Day {{ .N }}, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", day{{ .N }}.B(challenge.FromFile()))
		},
	})
	
	flags := day.PersistentFlags()

	flags.StringP("input", "i", "./challenge/day{{ .N }}/input.txt", "Input File to read")
	_ = viper.BindPFlags(flags)
	rootCmd.AddCommand(day)
}
`
	problemTemplate = `package day{{ .N }}

import (
    "fmt"

    "github.com/spf13/cobra"

    "github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
)

func {{ .AB }}(challenge *challenge.Input) int {
    return 0
}
`

	testTemplate = `package day{{ .N }}

import (
	"testing"

	"github.com/stretchr/testify/require"

    "github.com/heldeen/aoc2020/challenge"
)

func Test{{ .AB }}(t *testing.T) {
	input := challenge.FromLiteral("foobar")

	result := {{ .AB }}(input)

	require.Equal(t, 42, result)
}
`
)

type metadata struct {
	N  int
	AB string
}

func main() {
	date := time.Now().Format("20060102150405")
	logger, err := os.OpenFile(fmt.Sprintf("AoC.%s.log", date), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(logger)

	if len(os.Args) != 2 {
		log.Fatalf("expected 2 args but got %d", len(os.Args))
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		abort(err)
	}

	cmdP, err := util.CmdPkgPath()
	if err != nil {
		abort(err)
	}

	p, err := util.PkgPath(n)
	if err != nil {
		abort(err)
	}

	if err := os.MkdirAll(p, 0744); err != nil {
		abort(err)
	}

	funcs := template.FuncMap{
		"toLower": strings.ToLower,
	}

	gluePath := filepath.Join(cmdP, fmt.Sprintf("importDay%d.go", n))
	if _, stat := os.Stat(gluePath); stat != nil && os.IsNotExist(stat) {
		genFile(gluePath, glueTemplate, funcs, metadata{N: n, AB: "A"})
	}

	for _, ab := range []string{"A", "B"} {
		genFile(filepath.Join(p, fmt.Sprintf("%s.go", strings.ToLower(ab))), problemTemplate, funcs, metadata{N: n, AB: ab})
		genFile(filepath.Join(p, fmt.Sprintf("%s_test.go", strings.ToLower(ab))), testTemplate, funcs, metadata{N: n, AB: ab})
	}

	goimports := exec.Command("goimports", "-w", p)
	if err := goimports.Run(); err != nil {
		abort(err)
	}

	inputOutputPath := filepath.Join(p, "input.txt")
	if _, stat := os.Stat(inputOutputPath); os.IsNotExist(stat) {
		log.Printf("fetching input for day...%d\n", n)
		problemInput, err := getInput(n)
		if err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile(inputOutputPath, problemInput, 0644); err != nil {
			panic(err)
		}

	} else {
		log.Print("input already downloaded, skipping...")
	}

	log.Printf("Generated problems for day %d.", n)
}

func genFile(path, t string, funcs template.FuncMap, m metadata) {
	log.Println("creating", path)
	if _, stat := os.Stat(path); os.IsNotExist(stat) {
		t := template.Must(template.New(path).Funcs(funcs).Parse(t))
		cf, err := os.Create(path)
		if err != nil {
			log.Fatalf("creating path %v", err)
		}

		defer mustClose(cf)
		if err := t.Execute(cf, m); err != nil {
			log.Fatalf("exectute template %v", err)
		}
	} else {
		log.Println(path, "already exists, skipping...")
	}
}

func chromeCookiePath() (string, error) {
	if p, set := os.LookupEnv("CHROME_PROFILE_PATH"); set {
		return filepath.Join(p, "Cookies"), nil
	}

	if runtime.GOOS == "windows" {
		localAppData, err := os.UserCacheDir()
		return filepath.Join(localAppData, "Google", "Chrome", "User Data", "Default", "Cookies"), err
	}
	if runtime.GOOS == "linux" {
		homeDir, err := os.UserHomeDir()
		return filepath.Join(homeDir, ".config", "google-chrome", "Default", "Cookies"), err
	}

	return "", fmt.Errorf("chrome cookie path for GOOS %s not implemented, set CHROME_PROFILE_PATH instead", runtime.GOOS)
}

func getInput(day int) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2020/day/%d/input", day), nil)
	if err != nil {
		return nil, err
	}

	//cookiePath, err := chromeCookiePath()
	//if err != nil {
	//	return nil, err
	//}
	//
	//cookies, err := chrome.ReadCookies(cookiePath, kooky.Valid, kooky.Name("session"), kooky.Domain(".adventofcode.com"))
	//if err != nil {
	//	return nil, err
	//}
	//
	//if len(cookies) != 1 {
	//	return nil, log.Errorf("session cookie not found or too many results. Got %d, want 1, ensure that you are logged in", len(cookies))
	//}
	//
	//sessionToken := cookies[0].HTTPCookie()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadFile(filepath.Join(homeDir, ".aocdlconfig"))
	if err != nil {
		return nil, err
	}

	aocSession := &struct {
		S string `json:"session-cookie"`
	}{}

	err = json.Unmarshal(b, aocSession)
	if err != nil {
		return nil, err
	}

	exp, err := time.Parse(time.RFC3339, "2030-10-31T15:53:59.991Z")
	if err != nil {
		return nil, err
	}

	ck := http.Cookie{
		Name:       "session",
		Value:      aocSession.S,
		Path:       "/",
		Domain:     ".adventofcode.com",
		Expires:    exp,
		RawExpires: "",
		MaxAge:     0,
		Secure:     true,
		HttpOnly:   true,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	}

	req.AddCookie(&ck)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer mustClose(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code %s: %s", resp.Status, body)
	}

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func mustClose(c io.Closer) {
	if c == nil {
		return
	}

	if err := c.Close(); err != nil {
		log.Fatalf("error closing io.Closer: %w", err)
	}
}

func abort(err error) {
	log.Fatalf("%s\n\nsyntax: go run gen/problem.go <day> <a|b>\n", err.Error())
}
