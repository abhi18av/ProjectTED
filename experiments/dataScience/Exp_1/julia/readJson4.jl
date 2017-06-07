# AIM : To have a output json in a more structured form like
# {
#id:
#{
#  talkTitle:
#    {
#    langCodes,

#    }
#  }
#}


using JSON

filelist = []

for f in filter(x -> endswith(x, "json"), readdir("../../sampleJSON"))
    fullFileName = "../../sampleJSON/" * f
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

lang_keys = collect(keys(_contents["transcript"]))

fileNameAndLanguageList[talkName] = lang_keys


_x =  @sprintf("%04d", 10)
finalData[_x] = fileNameAndLanguageList


id = id + 1
end

#println(fileNameAndLanguageList)


json_output = JSON.json(finalData, 2)
write("finalData.json", json_output)
