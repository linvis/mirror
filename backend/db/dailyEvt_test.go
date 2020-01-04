package db

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDailyEvt(t *testing.T) {
	Convey("Test DailyEvt db operate", t, func() {
		// test DailyEvt
		evt := DailyEvt{
			UserID:        255,
			EvtType:       1,
			EvtItem:       2,
			EvtDate:       11111,
			StartTimeUnix: 1234,
			EndTimeUnix:   4567,
			Duration:      111,
			Num:           5,
			Comment:       "hello test",
		}

		err := CreateDailyEvt(&evt)
		So(err, ShouldBeNil)

		all_evt, err := GetAllDailyEvt(evt.UserID)
		So(err, ShouldBeNil)
		So(len(all_evt), ShouldNotEqual, 0)

		for _, e := range all_evt {
			err = DelDailyEvt(e.EvtID)
			So(err, ShouldBeNil)
		}
	})

	Convey("Test DailyEvtType db operate", t, func() {
		evt_type := DailyEvtType{
			EvtName: "测试",
			UserID:  11,
		}
		err := CreateDailyEvtType(&evt_type)
		So(err, ShouldBeNil)

		all_evt_type, err := GetDailyEvtType(evt_type.UserID)
		So(err, ShouldBeNil)
		So(len(all_evt_type), ShouldNotEqual, 0)

		for _, e := range all_evt_type {
			err = DelDailyEvtType(e.EvtType)
			So(err, ShouldBeNil)
		}
	})

	Convey("Test DailyEvtItem db operate", t, func() {
		evt_item := DailyEvtItem{
			EvtItemName: "测试",
			EvtType:     1,
			UserID:      1,
		}
		err := CreateDailyEvtItem(&evt_item)
		So(err, ShouldBeNil)

		all_evt_item, err := GetDailyEvtItem(evt_item.UserID)
		So(err, ShouldBeNil)
		So(len(all_evt_item), ShouldNotEqual, 0)

		for _, e := range all_evt_item {
			err = DelDailyEvtItem(e.EvtItem)
			So(err, ShouldBeNil)
		}
	})
}
