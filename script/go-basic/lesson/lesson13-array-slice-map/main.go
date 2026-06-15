package main

import (
	"fmt"
	"slices"
)

func main() {
	//var numbers []int
	//fmt.Println(numbers)
	//
	//// Không có kích thước
	//slice := []int{1, 2, 3, 4, 5}
	//fmt.Println(slice)
	//
	//// Có kích thước
	//array := [5]int{1}
	//fmt.Println(array)
	//
	//fmt.Println("slice co phai la Slice khong ?", reflect.TypeOf(slice).Kind() == reflect.Slice)
	//fmt.Println("array co phai la Array khong ?", reflect.TypeOf(array).Kind() == reflect.Array)

	// Mảng có tổng phần tử là 5
	//[1:4]
	// 1 -> Tính từ vị trí thứ 0 -> n - 1 (n là tổng phần tử của mảng)
	// 		=> 0 -> 5 - 1
	//		=> 0 -> 4
	// 4 -> Tính từ vị trí thứ nhất 1 -> n (n là tổng phần tử của mảng)
	//		=> 1 -> 5

	//arr := [5]int{1, 2, 3, 4, 5}
	//fmt.Println(arr)
	//
	//slice := arr[1:4]
	//fmt.Println(slice)
	//fmt.Println("slice co phai la Slice khong ?", reflect.TypeOf(slice).Kind() == reflect.Slice)
	//fmt.Println("array co phai la Array khong ?", reflect.TypeOf(arr).Kind() == reflect.Array)
	//
	//slice := make([]int, 3, 5) // [0 0 0]
	//slice[0] = 1
	//slice[1] = 2
	//slice = append(slice, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14) // [1 2 0 3]
	//
	//fmt.Println(slice)
	//fmt.Println("Chieu dai cua slice", len(slice))         // Chieu dai la 3
	//fmt.Println("Dung luong toi da cua slice", cap(slice)) // Chua toi da la 5 phan tu
	//
	//slice := []int{1, 2, 3}
	//fmt.Println(slice)
	//fmt.Println("Chieu dai cua slice", len(slice))         // Chieu dai la 3
	//fmt.Println("Dung luong toi da cua slice", cap(slice)) // Chua toi da la 5 phan tu

	//apple := []string{"Apple", "Banana", "Pear", "Cherry", "Pear", "Banana", "Cherry"}
	//
	//for i := 0; i < len(apple); i++ {
	//	fmt.Println(apple[i])
	//}
	//
	//for i, val := range apple {
	//	fmt.Println(i, val)
	//}

	//school := [][]string{
	//	{"Tuan", "Ca", "Vy"},
	//	{"Thanh", "Chim", "Zhu"},
	//}
	//
	////fmt.Println(school[1][0])
	//for _, val := range school {
	//	for i, v := range val {
	//		fmt.Println(i, v)
	//	}
	//}

	//apple := []string{"Apple", "Banana", "Pear", "Cherry", "Pear", "Banana", "Cherry"}
	//apple = append(apple, "Apple 2")
	//fmt.Println(apple)

	//apple1 := []string{"Apple", "Banana", "Pear", "Cherry", "Pear", "Banana", "Cherry"}
	//apple2 := []string{"Trung Em", "Tin", "Teo"}
	//apple3 := []string{"Tý", "Tin", "Té"}
	//
	//apple1 = append(apple1, apple2...)
	//apple1 = append(apple1, apple3...)
	//fmt.Println(apple2)
	//fmt.Println(apple1)

	// subSlice := slice[1:] // [2 3 4 5]
	// subSlice := slice[:4] // [1 2 3 4]
	// subSlice := slice[1:4] // [2 3 4]
	//slice := []int{1, 2, 3, 4, 5}
	//fmt.Println("Slice cha")
	//fmt.Println(slice)
	//fmt.Println("Length slice: ", len(slice))
	//fmt.Println("Cap slice: ", cap(slice))
	//
	//fmt.Println("=--=-==-=-=-=-=-=-=-=-=-=-==-=-=--=-=-=-=-=-")
	//subslice := slice[2:4]
	//fmt.Println("Slice con")
	//subslice = append(subslice, 90, 100)
	//fmt.Println(subslice)
	//fmt.Println("Length subslice: ", len(subslice))
	//fmt.Println("Cap subslice: ", cap(subslice))

	//** Clone: Tạo bản sao của slice **//
	//copied := slices.Clone([]int{1, 2, 3})
	//fmt.Println(copied)

	//** So sánh 2 slice có giống nhau không **//
	//compareSlice := slices.Equal([]int{1, 2, 3}, []int{1, 2, 3})
	//fmt.Println(compareSlice)

	//** Tìm vị trí đầu tiên của phần tử **//
	//findFirstPosition := slices.Index([]int{1, 2, 3, 4, 2}, 2)
	//fmt.Println(findFirstPosition)

	//** Kiểm tra phần tử có nằm trong slice ko **//
	//itemSlice := slices.Contains([]int{1, 2, 3, 4, 2}, 10)
	//fmt.Println(itemSlice)

	//** slices.Insert(s, i, v ...) Chèn phần tử vào vị trí i **//
	//insertValueAnyPosition := slices.Insert([]int{1, 2, 3, 4}, 2, 10)
	//fmt.Println(insertValueAnyPosition)

	//** slices.Delete(s, i, j) Xoá phần tử từ vị trí i đến j-1 **//
	//deleteValueAnyPosition := slices.Delete([]int{1, 2, 3, 4}, 1, 3)
	//fmt.Println(deleteValueAnyPosition)

	//** Đảo ngược slice **//
	//s := []int{1, 2, 3}
	//slices.Reverse(s)
	//fmt.Println(s)

	//** Sắp xếp slice tăng dần **//
	//s := []int{3, 1, 2}
	//slices.Sort(s)
	//fmt.Println(s)

	//** Sắp xếp theo điều kiện tuỳ chỉnh **//
	//s := []int{3, 1, 2}
	//slices.SortFunc(s, func(a, b int) int {
	//	return a - b
	//})
	//fmt.Println(s)

	itemMax := slices.Max([]int{1, 2, 3})
	fmt.Println(itemMax)

	itemMin := slices.Min([]int{1, 2, 3})
	fmt.Println(itemMin)
}
