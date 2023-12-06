/*
Copyright © 2023 Calvin Colton

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "port-scan",
	Short: "Fast TCP port scanner",
	Long: `

   __/\\\\\\\\\\\\\________________________________________________          
    _\/\\\/////////\\\______________________________________________         
     _\/\\\_______\/\\\____________________________________/\\\______        
      _\/\\\\\\\\\\\\\/_______/\\\\\______/\\/\\\\\\\____/\\\\\\\\\\\_       
       _\/\\\/////////_______/\\\///\\\___\/\\\/////\\\__\////\\\////__      
        _\/\\\_______________/\\\__\//\\\__\/\\\___\///______\/\\\______     
         _\/\\\______________\//\\\__/\\\___\/\\\_____________\/\\\_/\\__    
          _\/\\\_______________\///\\\\\/____\/\\\_____________\//\\\\\___   
           _\///__________________\/////______\///_______________\/////____  
   _____/\\\\\\\\\\\_________________________________________________        
    ___/\\\/////////\\\_______________________________________________       
     __\//\\\______\///________________________________________________      
      ___\////\\\______________/\\\\\\\\___/\\\\\\\\\______/\\/\\\\\\___     
       ______\////\\\_________/\\\//////___\////////\\\____\/\\\////\\\__    
        _________\////\\\_____/\\\____________/\\\\\\\\\\___\/\\\__\//\\\_   
         __/\\\______\//\\\___\//\\\__________/\\\/////\\\___\/\\\___\/\\\_  
          _\///\\\\\\\\\\\/_____\///\\\\\\\\__\//\\\\\\\\/\\__\/\\\___\/\\\_ 
           ___\///////////_________\////////____\////////\//___\///____\///__


port-scan--short for port scanner--is a fast, lightweight CLI library that executes a TCP port scan on a list of hosts.

port-scan allows you to add, list, and delete hosts from the list.

port-scan executes a port scan on specified TCP ports. You can customize the target ports using a command line flag.
	`,
	Version: "0.1",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-cli-port-scan.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	versionTemplate := `{{printf "%s: %s - version %s\n" .Name .Short .Version}}`
	rootCmd.SetVersionTemplate(versionTemplate)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
