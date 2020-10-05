#!/usr/bin/env python3
# -*- coding: utf-8 -*-


class Node:
    """
    Node is used to construct a node of linked list
    """

    def __init__(self, val: int, next: '__class__' = None):
        """
        construct a node
        :param val: value of current node
        :param next: point to next node
        """
        self.val = val
        self.next = next


class MyLinkedList:
    """
    MyLinkedList is used to construct a linked list
    """

    def __init__(self, head: Node = None, tail: Node = None, size: int = 0):
        """
        construct a empty linked list
        :param head: head of linked list
        :param tail: tail of linked list
        :param size: size of linked list
        """
        self.head = head
        self.tail = tail
        self.size = size

    def get(self, index: int) -> int:
        """
        get the value of index-th node of the linked list
        :param index: index
        :return: value of the node
        """
        if index < 0 or index >= self.size:
            return -1

        p = self.head
        i = 0
        while i < index:
            p = p.next
            i += 1

        return p.val

    def add_at_head(self, val: int) -> None:
        """
        add a node of val in head of the linked list
        :param val: val of the node
        :return: None
        """
        node = Node(val)

        if self.head is None:
            self.head = node
            self.tail = node
        else:
            node.next = self.head
            self.head = node

        self.size += 1

    def add_at_tail(self, val: int) -> None:
        """
        add a node of val in the tail of the linked list
        :param val: val of the node
        :return: None
        """
        node = Node(val)

        if self.head is None:
            self.head = node
            self.tail = node
        else:
            self.tail.next = node
            self.tail = node

        self.size += 1

    def add_at_index(self, index: int, val: int) -> None:
        """
        add a node of val before the index-th node of the linked list
        :param index: index
        :param val: val of the node
        :return: None
        """
        if index < 0 or index > self.size:
            return

        node = Node(val)
        if 0 == index:  # add at head
            node.next = self.head
            self.head = node
            if 0 == self.size:
                self.tail = node
        elif self.size == index:  # add at tail
            self.tail.next = node
            self.tail = node
        else:
            p = self.head
            i = 0
            while i < index - 1:
                p = p.next
                i += 1
            node.next = p.next
            p.next = node

        self.size += 1

    def delete_at_index(self, index: int) -> None:
        """
        delete the index-th node of the linked list
        :param index: index
        :return: None
        """
        if index < 0 or index >= self.size:
            return

        if 0 == index:
            self.head = self.head.next
            if 1 == self.size:
                self.tail = None
        else:
            p = self.head
            i = 0
            while i < index - 1:
                p = p.next
                i += 1
            p.next = p.next.next
            if self.size - 1 == index:
                self.tail = p

        self.size -= 1

    def __str__(self) -> str:
        """
        return str of the linked list
        :return: str
        """
        p = self.head
        values = []

        while p is not None:
            values.append(str(p.val))
            p = p.next

        return ', '.join(values)


if '__main__' == __name__:
    my_list = MyLinkedList()

    my_list.add_at_head(1)
    print(my_list)

    my_list.add_at_head(2)
    print(my_list)

    my_list.add_at_head(3)
    print(my_list)

    print(my_list.get(0))
