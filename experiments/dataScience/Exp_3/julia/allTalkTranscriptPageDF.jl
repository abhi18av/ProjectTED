# AIM : Create a dataFrame from the information held by talk_videoPage in the json

cd("/Users/eklavya/Projects/Polyglot/ProjectTED/experiments/dataScience/Exp_3/julia")


using JSON
using CSV
using DataFrames

#file = "./laura_robinson_the_secrets_i_find_on_the_mysterious_ocean_floor.json"

filelist = []

for f in filter(x -> endswith(x, "json"), readdir("../../../tedTalkDB"))
    fullFileName = "../../../tedTalkDB/" * f
    push!(filelist, fullFileName)
end


# filelist


transcriptPageDF = DataFrame( AvailableTranscripts = DataArray[], #1
Rated = DataArray[],#2
#TalkTranscript = DataArray[],#3
DatePosted = String[] ) #4




for file in filelist


_contents = JSON.Parser.parsefile(file)
transcriptPage = _contents["TalkTranscriptPage"]

_keys = collect(keys(transcriptPage))

#println(keys)






#df = DataFrame(a = String[], b= String[])
#push!(df, ["a","b"])





push!(transcriptPageDF,
[
convert(DataArray, transcriptPage[_keys[1]]),
convert(DataArray, Array([transcriptPage[_keys[2]]])),
#convert(DataArray, transcriptPage[_keys[3]]),
transcriptPage[_keys[4]]
	 ])

end



#describe(_videoPageDF)
#size(_videoPageDF)

CSV.write("allTalkTranscriptPageDF.csv", transcriptPageDF; delim= ';')
