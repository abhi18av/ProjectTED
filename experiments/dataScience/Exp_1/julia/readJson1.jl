# AIM : Read and print the contents of a JSON file

using JSON


#cd("Users/eklavya/Projects/Polyglot/TedTalks/ProjectTED/experiments/dataScience/Exp_1/julia/")
#run(`ls`)


#file = open("./al-Are_you_a_giver_or_a_taker_.json", "r")


contents = JSON.Parser.parsefile("./al-Are_you_a_giver_or_a_taker_.json")
json_contents = JSON.json(contents["transcript"],2)

write("json_contents.json", json_contents)

#println(contents)

contents["talk_link"]
length(contents)

keys(contents)

transcript = contents["transcript"]
keys(transcript)
korean_text = transcript["ko"]
#println(korean_text)
json_korean_text = JSON.json(korean_text)

write("korean_text.json", korean_text)
