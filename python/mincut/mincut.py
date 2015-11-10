#!usr/bin/env python
#
# Karger's Min Cut algorithms 
#
import sys
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

    def __init__(self, v1, v2):
        self.v1 = v1
        self.v2 = v2

    def __str__(self):
        sout = " {v1} -> {v2}".format(v1=self.v1, v2=self.v2)
        return sout

class Graph:
    def __init__(self):
        self.V = {} 
        self.E = []

    def contract(self, i):
        edges = []
        se = self.E[i]
        v1 = se.v1
        v2 = se.v2

        # move edges from v2 to v1  
        for j in range(len(self.V[v2].edges)):
            if v1 != self.V[v2].edges[j].v2:
                e = Edge(v1, self.V[v2].edges[j].v2)
                self.V[v1].edges.append(e)

        # update vertices to point to merged vertex
        for k, v in self.V.iteritems():
            edges = []
            for j in range(len(self.V[k].edges)):
                e = None
                if self.V[k].edges[j].v2 == v2:
                    e = Edge(self.V[k].edges[j].v1, v1)
                else:
                    e = Edge(self.V[k].edges[j].v1, self.V[k].edges[j].v2)
                edges.append(e)
            self.V[k].edges = edges

        # remove self loops in vertex edges
        self.V[v1].edges = [ e for e in self.V[v1].edges if e.v1 != e.v2 ] 
       
        # update the edges
        for j in range(len(self.E)):
            if self.E[j].v1 == v2:
                self.E[j].v1 = v1 
            elif self.E[j].v2 == v2:
                self.E[j].v2 = v1
        
        # remove self loops in edges
        self.E = [ e for e in self.E if e.v1 != e.v2 ]

        # remove v2 vertex         
        del self.V[v2]


    def randomized_contraction(self):
        while len(self.V) > 2:
            i = random.randint(0,len(self.E)-1)
            self.contract(i)
        return len(self.E) / 2 
    
    def __str__(self):
        sout = ""
        for k,v in self.V.iteritems():
            sout += str(v)
        sout += "Edges:\n"
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
        g.V[row[0]] = v

    for i in range(len(table)):
        for j in range(1,len(table[i])):
            v1_label = g.V[table[i][0]].label
            v2_label = g.V[table[i][j]].label
            e = Edge(v1_label, v2_label)
            g.V[table[i][0]].edges.append(e)
            g.E.append(e)
    return g

if __name__ == "__main__":
    g = Graph()

