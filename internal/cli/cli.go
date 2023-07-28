package cli

import (
	"errors"
	"flag"
)

// ParseFlags parses the command line flags and returns them in an Options struct
func ParseFlags() (*Options, error) {
	var opts = NewOptions()
	flag.IntVar(&opts.length, "l", 12, "length of password (characters)")
	flag.StringVar(&opts.mode, "m", "a", "mode alphanumeric (a) or numeric (n)")
	flag.BoolVar(&opts.special, "s", false, "include special characters")

	flag.Parse()

	if err := ValidateOpts(opts); err != nil {
		return nil, err
	}

	return opts, nil
}

// ValidateOpts validates the options passed in from the command line
func ValidateOpts(opts *Options) error {
	if opts.GetLength() < 1 {
		return errors.New("length has to be equal to or higher than 1")
	}

	switch opts.GetMode() {
	case "a", "n", "an":
		// valid
	default:
		return errors.New("invalid mode")
	}

	return nil
}
