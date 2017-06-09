#Pkg.add("PyPlot")

cd("/Users/eklavya/Projects/Polyglot/ProjectTED/experiments/dataScience/Exp_2/julia")


using Gadfly
using JSON
using CSV
using DataFrames



file = "../data_laura_robinson_the_secrets_i_find_on_the_mysterious_ocean_floor.json"




_contents = JSON.Parser.parsefile(file)

id_string = collect(keys(_contents))[1]

id_int = parse(Int, id_string)

talk_name = collect(values(_contents))[1][1]

langs_string = collect(values(_contents))[1][2]

langs_count = length(langs_string)

df = DataFrame( id_string = id_string, langs_count = langs_count, talk_name=talk_name)

typeof(df)

df.colindex





#CSV.write("out.csv", df; delim =';')


plot(df, x= :id_string, y= :langs_count)
