import Data.List

replaceChar :: Char -> Char
replaceChar char = case char of
                     '+' -> ' '
                     _   -> char

replace :: String -> String
replace = map replaceChar

toInt :: (String -> Integer)
toInt x = read x ::Integer

getInput :: IO [Integer]
getInput = do
    dat <- readFile "input.txt"
    let y = (map$toInt) $lines $replace dat
    return y

main :: IO ()
main = do
    f <- getInput
    print (sum f)
