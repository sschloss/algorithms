#!/usr/bin/env python
#
# Counting inversions
#
import unittest
import inversions

class TestInversions(unittest.TestCase):

    # ('Inversions: ', 2407905288)
    def test_inversions(self):
        a = self.read_input('../data/IntegerArray.txt')
        count  = inversions.count_inversions(a)
        self.assertEqual(2407905288, count) 

    def read_input(self, filename):
        f = open(filename, 'r')
        a = []
        line = ""
        for line in f:
            stripped = line.strip()
            if not stripped:
                continue 
            a.append(int(stripped))
        return a


if __name__ == '__main__':
    unittest.main()
