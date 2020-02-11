package db

// import (
// 	. "github.com/smartystreets/goconvey/convey"
// )

// // func TestSleepRecordOp(t *testing.T) {
// // 	Convey("Test DailyEvt db operate", t, func() {
// // 		// test DailyEvt
// // 		rec := SleepRecord{
// // 			UserID:        255,
// // 			Date:          11111,
// // 			StartTimeUnix: 1234,
// // 			EndTimeUnix:   4567,
// // 			Duration:      111,
// // 		}

// // 		err := NewSleepRecord(&rec)
// // 		So(err, ShouldBeNil)

// // 		allRec, err := GetSleepRecord(rec.UserID)
// // 		So(err, ShouldBeNil)
// // 		So(len(*allRec), ShouldNotEqual, 0)

// // 		for _, e := range *allRec {
// // 			err = DeleteSleepRecord(e.RecordID)
// // 			So(err, ShouldBeNil)
// // 		}
// // 	})
// // }

// func init() {
// 	InitDB("./db.config")
// }
