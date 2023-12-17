import scala.io._

object a {
    def main(args: Array[String]) = {
        println(Source.fromFile("input.txt").getLines().map(x => (x.toInt) / 3 - 2).sum)
    }
}

