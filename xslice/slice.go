package xslice

import (
	"fmt"
	"reflect"
)

/**
 * @description: 通过下标删除单个切片数据
 * @return []any
 */
func DeleteByIndex[T any](s []T, index int) []T {
	// 下标越界，不进行删除操作
	if index < 0 || index >= len(s) {
		return s
	}
	copy(s[index:], s[index+1:])
	return s[:len(s)-1]
}

/**
 * @description: 通过下标删除多个切片数据
 * @return []any
 */
func DeleteByIndexs[T any](s []T, indexs []int) []T {
	var res []T
	for i, v := range s {
		if !containsIntVal(indexs, i) {
			res = append(res, v)
		}
	}
	return res
}

/**
 * @description: 通过值删除单个切片数据
 * @return []any
 */
func DeleteByValue[T any](s []T, val T) []T {
	return DeleteByValues(s, []T{val})
}

/**
 * @description: 通过下标删除多个切片数据
 * @return []any
 */
func DeleteByValues[T any](s []T, val []T) []T {
	var res []T
	for _, v := range s {
		if !containsVal(val, v) {
			res = append(res, v)
		}
	}
	return res
}

/**
 * @description:判断切片中是否存在个值，适用于int型匹配，因有针对性，性能更好
 * @param {[]int} slice
 * @param {int} value
 * @return bool
 */
func containsIntVal(slice []int, index int) bool {
	for _, val := range slice {
		if val == index {
			return true
		}
	}
	return false
}

/**
 * @description:判断切片中是否存在个值，适用于所有切片类型
 * @param {[]int} slice
 * @param {int} value
 * @return bool
 */
func containsVal[T any](slice []T, index T) bool {
	for _, val := range slice {
		if reflect.DeepEqual(val, index) {
			return true
		}
	}
	return false
}

/**
 * @description:对按照实际使用比例动态扩缩容，优化内存使用
 * @param {[]any} s 切换数组
 * @param {float32} ratio 预留分片的比例,设置时要>1
 * @param {int} minSize 切片的最小值,如果扩缩容后的size仍然小于minSize,按照minSize生成新切片
 * @return []any
 */
func SetSliceCapacity[T any](s []T, ratio float32, minSize int) ([]T, error) {
	if ratio < 1 {
		return s, fmt.Errorf("ratio need %s", ">=1")
	}
	_, l := cap(s), len(s)
	//计划扩缩容的值
	toc := int(float32(l) * ratio)

	//设置最小的阈值
	size := max(minSize, toc)

	//设置新的返回值
	ns := make([]T, 0, size)
	return append(ns, s...), nil
}
