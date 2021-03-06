# [1. 两数之和](https://leetcode-cn.com/problems/two-sum/)

## 题目

给定一个整数数组 `nums` 和一个目标值 `target`，请你在该数组中找出和为目标值的那 *两个* 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:
```text
给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
```

## 标签

- 数组
- 哈希表

## 思路

循环肯定避免不了，而暴力解法的复杂度太高了。
题目中说明假设每种输入只有一个答案，那么 map 用在这个地方应该刚刚好。

1. 直接用两个嵌套循环遍历数组，然后和目标值进行对比。
2. 用 map 存储元素的值、索引，检查每个元素所对应的目标元素是否在 map 中。
