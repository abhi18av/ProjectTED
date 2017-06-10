# AIM : Create a dataFrame from the information held by talkVideoPage in the json

cd("/Users/eklavya/Projects/Polyglot/ProjectTED/experiments/dataScience/Exp_3/julia")


using JSON
using CSV
using DataFrames

file = "./laura_robinson_the_secrets_i_find_on_the_mysterious_ocean_floor.json"


_contents = JSON.Parser.parsefile(file)
transcriptPage = _contents["TalkTranscriptPage"]

_keys = collect(keys(transcriptPage))

#println(keys)



 _transcriptPageDF = DataFrame( #AvailableTranscripts = String[], #1
                          #Rated = String[],#2
                          #TalkTranscript = String[],#3
                          DatePosted = String[] ) #4



#df = DataFrame(a = String[], b= String[])
#push!(df, ["a","b"])




push!(_transcriptPageDF,
[
#transcriptPage[_keys[1]],
#transcriptPage[_keys[2]],
#transcriptPage[_keys[3]],
transcriptPage[_keys[4]]])



#describe(videoPageDF)
#size(videoPageDF)

CSV.write("DF.csv", _transcriptPageDF; delim= ';')
