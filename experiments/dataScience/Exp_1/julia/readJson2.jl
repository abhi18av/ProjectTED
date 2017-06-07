# AIM : Read all files in a directory

arr = run(`ls`)

filelist = []

for f in filter(x -> endswith(x, "json"), readdir("../../sampleJSON"))
    push!(filelist, f)
end

#filelist[end]
