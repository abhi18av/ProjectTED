#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Created on Sun Jun 11 21:52:00 2017

@author: eklavya
"""


import json
import spacy


nlp = spacy.load("en")
contents = json.load(open("./finalData4.json"))
shahrukh = contents["0087"]["transcript_en"][0]
paragraphs = shahrukh.split("\n\n")

#remove all new lines from the paragraphs
arr = []
for p in paragraphs:
	if p == "":
		None
	else:
		arr.append(p)
		paragraphs = arr
del(arr)


# find the longest paragraph
#arr = []
#for p in paragraphs:
#	len(p)





doc1 = nlp(paragraphs[51])

for token in doc1:
	print(token)


for entity in doc1.ents:
	print(entity)


doc2 = nlp.make_doc(shahrukh)



"""

chunks = shahrukh.split("\n\n")



for i in range(len(chunks)):
     if chunks[i]  == "" or chunks[i] == "\n":
         continue
     print(i, chunks[i])
"""