# AIM : To have a output sorted json in a more structured form like
# {
#id:
#{
#  talkTitle:
#    {
#    langCodes,

#    }
#  }
#}

cd("/Users/eklavya/Projects/Polyglot/ProjectTED/experiments/dataScience/Exp_1/julia")


using JSON

filelist = ["../laura_robinson_the_secrets_i_find_on_the_mysterious_ocean_floor.json"]
#filelist[end]


finalData = Dict()

fileNameAndLanguageList = Dict()

id = 1

for f in filelist

talkName = split(f, "/")[end]
#println(talkName)

_contents = JSON.Parser.parsefile(f)

lang_keys = collect(keys(_contents["TalkTranscriptPage"]["TalkTranscript"]))

#fileNameAndLanguageList[talkName] = lang_keys


_x =  @sprintf("%04d", id)
finalData[_x] = Dict(talkName => lang_keys)


id = id + 1
end

#println(fileNameAndLanguageList)

#finalData = sort(collect(finalData), by=x->parse(Int, x[1]))
#sort(collect(finalData))
json_output = JSON.json(finalData, 2)
write("data_laura_robinson_the_secrets_i_find_on_the_mysterious_ocean_floor.json", json_output)
