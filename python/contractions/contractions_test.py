#!/usr/bin/env python
#
# Counting inversions
#
import unittest
import contractions
import re
import sys

class TestContraction(unittest.TestCase):

    def read_input(self, filename):
        f = open(filename, 'r')
        table = []
        line = ""
        for line in f:
            stripped = line.strip()
            if not stripped:
                continue 
            row = [ int(x) for x in re.split('\s+', stripped) ]
            table.append(row)
        return table 

    # ('Inversions: ', 2407905288)
    def test_min_cut(self):
        table = self.read_input('../../data/reducedMinCut.txt')
        g = contractions.create_graph(table)

        sys.stdout.write("Edges: {edges}\n".format(edges=len(g.E)))
#        self.assertEqual(2407905288, count) 



if __name__ == '__main__':
    unittest.main()
