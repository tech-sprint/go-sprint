# [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:

    Input: nums = [4,5,6,7,0,1,2], target = 0
    Output: 4

Example 2:

    Input: nums = [4,5,6,7,0,1,2], target = 3
    Output: -1

# 思路

1. 这是一个二分查找的变体，所以首先要非常熟悉二分查找。把数据分为两个部分，通过判断中间值落在左右哪个范围，来移动首尾哨兵。
2. 查表法优化代码
  * tp为1，target在右半部分。
  * mp为1，nums[mid]在右半部分。
  * tm为1，target > nums[mid]。

|  tp  |  mp  |  tm  |  sum |  se   |
| ---- | ---- | ---- | ---- | ----- |
|  0   |   0  |   0  |   0  | false |
|  0   |   0  |   1  |   1  | true  |
|  0   |   1  |   0  |   2  | false |
|  0   |   1  |   1  |   3  | false |
|  1   |   0  |   0  |   4  | true  |
|  1   |   0  |   1  |   5  | true  |
|  1   |   1  |   0  |   6  | false |
|  1   |   1  |   1  |   7  | true  |
