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

filelist = []

for f in filter(x -> endswith(x, "json"), readdir("../../../tedTalkDB"))
    fullFileName = "../../../tedTalkDB/" * f
    push!(filelist, fullFileName)
end

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
write("finalData.json", json_output)
