package stringx

type Ints struct {
	Values []int
}

func (i *Ints) Parse(input string) error {
	splitter, err := NewSplitter(input)
	if err != nil {
		return err
	}
	i.Values = splitter.Ints()
	return nil
}
