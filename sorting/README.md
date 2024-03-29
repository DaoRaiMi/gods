# <center> 排序(Sorting)

## 排序的稳定性
简单点说就是对于排序关键字相等的记录，在排序后是否发生位置交换。

稳定排序：不发生交换

不稳定排序：发生了位置交换

## 内排序与外排序
1. 内排序
> 在排序的整个过程中，待排序的记录全部被加载到了内存中。

2. 外排序
> 在排序的过程中，由于待排序的记录数太多，不可能全部加载到内存中，排序过程
> 中需要在内外存之间进行多次数据交换才能完成排序。
> 

## 常见的排序算法
### 冒泡排序
冒泡排序(Bubble Sort)是一种交换排序，它的基本原理就是：两两比较相邻记录的元素，
如果反序则交换，直到没有反序的元素为止。

冒泡排序的时间复杂度是O(n^2)

### 简单选择排序
简单选择排序(Simple Selection Sort)就是通过n-1次关键字比较，从n-1+1个记录
中选出关键字最小的记录，并和第i(1<=i<=n)个记录交换。

简单选择排序的时间复杂度是O(n^2)

### 插入排序
将待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。

从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。

### 快速排序
快速排序的基本思想是：先随机选择一个元素(通常是第一个或最后一个）当作轴(pivot)，通过一次排序将待排序的元素分成独立的两个部分，其中一
部分都比轴小，而另一部分的值都比轴大。然后再对两这部分执行快速排序，直到整个序列变成有序。

时间复杂度为O(logN)。快速排序是常见排序算法中最快的。