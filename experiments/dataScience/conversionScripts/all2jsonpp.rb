jsonFiles = Dir["*.json"]

for f in jsonFiles

newName1 =  f.split(".")[0] + "__1__" + "." +  f.split(".")[1]
cmd1 = "json2json " + f + " > " + newName1
system(cmd1)


newName2 =  f.split(".")[0] + "__2__" + "." + f.split(".")[1]
cmd2 = "jsonpp " + newName1 + " > " + newName2
system(cmd2)


end
