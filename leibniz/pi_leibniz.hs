import Data.Function

qtrPi :: Double -> Double
qtrPi 0 = 1
--qtrPi n = (qtrPi (n-1)) + (-1)**n / (2*n + 1)
qtrPi n = last/4 + (-1)**n / (2*n + 1) -- memoization
          where last = snd (leibniz !! (floor (n-1)))

leibniz :: [(Double, Double)]
leibniz = zip [0..] (map ((* 4) . qtrPi) [0..])

main :: IO ()
main = do
    mapM_ (\(a,b) -> putStrLn $ show (floor a) ++ ":\t" ++ show b) leibniz
