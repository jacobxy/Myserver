package execl

import (
	"fmt"
	"xlsx"
)

func LoadAchievement(file string) (bool, error) {
	fmt.Println("file : ", file)
	f, err := xlsx.OpenFile(file)
	checkError(err)
	for _, v := range f.Sheets {
		for _, r := range v.Rows {
			for _, c := range r.Cells {
				fmt.Println(c.String())
			}
		}
	}
	return true, nil
}

func init() {
	RegisterExeclFile("achievem ent", "src/execl/Achievement.xlsx")
	RegisterExecl("achievement", LoadAchievement)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
