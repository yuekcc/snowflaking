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

		id, err := worker.NextID()
		So(err, ShouldBeNil)
		So(id, ShouldNotBeEmpty)

		t.Log(id)
	})
}

func TestNextIDBatch(t *testing.T) {
	Convey("生成 ID", t, func() {
		worker, err := NewIDWorkder(1000)
		So(err, ShouldBeNil)
		So(worker, ShouldNotBeNil)

		for i := 0; i < 205; i++ {
			id, _ := worker.NextID()
			t.Log(id)

			if i == 50 {
				time.Sleep(2 * time.Second)
				t.Log("wait for 2 seconds")
			}
		}
	})
}

func BenchmarkSnowFlaking(b *testing.B) {
	b.StopTimer()

	worker, _ := NewIDWorkder(1000)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		worker.NextID()
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
