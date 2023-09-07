package xslice

import "reflect"

/**
 * @description: 通过下标删除单个分片数据
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
 * @description: 通过下标删除多个分片数据
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
 * @description: 通过值删除单个分片数据
 * @return []any
 */
func DeleteByValue[T any](s []T, val T) []T {
	return DeleteByValues(s, []T{val})
}

/**
 * @description: 通过下标删除多个分片数据
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
 * @description:判断分片中是否存在个值，适用于int型匹配，因有针对性，性能更好
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
 * @description:判断分片中是否存在个值，适用于所有分片类型
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
