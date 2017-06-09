#Pkg.add("PyPlot")

cd("/Users/eklavya/Projects/Polyglot/ProjectTED/experiments/dataScience/Exp_2/julia")


using Gadfly
using JSON
using CSV
using DataFrames



file = "../finalData3.json"


_contents = JSON.Parser.parsefile(file)

ids = collect(keys(_contents))

#df = DataFrame()
i = 1
for id in ids
#println(id)
talk_name = _contents[id]["talkName"]
lang_keys = _contents[id]["lang_keys"]
langs_count = length(lang_keys)

if i == 1
  df = DataFrame( id_string = id, langs_count = langs_count, talk_name=talk_name)
  i = i + 1
end

append!(df, DataFrame( id_string = id, langs_count = langs_count, talk_name=talk_name))
#append!(df, DataFrame( id_string = id, langs_count = langs_count, talk_name=talk_name))
end

println(df)




id_string = collect(keys(_contents))[1]

id_int = parse(Int, id_string)

talk_name = collect(values(_contents))[1][1]

langs_string = collect(values(_contents))[1][2]

langs_count = length(langs_string)

df = DataFrame( id_string = id, langs_count = langs_count, talk_name = talk_name)

#typeof(df)

df.colindex





#CSV.write("out.csv", df; delim =';')


plot(df, x= :id_string, y= :langs_count)
