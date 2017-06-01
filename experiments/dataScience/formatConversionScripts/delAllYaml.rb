
Dir.chdir(File.dirname(__FILE__))
jsonFiles = Dir["*.yaml"]

for f in jsonFiles

#newName1 =  f.split(".")[0] +  ".toml" 

#File.open(newName1, 'w')
cmd1 = "rm "  + f 
system(cmd1)

end
