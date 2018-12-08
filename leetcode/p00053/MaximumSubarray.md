# [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)

Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

    Input: [-2,1,-3,4,-1,2,1,-5,4],
    Output: 6
    Explanation: [4,-1,2,1] has the largest sum = 6.

Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

# 思路

如果最大子串包含第k个元素，那么这个最大值必然为 `f(k) = max{nums(k), nums(k) + f(k-1)}; f(1) = nums(1)`，使用这个公式吧数组 nums 映射为 f，找到f中的最大值，即为要求的解 result。

# 扩展，求出对应的子数组

把 f 的结果写出来，就能发现规律（其实在实现代码时也能注意的到）。从 result 在 f 中下标 k 开始，向回找到第一个 `nums[k - i] == f[k - i]`，则 `k - i` 为开始下标（i 从 0 开始）。即子串为 [k - i, k]。
