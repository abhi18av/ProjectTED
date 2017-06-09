# AIM : Create a dataFrame from the information held by talkVideoPage in the json

cd("/Users/eklavya/Projects/Polyglot/ProjectTED/experiments/dataScience/Exp_3/julia")


using JSON
using DataFrames

file = "./laura_robinson_the_secrets_i_find_on_the_mysterious_ocean_floor.json"


_contents = JSON.Parser.parsefile(file)
videoPage = _contents["TalkVideoPage"]

keys = collect(keys(videoPage))

println(keys)



 videoPageDF = DataFrame( VideoURL = String[], #1
                          Speaker = String[],#2
                          AvailableSubtitlesCount = String[],#3
                          Duraton = String[],#4
                          #TalkTopicsList = String[], #5
                          TimeFilmed = String[], #5
                          TalkViewsCount = String[], #6
                          TalkCommentsCount = String[] ) #7


push!(videoPageDF,
      videoPage[string(keys[1])],
      videoPage[string(keys[2])],
      videoPage[string(keys[3])],
      videoPage[string(keys[4])],
      #videoPage[string(keys[5])],
      videoPage[string(keys[6])],
      videoPage[string(keys[7])])



"""

#      push!(videoPageDF,
#            videoPage[keys[1]],
#            videoPage[keys[2]],
#            videoPage[keys[3]],
#            videoPage[keys[4]],
#            #videoPage[keys[5]],
#            videoPage[keys[6]],
#            videoPage[keys[7]])



println(videoPage)
"""
