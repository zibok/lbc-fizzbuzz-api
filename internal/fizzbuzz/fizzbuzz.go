package fizzbuzz

import "strconv"

type Config struct {
	Limit        int
	FirstModulo  int
	SecondModulo int
	FirstWord    string
	SecondWord   string
}

func Generate(config Config) []string {
	result := make([]string, 0, config.Limit)

	for i := 1; i <= config.Limit; i++ {
		switch {
		case i%(config.FirstModulo*config.SecondModulo) == 0:
			result = append(result, config.FirstWord+config.SecondWord)
		case i%config.FirstModulo == 0:
			result = append(result, config.FirstWord)
		case i%config.SecondModulo == 0:
			result = append(result, config.SecondWord)
		default:
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}
