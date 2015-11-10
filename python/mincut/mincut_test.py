#!/usr/bin/env python
#
# Karger's Min Cut algorithm test 
#
import unittest
import mincut
import re
import sys

class TestMinCut(unittest.TestCase):

    def read_input(self, filename):
        f = open(filename, 'r')
        table = []
        line = ""
        for line in f:
            stripped = line.strip()
            if not stripped:
                continue 
            row = [ x for x in re.split('\s+', stripped) ]
            table.append(row)
	f.close()
        return table 

    def test_min_cut(self):
        result = 0
	# table = self.read_input('../../data/kargerMinCut.txt')
	table = self.read_input('../../data/reducedMinCut.txt')
	g = mincut.create_graph(table)
        for i in range(100):
		g = mincut.create_graph(table)
		cuts = g.randomized_contraction()
		sys.stdout.write("Cuts: {cuts}\n".format(cuts=cuts))
		if i == 0 or cuts < result:
			result = cuts

		sys.stdout.write("Min cut: {mincut}\n".format(mincut=result))


if __name__ == '__main__':
    unittest.main()

