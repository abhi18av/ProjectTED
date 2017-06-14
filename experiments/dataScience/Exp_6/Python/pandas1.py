import json
import pandas
import os
import glob


tedTalkDB = "/Users/eklavya/Projects/Polyglot/ProjectTED/experiments/tedTalkDB/"

files = glob.glob(tedTalkDB + "*.json")

filelist = []


for f in files:
    content = json.load(open(f))
    title_en = content["TalkTranscriptPage"]["TalkTranscript"]["en"]["LocalTalkTitle"]
    print(title_en)
