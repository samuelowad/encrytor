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

// Removed unused import: "os/user"

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt files",
	Run: func(cmd *cobra.Command, args []string) {
		pathPrompt := promptui.Prompt{
			Label: "File/Directory Path",
		}

		path, err := pathPrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		prompt := promptui.Select{
			Label: "Use Passkey or Master Key",
			Items: []string{"Passkey", "Master Key"},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		var password string
		if result == "Master Key" {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				fmt.Println("Error fetching home directory:", err)
				return
			}

			masterFilePath := filepath.Join(homeDir, ".ent_master.json")
			file, err := os.Open(masterFilePath)
			if err != nil {
				fmt.Println("Error opening master file:", err)
				return
			}
			defer file.Close()

			var masterData map[string]string
			json.NewDecoder(file).Decode(&masterData)

			prompt := promptui.Prompt{
				Label: "Master Password",
				Mask:  '*',
			}

			masterPassword, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			if masterPassword != masterData["master_password"] {
				fmt.Println("Invalid master password")
				return
			}

			entFilePath := filepath.Join(path, ".ent.json")
			file, err = os.Open(entFilePath)
			if err != nil {
				fmt.Println("Error opening .ent.json:", err)
				return
			}
			defer file.Close()

			var entData []map[string]string
			json.NewDecoder(file).Decode(&entData)

			for _, entry := range entData {
				if entry["path"] == path {
					password = entry["key"]
					break
				}
			}
		} else {
			prompt := promptui.Prompt{
				Label: "Password",
				Mask:  '*',
			}

			password, err = prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
		}

		// Decrypt logic here using the password
		if password == "" {
			fmt.Println("No password found or provided")
			return
		}

		// Use util.ScanFiles to get all files in the directory
		fileData, err := util.ScanFiles(path)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		key := util.NormalizeKey(password)
		for _, file := range fileData {
			pkg.Decrypt(file.Path, key)
		}

		fmt.Printf("Files at %s decrypted successfully\n", path)
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
}
