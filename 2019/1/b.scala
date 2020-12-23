import scala.io._

object b {
    def main(args: Array[String]) = {
        println(Source.fromFile("input.txt").getLines().map(x => fuel(x.toInt)).sum)
    }

    def fuel(x: Int): Int = (x / 3) - 2 match {
        case x if (x <= 0) => 0
        case x => x + fuel(x)
    }
}

