langs = document.getElementsByClassName("h9")

for (l = 0; l < langs.length; l++) {

    lang_name = langs[l].textContent
    lang_code = langs[l].children[0].getAttribute("href")

    console.log(lang_name, " : ", lang_code)

}



/*

Afrikaans  :  /talks?language=af
Albanian  :  /talks?language=sq
Algerian Arabic  :  /talks?language=arq
Amharic  :  /talks?language=am
Arabic  :  /talks?language=ar
Armenian  :  /talks?language=hy
Assamese  :  /talks?language=as
Asturian  :  /talks?language=ast
Azerbaijani  :  /talks?language=az
Basque  :  /talks?language=eu
Belarusian  :  /talks?language=be
Bengali  :  /talks?language=bn
Bislama  :  /talks?language=bi
Bosnian  :  /talks?language=bs
Bulgarian  :  /talks?language=bg
Burmese  :  /talks?language=my
Catalan  :  /talks?language=ca
Cebuano  :  /talks?language=ceb
Chinese, Simplified  :  /talks?language=zh-cn
Chinese, Traditional  :  /talks?language=zh-tw
Chinese, Yue  :  /talks?language=zh
Creole, Haitian  :  /talks?language=ht
Croatian  :  /talks?language=hr
Czech  :  /talks?language=cs
Danish  :  /talks?language=da
Dutch  :  /talks?language=nl
Dzongkha  :  /talks?language=dz
English  :  /talks?language=en
Esperanto  :  /talks?language=eo
Estonian  :  /talks?language=et
Filipino  :  /talks?language=fil
Finnish  :  /talks?language=fi
French  :  /talks?language=fr
French (Canada)  :  /talks?language=fr-ca
Galician  :  /talks?language=gl
Georgian  :  /talks?language=ka
German  :  /talks?language=de
Greek  :  /talks?language=el
Gujarati  :  /talks?language=gu
Hakha Chin  :  /talks?language=cnh
Hausa  :  /talks?language=ha
Hebrew  :  /talks?language=he
Hindi  :  /talks?language=hi
Hungarian  :  /talks?language=hu
Hupa  :  /talks?language=hup
Icelandic  :  /talks?language=is
Igbo  :  /talks?language=ig
Indonesian  :  /talks?language=id
Ingush  :  /talks?language=inh
Irish  :  /talks?language=ga
Italian  :  /talks?language=it
Japanese  :  /talks?language=ja
Kannada  :  /talks?language=kn
Kazakh  :  /talks?language=kk
Khmer  :  /talks?language=km
Klingon  :  /talks?language=tlh
Korean  :  /talks?language=ko
Kurdish  :  /talks?language=ku
Kyrgyz  :  /talks?language=ky
Lao  :  /talks?language=lo
Latgalian  :  /talks?language=ltg
Latin  :  /talks?language=la
Latvian  :  /talks?language=lv
Lithuanian  :  /talks?language=lt
Luxembourgish  :  /talks?language=lb
Macedo  :  /talks?language=rup
Macedonian  :  /talks?language=mk
Malagasy  :  /talks?language=mg
Malay  :  /talks?language=ms
Malayalam  :  /talks?language=ml
Maltese  :  /talks?language=mt
Marathi  :  /talks?language=mr
Mauritian Creole  :  /talks?language=mfe
Mongolian  :  /talks?language=mn
Montenegrin  :  /talks?language=srp
Nepali  :  /talks?language=ne
Norwegian Bokmal  :  /talks?language=nb
Norwegian Nynorsk  :  /talks?language=nn
Occitan  :  /talks?language=oc
Pashto  :  /talks?language=ps
Persian  :  /talks?language=fa
Polish  :  /talks?language=pl
Portuguese  :  /talks?language=pt
Portuguese, Brazilian  :  /talks?language=pt-br
Punjabi  :  /talks?language=pa
Romanian  :  /talks?language=ro
Russian  :  /talks?language=ru
Rusyn  :  /talks?language=ry
Serbian  :  /talks?language=sr
Serbo-Croatian  :  /talks?language=sh
Silesian  :  /talks?language=szl
Sinhala  :  /talks?language=si
Slovak  :  /talks?language=sk
Slovenian  :  /talks?language=sl
Somali  :  /talks?language=so
Spanish  :  /talks?language=es
Swahili  :  /talks?language=sw
Swedish  :  /talks?language=sv
Swedish Chef  :  /talks?language=art-x-bork
Tagalog  :  /talks?language=tl
Tajik  :  /talks?language=tg
Tamil  :  /talks?language=ta
Tatar  :  /talks?language=tt
Telugu  :  /talks?language=te
Thai  :  /talks?language=th
Tibetan  :  /talks?language=bo
Turkish  :  /talks?language=tr
Turkmen  :  /talks?language=tk
Ukrainian  :  /talks?language=uk
Urdu  :  /talks?language=ur
Uyghur  :  /talks?language=ug
Uzbek  :  /talks?language=uz
Vietnamese  :  /talks?language=vi


*/
