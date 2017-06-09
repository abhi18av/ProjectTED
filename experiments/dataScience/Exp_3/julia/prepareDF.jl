cd("/Users/eklavya/Projects/Polyglot/ProjectTED/experiments/dataScience/Exp_2/julia")


using Gadfly
using JSON
using CSV
using DataFrames



file = "./finalData3.json"


_contents = JSON.Parser.parsefile(file)

ids = collect(keys(_contents))

df = DataFrame( id_string = String[], langs_count = Int[], talk_name=String[])
#df = DataFrame()
i = 1
for id in ids
#println(id)
talk_name = _contents[id]["talkName"]
lang_keys = _contents[id]["lang_keys"]
langs_count = length(lang_keys)


push!(df, [id, langs_count, talk_name])
#append!(df, DataFrame( id_string = id, langs_count = langs_count, talk_name=talk_name))
end

#println(df)
#typeof(df)

CSV.write("finalData3.csv", df; delim=';')


