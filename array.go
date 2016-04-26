package goext

import "errors"

//ArrayChunk 将数组分割成二维数组
func ArrayChunk(raw interface{}, size int) ([][]interface{}, error) {

	arr := raw.([]interface{})

	if size < 1 {
		return nil, errors.New("size is zero.")
	}

	var ret [][]interface{}

	l := len(arr)

	for i := 0; i < l; i += size {

		end := i + size

		if end > l {
			end = l
		}

		ret = append(ret, arr[i:end])
	}

	return ret, nil

}
