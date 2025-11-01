package main


import (
	"github.com/fred1268/go-clap/clap"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/shri-acha/lookForSecrets.git/internal"
	"github.com/shri-acha/lookForSecrets.git/config"
)

func main() {

	godotenv.Load(".env")

	if len(os.Args) < 2 { 
		fmt.Println("Missing arguments!")
		os.Exit(1)
	}

	var err error
	var cfg config.InputConfig
	cfg = config.InputConfig{Scan:false,ScanIdx:-1} // print mode default

	if _, err = clap.Parse(os.Args[1:], &cfg); err != nil {
		fmt.Println(err)	
		os.Exit(1)
	}

	//default mode just prints out
	// bucket's values
	if !cfg.Scan && len(cfg.FilePath)>0 { 
		internal.PrintBucketHead(&cfg)
	}
	if cfg.Scan && cfg.ScanIdx >= 0 {
		internal.ScanKeywordMatch(&cfg)
	}else if cfg.Scan && cfg.ScanIdx < 0 {
		fmt.Println("missing scan index!")
	}

}
