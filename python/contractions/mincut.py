#!usr/bin/env python
#
# Karger's Min Cut algorithms 
#
import sys
import math
import random

class Vertex:
    def __init__(self):
        self.label = "" 
        self.edges = []

    def __str__(self):
        sout = "Label: {label}\n".format(label=self.label)
        for e in self.edges:
            sout += str(e)
        sout += "\n"
        return sout

class Edge:
    def __init__(self):
        self.v1 = None
        self.v2 = None

    def __str__(self):
        sout = " {v1} -> {v2}".format(v1=self.v1.label, v2=self.v2.label)
        return sout

class Graph:
    def __init__(self):
        self.V = []
        self.E = []
    
#    def contract(self, i):
#        for j in range(len(self.E[i].v2.edges)):
#            self.E[i].v2.edges[j].v1 = self.E[i].v1
#
#        for k in range(len(self.E[i].v2.edges)):
#            if self.E[i].v2.edges[k].v2.label != self.E[i].v1.label:
#                self.E[i].v1.edges.append(self.E[i].v2.edges[k])
#
#        self.E[i].v2.edges = []
#
#        for l in range(len(self.V)):
#              self.V[l].edges = [ e for e in self.V[l].edges if e.v2 != self.E[i].v2 ]
#
#        self.V = [ v for v in self.V if v != self.E[i].v2 ]
#        self.E = [ e for e in self.E if e.v1.label != e.v2.label and e.v2 != self.E[i].v2 ]

    def contract(self, i):
        pass

    def randomized_contraction(self):
        while len(self.V) > 2:
            i = random.randint(0,len(self.E)-1)
            self.contract(i)

        return len(self.E) 
    
    def __str__(self):
        sout = ""
        for v in self.V:
            sout += str(v)
        for e in self.E:
            sout += str(e)
        sout += "\n"
        return sout

def create_graph(table):
    g = Graph()
    #   init vertices 
    for row in table:
        v = Vertex()
        v.label = row[0] 
        g.V.append(v)

    for i in range(len(table)):
        for j in range(1,len(table[i])):
            e = Edge() 
            e.v1 = g.V[i]
            e.v2 = g.V[table[i][j]-1]
            g.V[i].edges.append(e)
            g.E.append(e)
    return g

if __name__ == "__main__":
    g = Graph()

