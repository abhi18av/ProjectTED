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


inSentences = [%Q("My name is alan."), %Q("I am 20 years old.")]

for s in inSentences
	cmd1 = "trans -brief en:ja " + s +  " >> out.txt"
	system(cmd1)
end



	#outSentences.push(x)
	#puts cmd1
