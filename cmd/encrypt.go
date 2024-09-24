package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/samuelowad/encryptor/pkg"
	"github.com/samuelowad/encryptor/util"
	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt files",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Prompt{
			Label: "Password",
			Mask:  '*',
		}

		password, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		pathPrompt := promptui.Prompt{
			Label: "File/Directory Path",
		}

		path, err := pathPrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fileData, err := util.ScanFiles(path)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		entFilePath := filepath.Join("./", ".ent.json")
		var entData []map[string]string

		if _, err := os.Stat(entFilePath); err == nil {
			file, err := os.Open(entFilePath)
			if err != nil {
				fmt.Println("Error opening .ent.json:", err)
				return
			}
			defer file.Close()

			json.NewDecoder(file).Decode(&entData)
		}

		key := util.NormalizeKey(password)
		for _, file := range fileData {
			pkg.Encrypt(file.Path, key)
			entData = append(entData, map[string]string{"key": password, "path": file.Path})
		}

		file, err := os.Create(entFilePath)
		if err != nil {
			fmt.Println("Error creating .ent.json:", err)
			return
		}
		defer file.Close()

		json.NewEncoder(file).Encode(entData)
		fmt.Printf("fileData = %v\n", fileData)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
}
