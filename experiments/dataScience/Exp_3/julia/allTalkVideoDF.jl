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

videoPageDF = DataFrame( VideoURL = String[], #1
                         Speaker = String[],#2
                         AvailableSubtitlesCount = String[],#3
                         Duraton = String[],#4
                         #TalkTopicsList = String[], #5
                         TimeFilmed = String[], #6
                         TalkViewsCount = String[], #7
                         TalkCommentsCount = String[] ) #8





for file in filelist

_contents = JSON.Parser.parsefile(file)
_videoPage = _contents["TalkVideoPage"]

_keys = collect(keys(_videoPage))

#println(keys)






#df = DataFrame(a = String[], b= String[])
#push!(df, ["a","b"])




push!(videoPageDF,
[_videoPage[_keys[1]],
_videoPage[_keys[2]],
_videoPage[_keys[3]],
_videoPage[_keys[4]],
#_videoPage[keys[5]],
_videoPage[_keys[6]],
_videoPage[_keys[7]],
_videoPage[_keys[8]]])


end



#describe(_videoPageDF)
#size(_videoPageDF)

CSV.write("allTalkVideoPageDF.csv", videoPageDF; delim= ';')
