package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"vietanh.com/metaverse/activity"
	"vietanh.com/metaverse/huma"
)

func addMyFamily(me *huma.Huma) {
	wife := huma.Huma{Name: "Anh", Birthday: "10/03/1995", Sex: -1}
	mom := huma.Huma{Name: "Hoan", Birthday: "05/07/1957", Sex: -1}
	dad := huma.Huma{Name: "Chau", Birthday: "22/12/1954", Sex: 1}
	daughter := huma.Huma{Name: "Linh", Birthday: "02/01/2021", Sex: -1}

	siblings := [3]huma.Huma{
		{Name: "Hue", Birthday: "13/05/1986", Sex: -1},
		{Name: "Lan", Birthday: "20/10/1981", Sex: -1},
		{Name: "Nga", Birthday: "13/05/1980", Sex: -1},
	}

	me.Relatives = append(me.Relatives, wife)
	me.Relatives = append(me.Relatives, mom)
	me.Relatives = append(me.Relatives, dad)
	me.Relatives = append(me.Relatives, daughter)
	for _, sibling := range siblings {
		me.Relatives = append(me.Relatives, sibling)
	}
}

func main() {
	dsn := "root:Yennhi0@@tcp(127.0.0.1:3306)/metaverse?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	var me huma.Huma
	db.First(&me, 1)
	// addMyFamily(&me)
	me.SelfIntroduce()

	for {
		choices := []string{"0:exit, 1:Show_today_acts", "2:Feed_Viet", "3:Sleep_Viet"}
		fmt.Println("Please chose: ", choices)
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 0:
			return
		case 1:
			me.ReportDailyActivities()
		case 2:
			me.Do(activity.Activity{Name: "Eat Breakfast", StartAt: "10:00 AM", EndAt: "10:10 AM"})
			fmt.Println("Viet is feed. Yay!")
		case 3:
			fmt.Println("Viet is slept. zzz..zz.")
		default:
			fmt.Println("I don't know what do you want to do with viet!")
		}
	}
}
