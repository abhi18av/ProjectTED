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
shahrukh = contents["0013"]["transcript_en"][0]
doc1 = nlp(shahrukh)

doc1.ents


doc2 = nlp.make_doc(shahrukh)



"""

chunks = shahrukh.split("\n\n")



for i in range(len(chunks)):
     if chunks[i]  == "" or chunks[i] == "\n":
         continue
     print(i, chunks[i])
"""