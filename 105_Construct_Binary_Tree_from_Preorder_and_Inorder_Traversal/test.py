#!/usr/bin/env python3
# -*- coding: utf-8 -*-
from typing import List


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

    def __str__(self):
        if self is None:
            return ""
        return '[left: {}, val: {}, right: {}]'.format(self.left, self.val, self.right)


class Solution:
    def build_tree(self, preorder: List[int], inorder: List[int]) -> TreeNode:
        if 0 == len(preorder):
            return None

        val = preorder[0]
        for i in range(len(inorder)):
            if val == inorder[i]:
                break

        node = TreeNode(val=val)
        node.left = self.build_tree(preorder[1:1 + i], inorder[:i])
        node.right = self.build_tree(preorder[1 + i:], inorder[i + 1:])
        return node


if '__main__' == __name__:
    print(Solution().build_tree(preorder=[3, 9, 20, 15, 7], inorder=[9, 3, 15, 20, 7]))
