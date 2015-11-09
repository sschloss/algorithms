#!/usr/bin/env python
#
# Counting inversions
#
import unittest
import mincut
import re
import sys
import random

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

    def test_min_cut(self):
        result = 0

        table = self.read_input('../../data/reducedMinCut.txt')
        g = mincut.create_graph(table)

        i = 0
        print "Edge: " + str(i) + "\n"
        print g 
        g.contract(i)
        print g 
#       while len(g.E) > 2:
#            i = random.randint(0,len(g.E)-1)
        
#        for i in range(100):
#            table = self.read_input('../../data/reducedMinCut.txt')
##            table = self.read_input('../../data/kargerMinCut.txt')
#            g = mincut.create_graph(table)
#            g.randomized_contraction()
#            cuts = len(g.E)
#            sys.stdout.write("Cuts: {cuts}\n".format(cuts=cuts))
#            if i == 0 or cuts < result:
#                result = cuts
#        
#        sys.stdout.write("Min cut: {mincut}\n".format(mincut=result))
#        self.assertEqual(2407905288, count) 

    

if __name__ == '__main__':
    unittest.main()
