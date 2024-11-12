package stringx

type Strings struct {
	Values []string
}

func (s *Strings) Parse(input string) error {
	splitter, err := NewSplitter(input)
	if err != nil {
		return err
	}
	s.Values = splitter.Strings()
	return nil
}
