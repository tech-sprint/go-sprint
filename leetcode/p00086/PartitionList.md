# [Partition List](https://leetcode.com/problems/partition-list/)

Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

Example:

    Input: head = 1->4->3->2->5->2, x = 3
    Output: 1->2->2->4->3->5

# 题意

将链表使用所给值x分成两个部分，小于新的在前，大于等于x的在后，且保证这些节点的原来的相对位置。

# 思路

设置两个临时头结点（less、more），小于x的节点连在less上，其它的连在more上，最后再合并两个临时节点。
