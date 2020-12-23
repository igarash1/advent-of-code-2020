import io.AnsiColor._

object Day2 {
    def main(args: Array[String]): Unit = {
        TestPart1()
        TestPart2()
    }

    def TestPart1(): Unit = {
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
            }
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

    def TestPart2(): Unit = {
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

    def part2(input: String): Int = {
        val nums = input.split(",").map(Integer.parseInt)
        var noun = 0
        while (noun < 100) {
            var verb = 0
            while (verb < 100) {
                val newNums = new Array[Int](nums.length)
                for (i <- 0 until nums.length) {
                    newNums(i) = nums(i)
                }
                newNums(1) = noun
                newNums(2) = verb
                if (compute(newNums) == 19690720) {
                    return 100 * noun + verb
                }
                verb += 1
            }
            noun += 1
        }
        return -1
    }

    def compute(nums: Array[Int]): Int = {
        var i = 0
        while (i < nums.length) {
            val opcode = nums(i)
            if (opcode == 99) return nums(0)
            val a = nums(nums(i + 1))
            val b = nums(nums(i + 2))
            val out = nums(i + 3)
            if (1 <= opcode && opcode <= 2) {
                nums(out) = opcode match {
                    case 1 => a + b
                    case 2 => a * b
                }
            } else {
                return -1
            }
            i += 4
        }
        return nums(0)
    }
}

