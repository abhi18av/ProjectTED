inSentences = ["Hello!", "My name is Alan.", "How old are you?"]
outSentences = []

for s in inSentences
	cmd1 = "trans" + " -brief" + " en:es " + s +  +" >> outSentences.txt"
	system(cmd1)
	#outSentences.push(x)
	#puts cmd1
end

for i in outSentences

	puts i

end
