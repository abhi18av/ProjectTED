import json
import pandas as pd
import os
import glob


tedTalkDB = "/Users/eklavya/Projects/Polyglot/ProjectTED/experiments/tedTalkDB/"

files = glob.glob(tedTalkDB + "*.json")

filelist = []

"""
for f in files:
    content = json.load(open(f))
    title_en = content["TalkTranscriptPage"]["TalkTranscript"]["en"]["LocalTalkTitle"]
    print(title_en)
"""


f1 = tedTalkDB +  "shah_rukh_khan_thoughts_on_humanity_fame_and_love.json" 
f2 = tedTalkDB +  "sarah_kay_how_many_lives_can_you_live.json" 

df = pd.read_json(f1)

df.append(pd.read_json(f2))