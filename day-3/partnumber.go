package partnumber

import "regexp"

type Pos struct {
	Start, End, Line int
}

type Number struct {
	Raw string
	Pos Pos
}

type Symbol struct {
	Raw     string
	Pos     Pos
	Numbers []int // numbers matched
}

func ParseNumbers(data []string) []*Number {
	numbers := []*Number{}

	re, _ := regexp.Compile(`\d+`)
	for li, line := range data {
		matches := re.FindAllString(line, -1)
		if len(matches) > 0 {
			mpos := re.FindAllStringIndex(line, -1)

			for ii, match := range matches {
				nn := Number{Raw: match}
				nn.Pos = Pos{Line: li, Start: mpos[ii][0], End: mpos[ii][1]}
				numbers = append(numbers, &nn)
			}
		}
	}

	return numbers
}

func ParseSymbols(data []string) []*Symbol {
	symbols := []*Symbol{}

	re, _ := regexp.Compile(`[\%|\/|\*|\$|\&|\#|\-|\=|\@|\+]`)
	for li, line := range data {
		matches := re.FindAllString(line, -1)
		if len(matches) > 0 {
			mpos := re.FindAllStringIndex(line, -1)

			for ii, match := range matches {
				ss := Symbol{Raw: match}
				ss.Pos = Pos{Line: li, Start: mpos[ii][0], End: mpos[ii][1]}

				symbols = append(symbols, &ss)
			}
		}
	}

	return symbols
}

func Match(numbers []*Number, symbols []*Symbol) {
	for _, sym := range symbols {
		for idx, num := range numbers {
			if num.Pos.Line-1 <= sym.Pos.Line && sym.Pos.Line <= num.Pos.Line+1 {
				if num.Pos.Start-1 <= sym.Pos.Start && sym.Pos.Start <= num.Pos.End {
					sym.Numbers = append(sym.Numbers, idx)
				}
			}
		}
	}
}
