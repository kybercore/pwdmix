package cli

// Options holds the user's password options passed in from the command line
type Options struct {
	length  int
	mode    string
	special bool
}

// NewOptions creates a new Options struct
func NewOptions() *Options {
	return &Options{}
}

func (o Options) GetLength() int {
	return o.length
}

func (o Options) GetMode() string {
	return o.mode
}

func (o Options) GetSpecial() bool {
	return o.special
}
