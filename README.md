# Go排序

（个人博客：https://www.qingbo1011.top/，博客原文[点击此处前往](https://www.qingbo1011.top/2022/04/28/Go%E6%8E%92%E5%BA%8F%E7%AE%97%E6%B3%95/)）

[八大排序算法的Golang实现](https://juejin.cn/post/7066728890575618078#heading-0)

[十个动图带你搞懂排序算法（go实现版本）](https://zhuanlan.zhihu.com/p/320419705)

[Sorting Algorithms in Go](https://dev.to/adnanbabakan/sorting-algorithms-in-go-725)

[常见排序算法知识体系详解](https://pdai.tech/md/algorithm/alg-sort-overview.html)

**[go-algorithms](https://github.com/0xAX/go-algorithms)**

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220705184830.png)

术语说明：

- n: 数据规模
- k: “桶”的个数
- In-place: 占用常数内存，不占用额外内存
- Out-place: 占用额外内存
- 稳定：如果a原本在b前面，而a=b，排序之后a仍然在b的前面
- 不稳定：如果a原本在b的前面，而a=b，排序之后a可能会出现在b的后面
- 内排序：所有排序操作都在内存中完成
- 外排序：由于数据太大，因此把数据放在磁盘中，而排序通过磁盘和内存的数据传输才能进行

# 冒泡排序

冒泡排序（Bubble Sort）：

1. 比较相邻元素，如果前者比后者大，则进行位置交换
2. 从第一个开始比较到未确定位置的最后一个，每经过一轮，则代表当前最大的数到达本次最后的位置

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220705124631.gif)

代码如下：

```go
package main

import (
   "fmt"
)

func main() {
   s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
   fmt.Println(s)
   fmt.Println(bubbleSort(s))
}

func bubbleSort(s []int) []int {
   for i := 0; i < len(s); i++ {
      for j := 0; j < len(s)-1-i; j++ {
         if s[j] > s[j+1] {
            s[j], s[j+1] = s[j+1], s[j]
         }
      }
   }
   return s
}
```

> 时间复杂度：
>
> - 最坏：如果是倒序，则需要比较 n-1 + n-2 + n-3 +...+1=n(n-1)/2次，为O(n^2)
> - 平均：O(n^2)
> - 最好：如果数组本来就是有序的，则经过一轮比较即可，共需要比较n-1次，时间复杂度是 O(n)
>
> 空间复杂度：使用常数个辅助单元O(1)
>
> 稳定性：稳定，因为 i>j时，A[i]=A[j]，不会进行交换

# 选择排序

选择排序(Selection sort)：每一次选出最小者直接交换到左侧，省出了多余的元素交换。

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220705144832.gif)

代码如下：

```go
package main

import "fmt"

func main() {
   s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
   fmt.Println(s)
   fmt.Println(selectionSort(s))
}

func selectionSort(s []int) []int {
   for i := 0; i < len(s); i++ {
      min := i
      for j := i + 1; j < len(s); j++ { // 找到索引i的右边最小值索引
         if s[j] < s[min] {
            min = j
         }
      }
      s[i], s[min] = s[min], s[i] // 交换i和min
   }
   return s
}
```

> 时间复杂度：
>
> - 最坏：O(n^2)
> - 平均：O(n^2)
> - 最好：O(n^2)
>
> 空间复杂度：O(1)
>
> 稳定性：不稳定。当数列包含多个值相等的元素时，选择排序有可能打乱它们的原有顺序。

# 插入排序

插入排序（Insertion Sort）：维护一个有序区，把元素一个一个插入到有序区的适当位置，直到所有元素有序为止。

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220705152320.gif)

代码如下：

```go
package main

import "fmt"

func main() {
   s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
   fmt.Println(s)
   fmt.Println(insertionSort(s))
}

func insertionSort(s []int) []int {
   for i := 1; i < len(s); i++ {
      for j := 0; j < i; j++ { // 索引i往i左边的区间插入
         // 最开始的区间，只有索引为0的元素(j)，索引为1(i)的元素往该区间插入
         if s[i] < s[j] {
            s[j], s[i] = s[i], s[j]
         }
      }
   }
   return s
}
```

> 时间复杂度：
>
> - 最坏：O(n^2)：切片是逆序的
> - 平均：O(n^2)
> - 最好：O(n)： 切片本身就已经是有序了 
>
> 空间复杂度：O(1)
>
> 稳定性：稳定

# 希尔排序

[What is Shell Sort](https://www.youtube.com/watch?v=ddeLSDsYVp8)

**[[算法]六分钟彻底弄懂希尔排序，简单易懂](https://www.bilibili.com/video/BV1rE411g7rW?)**

希尔排序（Shell Sort）：设置希尔增量每次折半，逐步分组进行粗调，最后进行插入排序。

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220705171851.gif)

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220705172053.png)

代码如下：

```go
package main

import "fmt"

func main() {
   s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
   fmt.Println(s)
   fmt.Println(shellSort(s))
}

func shellSort(s []int) []int {
   for gap := len(s) / 2; gap > 0; gap = gap / 2 {
      for i := gap; i < len(s); i++ {
         for j := i; j >= gap && s[j-gap] > s[j]; j = j - gap {
            s[j], s[j-gap] = s[j-gap], s[j]
         }
      }
   }
   return s
}
```

> 时间复杂度：
>
> - 最坏：O(n^2)（最坏就等于插入排序了）
> - 平均：O(nlogn)
> - 最好：O(n^1.3)  
>
> 空间复杂度：O(1)
>
> 稳定性：不稳定

# 归并排序

归并排序（Merge Sort）：最小分组比较，依次合并

1. 把长度为n的输入序列分成两个长度为n/2的子序列；
2. 对这两个子序列分别采用归并排序；
3. 将两个排序好的子序列合并成一个最终的排序序列。

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220706103240.gif)

代码如下：

```go
package main

import "fmt"

func main() {
	s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
	fmt.Println(s)
	fmt.Println(mergeSort(s))
}

func mergeSort(s []int) []int {
	if len(s) <= 1 { // 递归终止条件
		return s
	}

	middle := len(s) / 2
	left := mergeSort(s[:middle])
	right := mergeSort(s[middle:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		// 对len(left)=0和len(right)=0情况的处理
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}

		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:] // left切片去掉left[0]元素
		} else {
			result = append(result, right[0])
			right = right[1:] // right切片去掉right[0]元素
		}
	}

	return result
}
```

> 时间复杂度：
>
> - 最坏：O(nlogn)
> - 平均：O(nlogn)
> - 最好：O(nlogn)
>
> 空间复杂度： O(n)
>
> 稳定性：稳定

# 快速排序

**快速排序（Quick Sort）**：快排是面试时经常会要写的排序算法。

**[QuickSort Algorithm](https://www.youtube.com/watch?v=7h1s2SojIRw)**

分治思想：快速排序是从冒泡元素演变而来。在快速排序中，元素的比较和交换是从两端向中间进行的，较大的元素一轮就能交换到后面的位置，而较小的元素一轮就能交换到前面的位置，元素每次移动的距离较远，所以比较次数和移动次数较少，速度较快。

1.  在待排序的元素任取一个元素作为基准（**通常选第一个元素，但最好的方法是从待排序元素中随机选取一个为基准**），称为基准元素（**pivot**）
2.  将待排序的元素进行分区，比基准元素大的元素放在它的右边，比基准元素小的放在它的左边
3.  对左右两个分区重复以上步骤，直到所有的元素都是有序的

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220706113857.gif)

代码如下：

```go
package main

import "fmt"

func main() {
   s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
   fmt.Println(s)
   fmt.Println(quicksort(s))
}

func quicksort(s []int) []int {
   if len(s) < 2 {
      return s
   }
   left, right := 0, len(s)-1
   pivot := 0 // pivot可以在[0,len(s)-1]之间随机取
   s[pivot], s[right] = s[right], s[pivot]
   for i := range s {
      if s[i] < s[right] {
         s[left], s[i] = s[i], s[left]
         left++
      }
   }
   s[left], s[right] = s[right], s[left]

   quicksort(s[:left])
   quicksort(s[left+1:])

   return s
}
```

> 时间复杂度：
>
> - 最坏：O(n^2) ：选取的pivot，最终只能确定1个元素位置（如`s := []int{1,2,3,4,5,6}`，每次都选第一个元素作为pivot，那么第一次只能确定1的位置，第二次只能确定2的位置。。。）
> - 平均：O(nlogn)
> - 最好：O(nlogn)：每次选择的pivot正好是中位数
>
> 空间复杂度：
>
> - 平均：O(logn)
> - 最坏：O(n)
>
> 稳定性：不稳定，因为基准元素的比较和交换是跳跃进行的

# 堆排序

[数据结构 堆 笔记](https://www.qingbo1011.top/2021/08/05/Java%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B9%8B%E5%A0%86%E3%80%81%E6%95%A3%E5%88%97%E3%80%81%E5%B9%B6%E6%9F%A5%E9%9B%86/#java%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B9%8B%E5%A0%86%E4%BC%98%E5%85%88%E9%98%9F%E5%88%97-%E6%95%A3%E5%88%97%E5%93%88%E5%B8%8C%E8%A1%A8%E5%92%8C%E5%B9%B6%E6%9F%A5%E9%9B%86)

堆排序（Heap Sort）：利用堆这种数据结构所设计的一种排序算法。堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：即**子结点的键值或索引总是小于（或者大于）它的父节点**。（本质是完全二叉树，根节点为最大值或者最小值）

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220706125228.gif)

代码如下：

```go
package main

import "fmt"

func main() {
   s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
   fmt.Println(s)
   fmt.Println(heapSort(s))
}

func heapSort(s []int) []int {
   for i := len(s)/2 - 1; i >= 0; i-- {
      s = sift(s, i, len(s))
   }
   for i := len(s) - 1; i >= 1; i-- {
      s[0], s[i] = s[i], s[0]
      s = sift(s, 0, i)
   }
   return s
}

func sift(s []int, i int, length int) []int {
   maxChild := 0
   for i*2+1 < length {
      if (i*2+1 == length-1) || (s[i*2+1] > s[i*2+2]) {
         maxChild = i*2 + 1
      } else {
         maxChild = i*2 + 2
      }
      if s[i] < s[maxChild] {
         s[i], s[maxChild] = s[maxChild], s[i]
         i = maxChild
      } else {
         break
      }
   }
   return s
}
```

> 时间复杂度：
>
> - 最坏：O(nlogn)
> - 平均：O(nlogn)
> - 最好：O(nlogn)
>
> 空间复杂度：O(1)
>
> 稳定性：不稳定

# 计数排序

计数排序（Counting Sort）：计数排序的核心在于将输入的数据值转化为键存储在额外开辟的数组空间中。 作为一种线性时间复杂度的排序，**==计数排序要求数据必须是有确定范围的整数==**。

计数排序(Counting sort)是一种稳定的排序算法。计数排序使用一个额外的数组C，其中第i个元素是待排序数组A中值等于i的元素的个数。然后根据数组C来将A中的元素排到正确的位置。它只能对整数进行排序。

1. 找出待排序的数组中最大和最小的元素；
2. 统计数组中每个值为i的元素出现的次数，存入数组C的第i项；
3. 对所有的计数累加（从C中的第一个元素开始，每一项和前一项相加）；
4. 反向填充目标数组：将每个元素i放在新数组的第C(i)项，每放一个元素就将C(i)减去1。

说的直白一点，就是把把最小值和最大值范围内的所有可能数列举出来，然后原数组中的数字出现的个数进行计数。（思路非常简单）

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220706150829.gif)

代码如下：

```go
package main

import "fmt"

func main() {
   s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8, -3}
   fmt.Println(s)
   fmt.Println(countingSort(s))
}

func countingSort(s []int) []int {
   min, max := getMaxAndMin(s)
   bias := 0 - min // 引入bias，是为了解决s中最小值为负数的情况
   result := make([]int, 0)
   // arrayOfCounts index-bias为s中的某个数，value为该数在s中的个数
   arrayOfCounts := make([]int, max-min+1)
   for i, _ := range s { // 就算s中有负数，处理后s中的min也为0
      s[i] = s[i] + bias
   }
   for _, v := range s {
      arrayOfCounts[v]++
   }
   for i, _ := range arrayOfCounts {
      for arrayOfCounts[i] > 0 {
         result = append(result, i-bias) // 这里别忘了要减去bias
         arrayOfCounts[i]--
      }
   }
   return result
}

func getMaxAndMin(s []int) (min int, max int) {
   for _, v := range s {
      if v < min {
         min = v
      }
      if v > max {
         max = v
      }
   }
   return
}
```

> 时间复杂度：
>
> - 最坏：O(n+k)
> - 平均：O(n+k)
> - 最好：O(n+k)
>
> 空间复杂度：O(k)
>
> 稳定性：稳定

# 桶排序

桶排序（Bucket Sort）：桶排序就是把数值按照范围进行划分，把数值依次放入一个个划分的范围内，称之为 `桶`，然后在桶内进行排序，然后依次输出每个桶的值。（桶排序是计数排序的升级版。它利用了函数的映射关系，高效与否的关键就在于这个映射函数的确定。）

1. 人为设置一个BucketSize，作为每个桶所能放置多少个不同数值（例如当BucketSize==5时，该桶可以存放｛1,2,3,4,5｝这几种数字，但是容量不限，即可以存放100个3）；
2. 遍历输入数据，并且把数据一个一个放到对应的桶里去；
3. 对每个不是空的桶进行排序，**可以使用其它排序方法，也可以递归使用桶排序**；（**注意**，如果递归使用桶排序为各个桶排序，则当桶数量为1时要手动减小BucketSize增加下一循环桶的数量，否则会陷入死循环，导致内存溢出。）
4. 从不是空的桶里把排好序的数据拼接起来。 

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220706163732.png)、

代码如下：

```go
package main

import (
   "fmt"
   "math"
   "sort"
)

func main() {
   s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
   fmt.Println(s)
   fmt.Println(bucketSort(s))
}

func bucketSort(s []int) []int {
   result := make([]int, 0)
   // 1.找到最小值和最大值
   min, max := math.MaxInt, math.MinInt
   for _, v := range s {
      if v < min {
         min = v
      }
      if v > max {
         max = v
      }
   }
   // 2.初始化桶
   bucketNum := (max-min)/len(s) + 1
   bucketList := make([][]int, 0)
   for i := 0; i < bucketNum; i++ {
      bucket := make([]int, 0)
      bucketList = append(bucketList, bucket)
   }
   // 3.把元素放到桶内
   for _, v := range s {
      number := (v - min) / len(s)
      bucketList[number] = append(bucketList[number], v)
   }
   // 4.桶内进行排序
   for _, bucket := range bucketList {
      sort.Ints(bucket) // 桶内排序，可以使用其它排序方法，也可以递归使用桶排序，这里我直接用go内置的sort包
   }
   // 5.对bucketList进行拼接
   for _, bucket := range bucketList {
      for _, v := range bucket {
         result = append(result, v)
      }
   }

   return result
}
```

> 时间复杂度：
>
> - 最坏： O(n^2)
> - 平均： O(n+k)
> - 最好： O(n+k)
>
> 空间复杂度： O(n+k)
>
> 稳定性：稳定

# 基数排序

基数排序（Radix Sort）：基数排序是按照低位先排序，然后收集；再按照高位排序，然后再收集；依次类推，直到最高位。有时候有些属性是有优先级顺序的，先按低优先级排序，再按高优先级排序。最后的次序就是高优先级高的在前，高优先级相同的低优先级高的在前。基数排序基于分别排序，分别收集，所以是稳定的。

1. 取得数组中的最大数，并取得位数；
2. arr为原始数组，从最低位开始取每个位组成radix数组；
3. 对radix进行计数排序（利用计数排序适用于小范围数的特点）；

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220706172727.gif)

代码如下：

```go
package main

import (
   "bytes"
   "encoding/binary"
   "fmt"
)

const digit = 4
const maxbit = -1 << 31

func main() {
   s := []int32{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
   fmt.Println(s)
   fmt.Println(radixSort(s))
}

func radixSort(data []int32) []int32 {
   buf := bytes.NewBuffer(nil)
   ds := make([][]byte, len(data))
   for i, e := range data {
      binary.Write(buf, binary.LittleEndian, e^maxbit)
      b := make([]byte, digit)
      buf.Read(b)
      ds[i] = b
   }
   countingSort := make([][][]byte, 256)
   for i := 0; i < digit; i++ {
      for _, b := range ds {
         countingSort[b[i]] = append(countingSort[b[i]], b)
      }
      j := 0
      for k, bs := range countingSort {
         copy(ds[j:], bs)
         j += len(bs)
         countingSort[k] = bs[:0]
      }
   }
   var w int32
   for i, b := range ds {
      buf.Write(b)
      binary.Read(buf, binary.LittleEndian, &w)
      data[i] = w ^ maxbit
   }
   return data
}
```

> 时间复杂度：
>
> - 最坏：O(n*k)
> - 平均：O(n*k)
> - 最好：O(n*k)
>
> 空间复杂度：O(n+k)
>
> 稳定性：稳定

# 不同情况下排序选择

总结来看：**快速排序和希尔排序**在排序速度上表现是比较优秀的，而归并排序次之。

On average, quicksort performs better than shell sort; but shell sort is more efficient than quicksort when the given data is already/almost sorted.

Shell sort does not require stack calls, whereas quicksort does.

> 平均而言，快速排序比希尔排序性能更好; 但是，当给定数据已经或几乎排序时，希尔排序比快速排序更有效。希尔排序不需要堆栈调用（递归），而快速排序需要。

我们用Go来实际benchmark一下。（选择插入排序、快速排序和堆排序，因为Go `sort`包中主要用到了这几种排序）

根据序列元素排序情况划分：

- 完全随机的情况（random）
- 有序/逆序的情况（sorted/reverse）
- 元素重复度较高的情况（mod8）

在此基础上，还需要根据序列长度的划分（16/128/1024）

| 序列元素排序情况 | 序列长度 | 排序算法比较结果 |
| :--------------: | :------: | :--------------: |
|      random      |  短序列  |  插入>堆排>快排  |
|      sorted      |  短序列  |  插入>堆排>快排  |
|     reverse      |  短序列  |        >>        |
|       mod8       |  短序列  |        >>        |
|      random      |  中序列  |  快排>堆排>插入  |
|      sorted      |  中序列  |  插入>堆排>快排  |
|     reverse      |  中序列  |        >>        |
|       mod8       |  中序列  |        >>        |
|      random      |  长序列  |  快排>堆排>插入  |
|      sorted      |  长序列  |  插入>堆排>快排  |
|     reverse      |  长序列  |        >>        |
|       mod8       |  长序列  |        >>        |

> 总结：
>
> - 在random情况下：
>   - 插入排序在短序列中速度最快
>   - 快速排序在其他情况下速度最快
>   - 堆排序相较于快速排序差距不大
> - 在sorted情况下：
>   - 插入排序最快
> - 

- **所有短序列和元素有序情况下，插入排序性能最好**
- **在大部分情况下，快速排序有较好的综合性能**
- **几乎在任何情况下，堆排序的表现都比较稳定**

> 形象比喻：
>
> - 插入排序 ---> 单车
> - 快速排序 ---> 汽车
> - 堆排序 ---> 地铁

# Go sort包

**[go sort Documentation（官网）](https://pkg.go.dev/sort)**

**[go sort包用法](https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter03/03.1.html)**

go `sort` 包实现了四种基本排序算法：**插入排序**、**归并排序**、**堆排序**和**快速排序**。但是这四种排序方法是不公开的，它们只被用于`sort`包内部使用。所以在对数据集合排序时不必考虑应当选择哪一种排序方法，只要实现了`sort.Interface`定义的三个方法：

- 获取数据集合长度的`Len()`方法
- 比较两个元素大小的`Less()`方法
- 交换两个元素位置的`Swap()`方法

就可以顺利对数据集合进行排序。sort包会根据实际数据自动选择高效的排序算法。 除此之外，为了方便对常用数据类型的操作，sort包提供了对`[]int`切片、`[]float64`切片和`[]string`切片完整支持，主要包括：

- **对基本数据类型切片的排序支持**
- **基本数据元素查找**
- **判断基本数据类型切片是否已经排好序**
- **对排好序的数据集合逆序**

## 对切片进行排序

这里主要介绍sort包对`[]int`切片、`[]float64`切片和`[]string`切片完整支持。

> 前面已经提到，sort包原生支持`[]int`、`[]float64`和`[]string`三种内建数据类型切片的排序操作，即不必我们自己实现相关的`Len()`、`Less()`和`Swap()`方法。

对于`[]int`切片、`[]float64`切片和`[]string`切片，直接调sort包排序即可。

### []int排序和查找

对`[]int`切片排序常使用`sort.Ints()`。

代码如下：

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8, -3}
	sort.Ints(s)
	fmt.Println(s) // [-3 1 3 3 5 6 8 8 9 15 45 66]
	// 降序排序
	s = []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8, -3}
	sort.Sort(sort.Reverse(sort.IntSlice(s))) // [15 6 8 3 5 9 1 45 66 3 8 -3]
}
```

查找：

```go
package main

import (
   "fmt"
   "sort"
)

func main() {
   s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8, -3}
   sort.Ints(s)
   fmt.Println(s) // [-3 1 3 3 5 6 8 8 9 15 45 66]
   fmt.Println(sort.SearchInts(s,1))  // 1
   fmt.Println(sort.SearchInts(s,3))  // 2
}
```

**==注意==**，`SearchInts()`的使用条件为：**切片s已经升序排序** 。以下是一个错误使用的例子：

```go
s := []int{5, 2, 6, 3, 1, 4} // 未排序的切片数据
fmt.Println(sort.SearchInts(s, 2)) // 会得出错误的结果！
```

> `sort`包定义了一个`IntSlice`类型，并且实现了`sort.Interface`接口：
>
> ```go
> type IntSlice []int
> 
> func (p IntSlice) Len() int           { return len(p) }
> func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
> func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
> //IntSlice 类型定义了 Sort() 方法，包装了 sort.Sort() 函数
> func (p IntSlice) Sort() { Sort(p) }
> //IntSlice 类型定义了 SearchInts() 方法，包装了 SearchInts() 函数
> func (p IntSlice) Search(x int) int { return SearchInts(p, x) }
> ```
>
> 并且提供的`sort.Ints()`方法使用了该IntSlice类型：
>
> ```go
> func Ints(a []int) { Sort(IntSlice(a)) }
> ```
>
> 所以，对[]int切片排序更常使用`sort.Ints()`，而不是直接使用IntSlice类型。
>
> 如果要查找整数 x 在切片 a 中的位置，相对于前面提到的 Search() 方法，*sort*包提供了 SearchInts():
>
> ```go
> func SearchInts(a []int, x int) int
> ```

### []float64排序和查找

排序：

```go
package main

import (
   "fmt"
   "sort"
)

func main() {
   s := []float64{3.14, 5.12, 6.99, 3.234, 23.4762}
   sort.Float64s(s)
   fmt.Println(s) // [3.14 3.234 5.12 6.99 23.4762]
   // 降序排序
   s = []float64{3.14, 5.12, 6.99, 3.234, 23.4762}
   sort.Sort(sort.Reverse(sort.Float64Slice(s)))
   fmt.Println(s) // [23.4762 6.99 5.12 3.234 3.14]

}
```

查找：（注意同样是要求已经升序排序的切片）

```go
package main

import (
   "fmt"
   "sort"
)

func main() {
   s := []float64{3.14, 5.12, 6.99, 3.234, 23.4762}
   sort.Float64s(s)
   fmt.Println(s)                            // [3.14 3.234 5.12 6.99 23.4762]
   fmt.Println(sort.SearchFloat64s(s, 5.12)) // 2

}
```

> Float64Slice内部实现：
>
> ```go
> type Float64Slice []float64
> 
> func (p Float64Slice) Len() int           { return len(p) }
> func (p Float64Slice) Less(i, j int) bool { return p[i] < p[j] || isNaN(p[i]) && !isNaN(p[j]) }
> func (p Float64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
> func (p Float64Slice) Sort() { Sort(p) }
> func (p Float64Slice) Search(x float64) int { return SearchFloat64s(p, x) }
> ```
>
> 与 Sort()、IsSorted()、Search() 相对应的三个方法：
>
> ```go
> func Float64s(a []float64)  
> func Float64sAreSorted(a []float64) bool
> func SearchFloat64s(a []float64, x float64) int
> ```
>
> 要说明一下的是，在上面 Float64Slice 类型定义的 Less 方法中，有一个内部函数 `isNaN()`。 isNaN() 与*math*包中 `IsNaN()` 实现完全相同，*sort*包之所以不使用 math.IsNaN()，完全是基于包依赖性的考虑，应当看到，*sort*包的实现不依赖与其他任何包。

### []string排序和查找

两个string对象之间的大小比较是基于**字典序**的。

排序：

```go
package main

import (
   "fmt"
   "sort"
)

func main() {
   s := []string{"world", "b", "a", "c", "hello"}
   sort.Strings(s)
   fmt.Println(s) // [a b c hello world]
   // 降序排序
   sort.Sort(sort.Reverse(sort.StringSlice(s)))
   fmt.Println(s) // [world hello c b a]

}
```

查找：（注意同样是要求已经升序排序的切片）

```go
package main

import (
   "fmt"
   "sort"
)

func main() {
   s := []string{"world", "b", "a", "c", "hello"}
   sort.Strings(s)
   fmt.Println(s)                              // [a b c hello world]
   fmt.Println(sort.SearchStrings(s, "hello")) // 3
}
```

## 数据集合排序

参考 **[go sort包用法](https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter03/03.1.html)**

下面是一个使用 sort 包对学生成绩排序的示例：

```go
package main

import (
    "fmt"
    "sort"
)

// 学生成绩结构体
type StuScore struct {
    name  string    // 姓名
    score int   // 成绩
}

type StuScores []StuScore

//Len()
func (s StuScores) Len() int {
    return len(s)
}

//Less(): 成绩将有低到高排序
func (s StuScores) Less(i, j int) bool {
    return s[i].score < s[j].score
}

//Swap()
func (s StuScores) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func main() {
    stus := StuScores{
                {"alan", 95},
                {"hikerell", 91},
                {"acmfly", 96},
                {"leao", 90},
                }

    // 打印未排序的 stus 数据
    fmt.Println("Default:\n\t",stus)
    //StuScores 已经实现了 sort.Interface 接口 , 所以可以调用 Sort 函数进行排序
    sort.Sort(stus)
    // 判断是否已经排好顺序，将会打印 true
    fmt.Println("IS Sorted?\n\t", sort.IsSorted(stus))
    // 打印排序后的 stus 数据
    fmt.Println("Sorted:\n\t",stus)
}
```

##  []interface排序和查找

参考 **[go sort包用法](https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter03/03.1.html)**



------

# 进阶：pdqsort

之前在字节青训营的学习中，学习了从零开始打造目前业界性能一流的排序算法pdqsort（Pattern-Defeating-QuickSort）。这里简单记录一下。（将来会是Go1.19的默认排序算法）

我们都知道Go的sort包里已经给我们设计好了排序算法。那么Go的排序算法有提升空间吗？

> 什么是最快的排序算法？
>
> - Python中使用timsort
> - C++中使用introsort
> - Rust中使用pdqsort
> - Go(<=1.18)使用introsort

字节团队内部在21年10月的时候将自己的想法提供给了官方。[sort: use pdqsort #50154](https://github.com/golang/go/issues/50154)

pdqsort重新实现了Go的排序算法，在某些常见场景中比之前算法快了~10倍，成为Go1.19的默认排序算法。

|               |   Best    |      Avg      |     Worst     |
| :-----------: | :-------: | :-----------: | :-----------: |
| InsertionSort |   O(n)    |    O(n^2)     |    O(n^2)     |
|   QuickSort   | O(n*logn) |   O(n*logn)   |    O(n^2)     |
|   HeapSort    | O(n*logn) |   O(n*logn)   |   O(n*logn)   |
|  ==pdqsort==  | **O(n)**  | **O(n*logn)** | **O(n*logn)** |

pdqsort（pattern-defeating-quicksort）是一种不稳定的混合排序算法，它的不同版本被应用在C++ BOOST、Rust以及Go1.19中。它对常见的序列类型做了特殊的优化，是的在不同条件下都能拥有不错的性能。

## pdqsort - version1

结合三种排序方法的优点：

- 对于短序列（小于一定长度）我们使用插入排序
- 其他情况，使用快速排序来保证整体性能（选择首个元素作为pivot）
- 当快速排序表现不佳时，使用堆排序来保证最坏情况下时间复杂度仍然为O(n*logn)

Q&A：

- 短序列的具体长度是多少呢？
  - 12~32，在不同语言和场景下会有不同，在泛型版本根据测试选定24
- 如何得知快速排序表现不佳，以及如何及时切换到堆排序？
  - 当最终pivot的位置离序列两端很近时（距离小于length/8），判定其表现不佳，当**这种情况的次数**达到limit（即bits.Len(length)）时，切换到堆排序

 ![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220707162115.png)



## pdqsort - version2

如何让pdqsort速度更快？

- **尽量使得QuickSort的pivot为序列的中位数（改进pivot）**

关于pivot的选择：

- 使用首个元素作为pivot（最简单的方案）
  - 实现简单，但是效果往往不好，例如在sorted情况下性能很差
- 遍历数组，寻找真正的中位数
  - 遍历比对代价很高，性能不好

> 我们要在平衡`寻找pivot所需要的开销`和`pivot带来的性能优化`。

最终的解决方案是：**寻找近似中位数**！

优化pivot选择：根据序列长度的不同，来决定选择策略

- ~~短序列（<=8），选择固定元素~~（忽略，短序列直接采用插入排序了）
- 中序列（<=50），采样三个元素，取中位数
- 长序列（>50），采样九个元素，去中位数

> 这里顺带引出：pivot的采样方式使得我们有探知序列当前状态的能力。
>
> - 采样的元素都是逆序排序 ---> 序列可能已经逆序 ---> 翻转整个序列
> - 采样的元素都是顺序排序 ---> 序列可能已经有序 ---> 使用插入排序
>
> （注意：插入排序实际使用partiallnsertionSort，即有限制次数的插入排序。为了防止因为误判而all in插入排序从而影响性能。）

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220707164755.png)

> version1升级到version2优化总结：
>
> - 升级pivot选择策略（近似中位数）
> - 发现序列可能逆序，则反转序列（应对reverse场景）
> - 发现序列可能有序，使用有限制插入排序（应对sorted场景）



## pdqsort - final version

最后，还有什么场景我们没有优化？

- 短序列情况
  - 使用插入插入（version1）
- 极端情况
  - 使用堆排序保证算法可行性（verison1）
- 完全随机的情况（random）
  - 更好的pivot选择策略（version2）
- 有序/逆序的情况（sorted/reverse）
  - 根据序列状态翻转或者插入排序（version2）

最后还需要优化的场景：**如何优化重复元素很多的情况？**

> 采用pivot的时候检测重复度？
>
> - 不是很好，因为采样数量有限，不一定能采样到相同元素
>
> 解决方案：**如果两次partition生成的pivot相同，即partition进行了无效分割，此时认为pivot的值为重复元素。**

优化重复元素较多的情况：

- 当检测此时pivot和上次相同时（发生在leftSubArray），使用partitionEqual将重复元素排列在一起，减少重复元素对于pivot选择的干扰

当pivot选择策略表现不佳时：随机交换元素。（避免一些极端情况使得QuickSort总是表现不佳，以及黑客攻击情况）

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220707172423.png)

> 标黄可能发生也可能不发生。



