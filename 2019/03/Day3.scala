package Day3

import io.AnsiColor._

object Day3 {
    def main(args: Array[String]): Unit = {
        testPart1()
        testPart2()
    }

    def testPart1(): Unit = {
        val tests = Seq[(String, Int)](
            ("1,0,0,0,99"
                , 2),
            ("R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83"
                , 159),
            ("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
                , 135))
        tests.zipWithIndex.foreach(tc => {
            val actual = part1(tc._1._1)
            printf("#%d ", tc._2)
            if (tc._1._2 != actual) {
                printf(s"${RED}FAILED${RESET} expected: %d, actual: %d\n", tc._1._2, actual)
            } else {
                printf(s"${GREEN}OK${RESET}\n")
            }
        })
    }

    def part1(input: String): Int = {
    }

    def testPart2(): Unit = {

    }

    def part2(input: String): Int = {

    }
}

