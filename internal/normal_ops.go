package internal

import (
	"fmt"
	"os"
	"encoding/csv"
	"io"
	"github.com/shri-acha/lookForSecrets.git/config"
)

func PrintBucketHead (cfg *config.InputConfig){

		file,err := os.Open(cfg.FilePath)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer file.Close()

		reader := csv.NewReader(file)

		for i:=0;i<10;i++{
			// prints out the first 10 lines of the csv, if any
			line,err := reader.Read() 

			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(line)
		}
} 
