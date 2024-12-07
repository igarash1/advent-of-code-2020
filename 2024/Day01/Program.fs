open System

let readUntilEOF () =
    let left = ResizeArray<int>()
    let right = ResizeArray<int>()
    
    let rec readLines () =
        match Console.ReadLine() with
        | null -> () // Stop reading when EOF is reached
        | line ->
            let parts = line.Split("   ")
            left.Add(int parts.[0])
            right.Add(int parts.[1])
            readLines () // Continue reading lines
    
    readLines ()
    List.sort (List.ofSeq left), List.sort (List.ofSeq right)

[<EntryPoint>]
let main _ =
    let left, right = readUntilEOF ()
    let rightFreq = 
        right   |> Seq.groupBy id
                |> Seq.map (fun (num, sq) -> num, Seq.length sq)
                |> dict
    let totalDiff = ([|0 .. left.Length - 1|]
    |> Array.map (fun i ->
                    let found, value = rightFreq.TryGetValue left.[i]
                    if found then
                        Math.Abs(left.[i] * value)
                    else
                        0)
    |> Array.reduce (fun a b -> a + b))
    printfn "%d" totalDiff
    0