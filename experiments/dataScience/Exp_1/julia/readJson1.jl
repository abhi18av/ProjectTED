using JSON


#cd("Users/eklavya/Projects/Polyglot/TedTalks/ProjectTED/experiments/dataScience/Exp_1/julia/")
#run(`ls`)


#file = open("./al-Are_you_a_giver_or_a_taker_.json", "r")


contents = JSON.Parser.parsefile("./al-Are_you_a_giver_or_a_taker_.json")

#println(contents)

contents["talk_link"]
length(contents)

keys(contents)

transcript = contents["transcript"]
keys(transcript)
println(transcript["ko"])
