package heap

//构建一个大顶堆
func NewHeap(arr []int, i, length int) {
	temp := arr[i] //先保存当前需要局部调整结构的子堆的根节点

	// for循环比较一个字堆中的三个节点的大小
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if arr[k] < arr[k+1] && k+1 < length {
			// 表示左子节点小于右子节点，此时要将k指向右子节点
			k++
		}
		if arr[k] > temp {
			// 表示该子堆的根节点小于右子节点，右子节点与根节点交换位置
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}
	// 把根节点移动到最终于根节点交换位置的子节点上
	arr[i] = temp
}
