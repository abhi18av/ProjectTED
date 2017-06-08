
#cd("/Users/eklavya/Projects/Polyglot/ProjectTED/experiments/dataScience/utils")

using JSON

filelist = []

for f in filter(x -> endswith(x, "json"), readdir("../../tedTalkDB"))
    fullFileName = "../../tedTalkDB/" * f
    push!(filelist, fullFileName)
end

#filelist[end]

_finalData = Dict()

_fileNameAndLanguageList = Dict()

_id = 1

function getDict(jsonPath)

split(jsonPath, "/")

for f in filelist

talkName = split(f, "/")[end]
#println(talkName)

_contents = JSON.Parser.parsefile(f)

lang_keys = collect(keys(_contents["TalkTranscriptPage"]["TalkTranscript"]))

#fileNameAndLanguageList[talkName] = lang_keys


_x =  @sprintf("%04d", id)
_finalData[_x] = Dict(talkName => lang_keys)


_id = _id + 1
end


_finalData

end
