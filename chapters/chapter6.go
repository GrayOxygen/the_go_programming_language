package chapters

type Point [2]int

func (p *Point) Where() string {
	if p[0] > 0 && p[1] > 0 {
		return "eath"
	}
	return "other plant"
}
