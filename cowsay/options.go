package cowsay

type Options struct {
	Eyes      Eyes
	Tongue    Tongue
	Thoughts  Thoughts
	File      string
	CowPath   string
	WrapWidth int
	NewLines  bool
}

func NewOptions() *Options {
	return &Options{}
}
