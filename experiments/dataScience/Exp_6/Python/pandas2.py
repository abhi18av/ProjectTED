import pandas as pd
import json
import os
import glob


tedTalkDB = "/Users/eklavya/Projects/Polyglot/ProjectTED/experiments/tedTalkDB/"



ls1 = [1,2]
ls2 = [1,2,[3,4]]
ls3 = ['a', 'b','c','d',ls2]
df_ls = pd.DataFrame([ls1,ls2,ls3])


dc1 = {1:"a"}
dc2 = {1:"a", 2:"b", 3:{4:"c", 5:"d"}}
df_dc = pd.DataFrame([dc1,dc2])




f1 = tedTalkDB +  "shah_rukh_khan_thoughts_on_humanity_fame_and_love.json" 


df = pd.read_json(f1)


In [46]: with open("dataFrame.html", "w") as f:
       2     f.write(str(df.to_html))


