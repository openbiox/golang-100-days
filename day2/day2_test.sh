for i in 1 2 3 4
do
  go build ./day2_00${i}.go
done


echo "a\na\nb\nc\na" | ./day2_001

echo "a\na\nb\nc" | ./day2_002

echo "day2_002 for file param:"
./day2_002 ./data/text.txt

echo "day2_004 for file param:"
./day2_004 ./data/text2.txt ./data/text3.txt

for i in 1 2 3 4
do
  rm ./day2_00${i}
done
