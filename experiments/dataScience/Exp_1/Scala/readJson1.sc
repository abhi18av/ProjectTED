import scala.util.parsing.json._

val result = JSON.parseFull("""
  {"name": "Naoki",  "lang": ["Java", "Scala"]}
""")

result match {
  case Some(e) => println(e) // => Map(name -> Naoki, lang -> List(Java, Scala))
  case None => println("Failed.")
}
