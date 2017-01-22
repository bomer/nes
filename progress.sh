a=`grep -o 'DONE' 6502.html  | wc -l`
b=56

echo "Done $a out of $b instructions"
echo "scale=2 ; $a / $b" | bc 

