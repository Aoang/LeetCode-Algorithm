# [6. Z 字形变换](https://leetcode-cn.com/problems/zigzag-conversion/)

## 题目

将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
```text
P   A   H   N
A P L S I I G
Y   I   R
```

之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：
```text
string convert(string s, int numRows);
```

示例 1:
```text
输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
```

示例 2:
```text
输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G
```

## 标签

- 字符串

## 思路

1. 通过从左向右迭代字符串，确定字符位于 Z 字形图案中的哪一行
2. 按照与逐行读取 Z 字形图案相同的顺序访问字符串
