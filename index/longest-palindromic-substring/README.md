# [5. 最长回文子串](https://leetcode.com/problems/longest-palindromic-substring/)

## 题目

给定一个字符串 `s`，找到 `s` 中最长的回文子串。你可以假设 `s` 的最大长度为 1000。

示例 1:
```text
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
```

示例 2:
```text
输入: "cbbd"
输出: "bb"
```

## 标签

- 字符串
- 动态规划

## 解题思路

马拉车算法 Manacher‘s Algorithm 是用来查找一个字符串的最长回文子串的线性方法，它将时间复杂度提升到了线性。