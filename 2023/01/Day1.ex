defmodule Day1 do

end


numbers = IO.read(:stdio, :all)
  |> String.split("\n")
  |> Enum.map(fn s -> String.replace(s, ~r/[^\d]/, "") end)
  |> Enum.map(fn s -> String.at(s, 0) <> String.at(s, String.length(s) - 1) end)
  |> Enum.map(fn s -> String.to_integer(s) end)
  |> IO.inspect()
  |> Enum.sum()


