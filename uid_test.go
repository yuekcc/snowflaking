package snowflaking

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/yuekcc/goSnowFlake"
)

func TestNextID(t *testing.T) {
	Convey("生成 ID", t, func() {
		worker, err := NewIDWorkder(1000)
		So(err, ShouldBeNil)
		So(worker, ShouldNotBeNil)

		id, err := worker.Next()
		So(err, ShouldBeNil)
		So(id, ShouldNotBeEmpty)

		t.Log(id)
	})
}

func TestNextIDBatch(t *testing.T) {
	Convey("批量生成 ID", t, func() {
		worker, err := NewIDWorkder(1000)
		So(err, ShouldBeNil)
		So(worker, ShouldNotBeNil)

		for i := 0; i < 205; i++ {
			id, _ := worker.Next()
			t.Log(id)

			if i == 50 {
				time.Sleep(2 * time.Second)
				t.Log("wait for 2 seconds")
			}
		}
	})
}

func TestUniCheck(t *testing.T) {
	Convey("重复率检查", t, func() {
		worker, _ := NewIDWorkder(2000)
		ds := make(map[string]bool)
		countMax := 10000000
		for i := 0; i < countMax; i++ {
			id, _ := worker.Next()
			ds[id] = true
		}

		So(len(ds), ShouldEqual, countMax)
	})
}

func BenchmarkSnowFlaking(b *testing.B) {
	b.StopTimer()

	worker, _ := NewIDWorkder(1000)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		worker.Next()
	}
}

func BenchmarkSnowFlake(b *testing.B) {
	b.StopTimer()
	iw, _ := goSnowFlake.NewIdWorker(1)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		iw.NextId()
	}
}
