# AIM : Read all json files and store the number of languages per talk in a json output.
using JSON

filelist = []

for f in filter(x -> endswith(x, "json"), readdir("../../sampleJSON"))
    fullFileName = "../../sampleJSON/" * f
    push!(filelist, fullFileName)
end

#filelist[end]


fileNameAndLanguageList = Dict()

for f in filelist

talkName = split(f, "/")[end]
#println(talkName)

_contents = JSON.Parser.parsefile(f)

lang_keys = keys(_contents["transcript"])

fileNameAndLanguageList[talkName] = lang_keys

end

#println(fileNameAndLanguageList)


json_output = JSON.json(fileNameAndLanguageList, 2)
write("out.json", json_output)
