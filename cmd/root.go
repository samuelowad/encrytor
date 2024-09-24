/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/samuelowad/encryptor/util"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "encryptor",
	Short: "A tool to encrypt and decrypt files",
	Run: func(cmd *cobra.Command, args []string) {
		welcomeAnimation()

		prompt := promptui.Select{
			Label: "Select Action",
			Items: []string{"Encrypt", "Decrypt"},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case "Encrypt":
			encryptCmd.Run(cmd, args)
		case "Decrypt":
			decryptCmd.Run(cmd, args)
		default:
			fmt.Println("Invalid selection")
		}
	},
}

func welcomeAnimation() {
	clearScreen()

	rand.Seed(time.Now().UnixNano())
	selectedDesign := util.EncryptorDesigns[rand.Intn(len(util.EncryptorDesigns))]

	colorFuncs := []func(a ...interface{}) string{
		color.New(color.FgRed).SprintFunc(),
		color.New(color.FgGreen).SprintFunc(),
		color.New(color.FgYellow).SprintFunc(),
		color.New(color.FgBlue).SprintFunc(),
		color.New(color.FgMagenta).SprintFunc(),
		color.New(color.FgCyan).SprintFunc(),
	}
	selectedColor := colorFuncs[rand.Intn(len(colorFuncs))]

	fmt.Println(selectedColor(selectedDesign))

	fmt.Print("Loading")
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(selectedColor("."))
	}
	fmt.Println(selectedColor("\nWelcome to Encryptor!"))
	time.Sleep(1 * time.Second)
	clearScreen()
}

func clearScreen() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.encryptor.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Add encrypt and decrypt commands to rootCmd
	rootCmd.AddCommand(encryptCmd)
	rootCmd.AddCommand(decryptCmd)
}
