require 'json'

file = File.read("./talkIndex.json")

halfLinks = JSON.parse(file)

for link in halfLinks
# for i in 1..10
	url = "https://www.ted.com" + link
	#puts url
system('./goTALK', url)
print(url) 
end
