package structs

import (
	"fmt"
	"strconv"
)

type Business struct {
	Name string
	Industry string
	Founded uint
}

func (b *Business) GetName() {
	fmt.Println("The businesses name is: " + b.Name)
}

func (b *Business) GetIndustry() {
	fmt.Println("The businesses industry is: " + b.Industry)
}

func (b *Business) GetYearFounded() {
	fmt.Println("The businesses was founded in: " + strconv.FormatUint(uint64(b.Founded), 10))
}