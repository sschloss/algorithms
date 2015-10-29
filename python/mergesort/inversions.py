#!/usr/bin/env python

import sys
import math

def merge_sort(a, data): 
    if len(a) < 2:
        return a
    else:
        mid = int(math.floor(len(a) / 2.0))
        merged = []
        l = merge_sort(a[:mid], data)
        r = merge_sort(a[mid:], data)
        return merge(l, r, data)

def merge(left, right, data): 
    result = []

    i = j = 0
    while i < len(left) and j < len(right):
        if left[i] < right[j]:
            result.append(left[i])
            i = i + 1
        else:
            result.append(right[j])
            j = j + 1
            # count the inversions
            data['inv'] = data['inv'] + len(left[i:])

    while i < len(left):
        result.append(left[i])
        i = i + 1

    while j < len(right):
        result.append(right[j])
        j = j + 1

    return result 

def num(s):
    try:
        return int(s)
    except ValueError:
        return float(s)



