package main

import (
	"flag"
	"github.com/pkg6/gruploader-action/gruploader/clouds"
	"github.com/pkg6/gruploader-action/gruploader/flags"
	"github.com/pkg6/gruploader-action/gruploader/logger"
	"github.com/pkg6/gruploader-action/gruploader/utils"
)

var (
	args flags.Flags
	err  error
)

func init() {
	flag.StringVar(&args.Cloud, "cloud", "github", "Which cloud do you use for git?")
	flag.StringVar(&args.File, "f", "", "File path to upload.")
	flag.StringVar(&args.Repo, "repo", "", "repo, e.g., 'pkg6/gruploader-action'.")
	flag.StringVar(&args.Tag, "tag", "", "Git tag to identify a  Release in repo.")
	flag.BoolVar(&args.Overwrite, "overwrite", false, "Overwrite release asset if it's already exist.")
	flag.UintVar(&args.Retry, "retry", 2, "How many times to retry if error occur.")
}
func main() {
	flag.Parse()
	err = utils.ValidateStruct(&args)
	if err != nil {
		logger.Fatalf(err)
	}
	if err = args.DoRetryUpload(func(args *flags.Flags) error {
		cloud, err := clouds.New(args)
		if err != nil {
			return err
		}
		return cloud.Upload(args.File)
	}); err != nil {
		logger.Fatalf(err)
	}
}
