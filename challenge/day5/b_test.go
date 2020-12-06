package day5

import (
	"github.com/smartystreets/assertions"
	. "github.com/smartystreets/assertions"
	"testing"
)

func TestB(t *testing.T) {

	assert := assertions.New(t)

	l := "BFFFBBFRRR"
	assert.So(calcRow(l), ShouldEqual, 70)
	assert.So(calcColumn(l), ShouldEqual, 7)
	assert.So(calcSeatId(l), ShouldEqual, 567)

	l = "FFFBBBFRRR"
	assert.So(calcRow(l), ShouldEqual, 14)
	assert.So(calcColumn(l), ShouldEqual, 7)
	assert.So(calcSeatId(l), ShouldEqual, 119)

	l = "BBFFBBFRLL"
	assert.So(calcRow(l), ShouldEqual, 102)
	assert.So(calcColumn(l), ShouldEqual, 4)
	assert.So(calcSeatId(l), ShouldEqual, 820)

}
