package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"sigs.k8s.io/yaml"
)

var (
	rootCmd = &cobra.Command{
		Use:   "yaml2json",
		Short: "yaml2json",
		Run: func(cmd *cobra.Command, args []string) {
			reverse, _ := cmd.Flags().GetBool("reverse")
			if reverse {
				//json2yaml
				fmt.Println("json2yaml")
			} else {
				//yaml2json
				YamlToJson()
			}
		},
	}
)

func init() {
	rootCmd.Flags().BoolP("reverse", "r", false, "Add Floating Numbers")
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func YamlToJson() {
	filePath := os.Args[1]
	abs, err := filepath.Abs(filePath)
	Exit(err)
	file, err := ioutil.ReadFile(abs)
	Exit(err)

	split := bytes.Split(file, []byte("---"))
	jsonData := []byte{}
	newLine := []byte("\n")
	for _, item := range split {
		if bytes.Equal(item, []byte("")) {
			continue
		}
		indent, _ := yaml.YAMLToJSON(item)
		//indent, _ := json.MarshalIndent(string(i), "", "    ")

		jsonData = append(jsonData, indent...)
		jsonData = append(jsonData, newLine...)
	}

	Exit(err)
	pwd, _ := os.Getwd()

	ioutil.WriteFile(pwd+"/res.json", jsonData, fs.ModePerm)
}

func Exit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
