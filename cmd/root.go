package cmd

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"io"
	"net/url"
	"os"
	"runtime"
	"strings"

	"github.com/GreatGodApollo/gospacebin"
	"github.com/GreatGodApollo/lunar/internal"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mitchellh/go-homedir"

	"github.com/atotto/clipboard"
)

var (
	cfgFile    string
	file       string
	instance   string
	resultBase string
	extension  string
	raw        bool
	copy       bool
	useColor   bool
)

var rootCmd = &cobra.Command{
	Version: "0.5.0",
	Use:   "lunar",
	Short: "A CLI for Spacebin",
	Long: `Lunar is a CLI for Spacebin that allows you to easily make documents.
This application can be used in a couple of different ways
to quickly create a document on an instance.

You can either pipe a document into lunar by doing:
'command | lunar'

or upload a document directly:
'lunar -f file.txt'`,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.ConfigFileUsed() != "" {
			if viper.GetString("instance") != "" && !cmd.Flags().Lookup("instance").Changed {
				instance = viper.GetString("instance")
			}
			if viper.GetString("result-url") != "" && !cmd.Flags().Lookup("result-url").Changed {
				resultBase = viper.GetString("result-url")
			}
		}
		if _, err := url.ParseRequestURI(instance); err != nil {
			fmt.Println(internal.NewMessage("Invalid instance URL!", color.FgRed))
			return
		}
		if _, err := url.ParseRequestURI(resultBase); err != nil {
			fmt.Println(internal.NewMessage("Invalid result URL!", color.FgRed))
			return
		}

		spacebin := gospacebin.NewClient(instance)

		if isPipe() {
			doc, err := postDoc(spacebin, os.Stdin)
			if err != nil {
				handleError(err)
				return
			}
			printDoc(doc)
			return
		} else {
			if file != "" {
				if !fileExists(file) {
					fmt.Println(internal.NewMessage("File does not exist!", color.FgRed))
					return
				}
			} else {
				fmt.Println(internal.NewMessage("You need to provide a file!", color.FgRed))
				return
			}
			f, err := os.Open(file)
			if err != nil {
				handleError(err)
				return
			}
			defer f.Close()
			doc, err := postDoc(spacebin, f)
			if err != nil {
				handleError(err)
				return
			}
			printDoc(doc)
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lunar.yaml)")
	rootCmd.Flags().StringVarP(&file, "file", "f", "", "the file to upload")
	rootCmd.Flags().StringVarP(&instance, "instance", "i", "https://spaceb.in", "the spacebin instance")
	rootCmd.Flags().StringVar(&resultBase, "result-url", "https://spaceb.in", "the base url for response")
	rootCmd.Flags().StringVarP(&extension, "extension", "e","none", "the file extension")
	rootCmd.Flags().BoolVarP(&raw, "raw", "r", false, "do you want the raw url")
	rootCmd.Flags().BoolVarP(&copy, "copy", "c", false, "copy the url to your clipboard")
	noColor := runtime.GOOS == "windows" || color.NoColor
	rootCmd.Flags().BoolVar(&useColor, "color", !noColor, "use color")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".lunar" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".lunar")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		color.NoColor = !useColor || color.NoColor
		fmt.Println(internal.NewMessage("Using config file:", color.FgMagenta).Then(viper.ConfigFileUsed(), color.FgGreen, color.Bold))
	}
}

// check if a file exists
func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// check if it's being used as a pipe
func isPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode() & os.ModeCharDevice == 0
}

// read the contents of a file
func readFile(r io.Reader) string {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	ctnt := []string{}
	for scanner.Scan() {
		ctnt = append(ctnt, scanner.Text())
	}
	return strings.Join(ctnt, "\n")
}

// put a doc on the instance
func postDoc(spacebin *gospacebin.Client, r io.Reader) (*gospacebin.HashDocument, error) {
	input := readFile(r)
	opts := gospacebin.NewCreateDocumentOpts(input).SetExtension(extension)
	return spacebin.CreateDocument(opts)
}

// handle an error
func handleError(err error) {
	fmt.Println(internal.NewMessage("An error occurred:", color.FgRed).Then(err.Error(), color.FgRed, color.Bold))
}

// hande the printing of a doc
func printDoc(doc *gospacebin.HashDocument) {
	uri := resultBase + "/" + doc.ID
	if raw {
		uri += "/raw"
	}
	fmt.Println(internal.NewMessage("Check out your paste @", color.FgGreen).Then(uri, color.FgCyan, color.Bold))
	if copy {
		if clipboard.Unsupported {
			fmt.Println(internal.NewMessage("Your platform is unsupported for clipboard copying.", color.FgRed))
			return
		}
		_ = clipboard.WriteAll(uri)
		fmt.Println(internal.NewMessage("URL copied to your clipboard!", color.FgGreen))
	}
}
