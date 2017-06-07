# AIM : Read all json files and store the number of languages per talk in a json output.
using JSON

filelist = []

for f in filter(x -> endswith(x, "json"), readdir("../../sampleJSON"))
    fullFileName = "../../sampleJSON/" * f
    push!(filelist, fullFileName)
end

#filelist[end]


fileNameAndLanguage = Dict()

for f in filelist

talkName = f
#println(talkName)
contents = JSON.Parser.parsefile(f)

json_contents = JSON.json(contents["transcript"], 2)

write("out.json", json_contents)


end
