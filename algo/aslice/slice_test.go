package aslice

import (
	"github.com/puresnr/go-cell/cast"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testSli    = make([]int, 0, 10000)
	testSliStr = make([]struct{ Name string }, 0, 10000)
)

func init() {
	for i := 0; i != 10000; i++ {
		testSli = append(testSli, i)
		testSliStr = append(testSliStr, struct{ Name string }{cast.ToString(i)})
	}
}

func TestErase(t *testing.T) {
	for i := 0; i != 20000; i++ {
		Erase(&testSli, i)
	}

	assert.Equal(t, 0, len(testSli))
}

func TestEraseIf(t *testing.T) {
	for i := 0; i != 20000; i++ {
		EraseIf(&testSliStr, func(tstr struct{ Name string }) bool { return tstr.Name == cast.ToString(i) })
	}

	assert.Equal(t, 0, len(testSliStr))
}

func TestFind(t *testing.T) {
	testStr := []string{"1", "2", "3", "4"}
	assert.Equal(t, Find(testStr, "1"), 0)
	assert.Equal(t, Find(testStr, "4"), 3)
	assert.Equal(t, Find(testStr, "5"), InvalidIdx)

	testInt := []int64{1, 2, 3, 4}
	assert.Equal(t, Find(testInt, 1), 0)
	assert.Equal(t, Find(testInt, 2), 1)
	assert.Equal(t, Find(testInt, 5), InvalidIdx)
}

func TestFindIf(t *testing.T) {
	type TestStr struct {
		Name string
		Age  int
	}

	testSli := []*TestStr{{"a", 13}, {"b", 14}, {"a", 14}}
	compare := func(cv, ct *TestStr) bool { return cv.Age == ct.Age && cv.Name == ct.Name }

	assert.Equal(t, FindIf(testSli, func(ts *TestStr) bool { return compare(ts, &TestStr{"a", 15}) }), InvalidIdx)
	assert.Equal(t, FindIf(testSli, func(ts *TestStr) bool { return compare(ts, &TestStr{"a", 13}) }), 0)
	assert.Equal(t, FindIf(testSli, func(ts *TestStr) bool { return compare(ts, &TestStr{"a", 14}) }), 2)
}

func TestExist(t *testing.T) {
	testInt := []int64{1, 2, 3, 4}
	assert.Equal(t, Exist(testInt, 1), true)
	assert.Equal(t, Exist(testInt, 2), true)
	assert.Equal(t, Exist(testInt, 5), false)
}

func TestExistIf(t *testing.T) {
	type TestStr struct {
		Name string
		Age  int
	}

	testSli := []*TestStr{{"a", 13}, {"b", 14}, {"a", 14}}
	compare := func(cv, ct *TestStr) bool { return cv.Age == ct.Age && cv.Name == ct.Name }

	assert.Equal(t, ExistIf(testSli, func(ts *TestStr) bool { return compare(ts, &TestStr{"a", 15}) }), false)
	assert.Equal(t, ExistIf(testSli, func(ts *TestStr) bool { return compare(ts, &TestStr{"a", 13}) }), true)
	assert.Equal(t, ExistIf(testSli, func(ts *TestStr) bool { return compare(ts, &TestStr{"a", 14}) }), true)
}

func TestReverse(t *testing.T) {
	testNil := []uint64{}
	testOdd := []int{1, 2, 3, 4, 5}
	testEven := []uint32{1, 2, 3, 4}

	Reverse(testNil)
	Reverse(testOdd)
	Reverse(testEven)

	assert.Equal(t, len(testNil), 0)

	for i, v := range testOdd {
		assert.Equal(t, len(testOdd)-i, v)
	}

	for i, v := range testEven {
		assert.Equal(t, uint32(len(testEven)-i), v)
	}
}

func TestReverseCopy(t *testing.T) {
	testNil := []uint64{}
	testOdd := []int{1, 2, 3, 4, 5}
	testEven := []uint32{1, 2, 3, 4}

	ctnil := ReverseCopy(testNil)
	ctodd := ReverseCopy(testOdd)
	cteven := ReverseCopy(testEven)

	assert.Equal(t, len(testNil), len(ctnil))

	for i, v := range testOdd {
		assert.Equal(t, ctodd[len(ctodd)-1-i], v)
	}

	for i, v := range testEven {
		assert.Equal(t, cteven[len(cteven)-1-i], v)
	}
}
