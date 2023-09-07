package xslice

import (
	"reflect"
	"testing"
)

func TestDeleteByIndex(t *testing.T) {
	a := []string{"a", "b", "c", "d"}
	shouldA := []string{"a", "c", "d"}
	res := DeleteByIndex(a, 1)
	for i, v := range res {
		if v != shouldA[i] {
			t.Errorf("Expected %s, but got %s", shouldA, res)
		}
	}
}
func TestDeleteByIndexs(t *testing.T) {
	a := []string{"a", "b", "c", "d"}
	shouldA := []string{"a", "d"}
	res := DeleteByIndexs(a, []int{1, 2})
	if !reflect.DeepEqual(shouldA, res) {
		t.Errorf("Expected %s, but got %s", shouldA, res)
	}
}

func TestDeleteByValue(t *testing.T) {
	//测试字符串分片
	a := []string{"a", "b", "c", "d"}
	shouldA := []string{"a", "c", "d"}
	res := DeleteByValue(a, "b")
	if !reflect.DeepEqual(shouldA, res) {
		t.Errorf("Expected %s, but got %s", shouldA, res)
	}
}

func TestDeleteByValues(t *testing.T) {
	//测试字符串分片
	a := []string{"a", "b", "c", "d"}
	shouldA := []string{"a", "d"}
	res := DeleteByValues(a, []string{"b", "c"})
	if !reflect.DeepEqual(shouldA, res) {
		t.Errorf("Expected %s, but got %s", shouldA, res)
	}
	//测试结构体分片
	as := []struct{ s string }{{"a"}, {"b"}, {"c"}, {"d"}}
	shouldAs := []struct{ s string }{{"a"}, {"d"}}
	ress := DeleteByValues(as, []struct{ s string }{{"b"}, {"c"}})
	if !reflect.DeepEqual(shouldAs, ress) {
		t.Errorf("Expected %s, but got %s", shouldAs, res)
	}
}

func TestSetSliceCapacity(t *testing.T) {
	a := make([]int, 0, 1000)
	for i := 0; i < 100; i++ {
		a = append(a, 1)
	}
	//常规扩缩容
	res, err := SetSliceCapacity(a, 1.2, 100)
	shouldA := 120
	if err != nil || cap(res) != shouldA {
		t.Errorf("Expected len %d, but got len %d", shouldA, cap(res))
	}

	//ratio<1
	res, err = SetSliceCapacity(a, 0.9, 100)
	if err == nil {
		t.Error("Expected err msg , but got nil")
	}

	//最小值生效
	shouldA = 200
	res, err = SetSliceCapacity(a, 1.2, 200)
	if err != nil || cap(res) != shouldA {
		t.Errorf("Expected len %d, but got len %d", shouldA, cap(res))
	}

}
