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
            ("2,3,0,3,99"
                , 2),
            ("2,4,4,5,99,0"
                , 2),
            ("1,1,1,4,99,5,6,0,99"
                , 30),
            ("1,9,10,3,2,3,11,0,99,30,40,50"
                , 3500),
            ("1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,19,6,23,1,23,6,27,1,13,27,31,2,13,31,35,1,5,35,39,2,39,13,43,1,10,43,47,2,13,47,51,1,6,51,55,2,55,13,59,1,59,10,63,1,63,10,67,2,10,67,71,1,6,71,75,1,10,75,79,1,79,9,83,2,83,6,87,2,87,9,91,1,5,91,95,1,6,95,99,1,99,9,103,2,10,103,107,1,107,6,111,2,9,111,115,1,5,115,119,1,10,119,123,1,2,123,127,1,127,6,0,99,2,14,0,0"
                , 12490719))
        tests.zipWithIndex.foreach(tc => {
            val actual = part1(tc._1._1)
            printf("#%d ", tc._2)
            if (tc._1._2 != actual) {
                printf(s"${RED}FAILED${RESET} expected: %d, actual: %d\n", tc._1._2, actual)
            } else {
                printf(s"${GREEN}OK${RESET}\n")
            }3
        })
    }

    def part1(input: String): Int = {
        val nums = input.split(",").map(Integer.parseInt)
        if (input.length > 100) {
            nums(1) = 12
            nums(2) = 2
        }
        return compute(nums)
    }

    def testPart2(): Unit = {
        val tests = Seq[(String, Int)](
            ("1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,19,6,23,1,23,6,27,1,13,27,31,2,13,31,35,1,5,35,39,2,39,13,43,1,10,43,47,2,13,47,51,1,6,51,55,2,55,13,59,1,59,10,63,1,63,10,67,2,10,67,71,1,6,71,75,1,10,75,79,1,79,9,83,2,83,6,87,2,87,9,91,1,5,91,95,1,6,95,99,1,99,9,103,2,10,103,107,1,107,6,111,2,9,111,115,1,5,115,119,1,10,119,123,1,2,123,127,1,127,6,0,99,2,14,0,0"
                , 2003))
        tests.zipWithIndex.foreach(tc => {
            val actual = part2(tc._1._1)
            printf("#%d ", tc._2)
            if (tc._1._2 != actual) {
                printf(s"${RED}FAILED${RESET} expected: %d, actual: %d\n", tc._1._2, actual)
            } else {
                printf(s"${GREEN}OK${RESET}\n")
            }
        })
    }

}

