package huma

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"vietanh.com/metaverse/activity"
)

var Emotions = []string{"Happy", "Sad", "Blissful", "Angry", "Peaceful"}

var Sex = [3]int8{1, 0, -1} //1 male, -1 female, 0 gay

type Huma struct {
	gorm.Model
	Birthday        string
	Name            string
	Sex             int8
	Relatives       []Huma              `gorm:"-"`
	Friends         []Huma              `gorm:"-"`
	todayActivities []activity.Activity `gorm:"-"`
}

func (huma *Huma) Do(activity activity.Activity) {
	huma.todayActivities = append(huma.todayActivities, activity)
}

func (huma *Huma) ReportDailyActivities() {
	for i, act := range huma.todayActivities {
		fmt.Println(i+1, ". ", act.Name, " Start at", act.StartAt, ", End at", act.EndAt)
	}
}

func (huma *Huma) GetFirstName() string {
	return strings.Fields(huma.Name)[0]
}

func (huma *Huma) SelfIntroduce() {
	fmt.Println("Hi, my name is", huma.Name, "Sex: ", huma.Sex, ". I was born in ", huma.Birthday)
	fmt.Println("My first name is", huma.GetFirstName())
	fmt.Println("My relatives are:", huma.GetRelativeNames())
}

func (huma *Huma) GetRelativeNames() string {
	var names string
	for _, relative := range huma.Relatives {
		names = names + ", " + relative.Name
	}
	return names
}

func (huma *Huma) Speak(sentence string) {
	fmt.Println(huma.Name, " says: ", sentence)
}
