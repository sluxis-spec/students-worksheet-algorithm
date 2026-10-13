**GitHub username:** sluxis-spec
**Exercises completed:** Exercise 1-9
**Most surprising bug:** At first, I thought my program was stuck because nothing appeared in the terminal. Later, I realized that `fmt.Scan()` was waiting for my keyboard input. I learned that a program can be waiting for input without showing an error.

For each exercise you completed, copy and fill in this mini-report:

## Exercise 01 — Data_types
- **Story in my words:** Maya needs to store student information for a digital check-in desk. Each piece of information uses a data type that matches its purpose.
- **Input(s) and their types:** 
- age: int, with the value 19.
- mass: float64, with the value 62.5.
- checkedIn: bool, with the value true.
- studentID: string, with the value "00123".
- **Processing / logical rule:** The program assigns values to four variables and prints them on one line, separated by the | symbol.
- **Expected output (before coding):** 19 | 62.5 | true | 00123
- **Actual output:** 19 | 62.5 | true | 00123
- **Normal test:** Use the given values and check that all four values are printed correctly.
- **Boundary or invalid test:** Use a student ID such as "00001" and check whether the leading zeros are preserved.
- **One assumption or limitation:** The program uses fixed values instead of accepting information entered by a user.
- **What I changed in `main.go`:** I filled in the values for age, mass, checkedIn, and studentID using the data types specified in the task.
- **Why I believe the result is correct:** Because the program prints the expected values, and the student ID keeps its leading zeros because it is stored as a string.

## Exercise 02 — Cafe_sequence
- **Story in my words:** A student buys two coffees that cost Rp18,000 each and receives a Rp5,000 discount. The program calculates the subtotal before subtracting the discount.
- **Input(s) and their types:** 
- unitPrice: int, with the value 18000.
- quantity: int, with the value 2.
- discount: int, with the value 5000.
- **Processing / logical rule:** First, the program multiplies unitPrice by quantity to calculate the subtotal. Then, it subtracts discount from the subtotal to calculate the total.
- **Expected output (before coding):** 
- Subtotal: 36000
- Total: 31000
- **Actual output:** 
- Subtotal: 36000
- Total: 31000
- **Normal test:** Use a unit price of 18000, a quantity of 2, and a discount of 5000. The expected subtotal is 36000, and the expected total is 31000.
- **Boundary or invalid test:** Is using a discount greater than the subtotal, such as 40000. The calculation produces a negative total, so the program should validate the discount if negative totals are not allowed.
- **One assumption or limitation:** The program assumes the price, quantity, and discount are valid. It does not check whether the discount exceeds the subtotal.
- **What I changed in `main.go`:** I set the unit price, quantity, discount, and calculated the subtotal before the total, and printed both values.
- **Why I believe the result is correct:** Because the subtotal is calculated as 18000 * 2 = 36000. Subtracting the discount gives 36000 - 5000 = 31000.

## Exercise 03 — Factor_checker
- **Story in my words:** Maya wants to divide y volunteers into equal-sized groups of x people without leaving anyone out.
- **Input(s) and their types:** 
- x: int, the number of volunteers in each group.
- y: int, the total number of volunteers.
- **Processing / logical rule:** The program first checks whether either number is less than or equal to zero. If so, it prints invalid input. Otherwise, it checks whether y % x == 0. If the remainder is zero, x divides y evenly, so the program prints true; otherwise, it prints false.
- **Expected output (before coding):** 
- 5 20 : true
- 20 5 : false
- 123 123 : true
- 0 20 : invalid input
- **Actual output:** 
- 5 20 : true
- 20 5 : false
- 123 123 : true
- 0 20 : invalid input
- **Normal test:** If we enter 5 20. The program prints true because 20 % 5 == 0.
- **Boundary or invalid test:** If we enter 0 20. The program prints invalid input without performing the modulo operation.
- **One assumption or limitation:** The program checks whether both numbers are positive, but assumes the user enters valid integers.
- **What I changed in `main.go`:** I used fmt.Scan() to read two integers, added input validation before the modulo operation, and printed the result of y % x == 0.
- **Why I believe the result is correct:** The modulo operation determines whether there is a remainder after dividing y by x. A remainder of zero means the volunteers can be divided into equal-sized groups without anyone left over.

## Exercise 04 — Population
- **Story in my words:** A city wants to estimate its final population by adding births and incoming residents, then subtracting deaths and people who leave.
- **Input(s) and their types:** 
- start: int, with the value 1000.
- births: int, with the value 20.
- arrivals: int, with the value 5.
- deaths: int, with the value 10.
- departures: int, with the value 4.
- **Processing / logical rule:** The program calculates the final population using start + births + arrivals - deaths - departures.
- **Expected output (before coding):** Population: 1011
- **Actual output:** Population: 1011
- **Normal test:** Use the given values: initial population 1000, births 20, arrivals 5, deaths 10, and departures 4, and the expected final population is 1011.
- **Boundary or invalid test:** Try a values that would produce a negative final population. The result indicates that the input data may be inconsistent with a real population model or the model needs additional validation.
- **One assumption or limitation:** The model assumes that the initial population and all population changes are recorded accurately. It does not validate the input values.
- **What I changed in `main.go`:** I assigned the five population values, calculated the final population using the given formula, and printed the result.
- **Why I believe the result is correct:** The calculation is 1000 + 20 + 5 - 10 - 4 = 1011, which matches the expected output.

## Exercise 05 — A Composer's Pattern: Number Sequences
- **Story in my words:** A composer starts with two numbers. Each new number is calculated by adding the previous two numbers together.
- **Input(s) and their types:** 
- first = 1 (int)
- second = 1 (int)
  ----------------
- first = 5 (int)
- second = 3 (int)
  ----------------
- first = 10 (int)
- second = 20 (int)
- **Processing / logical rule:** The program calculates the third number by adding first and second. Then, it calculates the fourth number using second and third, and the fifth number using third and fourth.
- **Expected output (before coding):** 
- 2 3 5
  ----------------
- 8 11 19
  ----------------
- 30 50 80
- **Actual output:** 
- 2 3 5
  ----------------
- 8 11 19
  ----------------
- 30 50 80
- **Normal test:** 
Normal test: starting values 5 and 3 produce 30 50 80.
Additional test: starting values 1 and 1 should produce 2 3 5.
Additional test: starting values 10 and 20 should produce 8 11 19.
- **Boundary or invalid test:** Test the starting values 0 and 0. The expected output is 0 0 0. Zero is a valid starting value, and the program should still calculate the sequence correctly.
- **One assumption or limitation:** The program uses fixed starting values and does not ask the user for input. It calculates the sequence directly without using a loop.
- **What I changed in `main.go`:** I assigned the two starting values and calculated the next three numbers using addition.
- **Why I believe the result is correct:** Each new number is the sum of the previous two numbers, as required by the task.

## Exercise 06 — Gravity on Mars
- **Story in my words:** An astronaut's mass stays the same on Earth and Mars, but their weight changes because the gravitational force is different.
- **Input(s) and their types:** 
- massKg = 70.0 (float64)
- earthGravity = 9.8 (float64)
- marsFactor = 0.38 (float64)
- **Processing / logical rule:** The program calculates Earth weight by multiplying mass by Earth's gravity. It then calculates Mars weight by multiplying Earth weight by the Mars factor.
- **Expected output (before coding):** 
- Earth: 686.0 N
- Mars: 260.7 N
- **Actual output:** 
- Earth: 686.0 N
- Mars: 260.7 N
- **Normal test:** A mass of 70.0 kg should produce an Earth weight of 686.0 N and a Mars weight of 260.7 N.
- **Boundary or invalid test:** Test a mass of 0.0 kg. The expected output is Earth: 0.0 N and Mars: 0.0 N. The program accepts zero, but it does not validate negative mass values.
- **One assumption or limitation:** The program uses a fixed mass and fixed gravity values. It assumes that the input mass is valid.
- **What I changed in `main.go`:** I assigned the mass, Earth's gravity, and Mars factor. I calculated Earth weight first, then used it to calculate Mars weight. I also formatted the outputs to one decimal place.
- **Why I believe the result is correct:** The calculations follow the formulas provided in the task, and %.1f formats each result to one decimal place.

## Exercise 07 — Optional AI Preview: Tiny Genetic Algorithm Operations
- **Story in my words:** A four-switch controller can be any four-bit string (such as 1001). The objective is to reach 1111.
- **Input(s) and their types:** 
- a = "1001" (string)
- b = "0110" (string)
- fitness returns an int representing the number of 1 bits.
- **Processing / logical rule:** The program calculates the fitness of both parents. It then performs crossover by taking the first two bits of Parent A and the last two bits of Parent B. Finally, it changes the second bit of the child to 1.
- **Expected output (before coding):** 
- Parent A: 1001 fitness = 2
- Parent B: 0110 fitness = 2
- Child: 1010 fitness = 2
- Mutated: 1110 fitness = 3
- **Actual output:** 
- Parent A: 1001 fitness = 2
- Parent B: 0110 fitness = 2
- Child: 1010 fitness = 2
- Mutated: 1110 fitness = 3
- **Normal test:** The parent strings 1001 and 0110 should each have a fitness of 2. The child 1010 should also have a fitness of 2.
- **Boundary or invalid test:** Test the candidate "0000". Its fitness should be 0 because it contains no 1 bits. The fitness function assumes the candidate contains only binary characters (0 and 1); it does not validate other characters.
- **One assumption or limitation:** This is a demonstration of genetic algorithm operations, not a full genetic algorithm. It does not evolve a population through multiple generations.
- **What I changed in `main.go`:** I corrected the fitness function syntax and updated the print statements to show each candidate and its fitness.
- **Why I believe the result is correct:** The program counts the 1 bits, combines the parent strings as instructed, and changes the second bit of the child to 1.

## Exercise 08 — Cube-root Check: A Precise Rule
- **Story in my words:** A game checks whether an integer x is the cube root of another integer y. The door unlocks only when multiplying x by itself three times gives y.
- **Input(s) and their types:** 
- x (int): the submitted cube-root candidate.
- y (int): the number being checked.
- **Processing / logical rule:** The program calculates x*x*x and compares the result with y. It prints true if they are equal and false otherwise.
- **Expected output (before coding):** 
- Input 5 125 → true
- Input -3 -27 → true
- Input 12 12 → false
- **Actual output:** 
- Input 5 125 → true
- Input -3 -27 → true
- Input 12 12 → false
- **Normal test:** Input 5 125 should print true.
- **Boundary or invalid test:** Test the input 0 0. The expected output is true because 0 * 0 * 0 == 0. The program does not validate whether multiplication overflows for very large integers.
- **One assumption or limitation:** The program assumes both inputs are valid integers. Very large values can overflow the range of Go's int type, causing an incorrect comparison.
- **What I changed in `main.go`:** I corrected the variable declarations and multiplication expression so the program can read two integers and compare the cube of x with y.
- **Why I believe the result is correct:** The program directly checks whether x*x*x == y, following the task's requirement without using a floating-point cube-root approximation.

## Exercise 09 — Advanced Bonus: Evolve a Population
- **Story in my words:** Instead of modifying only one candidate, the program creates a population of four-bit candidates and tries to find the target 1111 through repeated generations.
- **Input(s) and their types:** 
- populationSize = 6 (int)
- chromosomeSize = 4 (int)
- maxGenerations = 10 (int)
- mutationRate = 0.1 (float64)
- Each candidate is a string containing four bits.
- fitness returns an int.
- **Processing / logical rule:** The program generates a random population and calculates each candidate's fitness. It selects parents, performs crossover, and applies mutation to create a new population. The process repeats until it finds a candidate with fitness 4 or reaches the generation limit.
- **Expected output (before coding):** The program prints the generation number, the best candidate, and its fitness. If the target 1111 is found, it prints Optimal solution found!. The exact output varies because the population and mutations are random.
- **Actual output:** 
- Generation 1: Best = 1110, Fitness = 3
- Generation 2: Best = 1110, Fitness = 3
- Generation 3: Best = 1110, Fitness = 3
- Generation 4: Best = 1110, Fitness = 3
- Generation 5: Best = 1111, Fitness = 4
- Optimal solution found!
- **Normal test:** Run the program and check that it prints the generation number, the best candidate, and its fitness. The fitness value should be between 0 and 4.
- **Boundary or invalid test:** Test a run where the target 1111 is not found within the generation limit. The program should stop after at most 10 generations without printing Optimal solution found!. Because the candidates are random, this result is not guaranteed to happen in every run.
- **One assumption or limitation:** The program uses a small population and a maximum of ten generations. It does not guarantee finding 1111 within this limit because random selection, crossover, and mutation may not produce the target in time.
- **What I changed in `main.go`:** I corrected the Go type declarations and conversions and initialized the best candidate using an actual member of the population.
- **Why I believe the result is correct:** I corrected the Go type declarations and conversions and initialized the best candidate using an actual member of the population.

## Task 2 — Fill in the blank

Practice for [GitHub CLI intro](docs/github-cli-intro.md) and [Go basic types](docs/go-basic-types.md). Replace every `____` with the missing word, symbol or value. **Predict first, then check** each Go answer by running it. Write your answers in your fork's `EXPLANATION.md` under *Task 2 — Fill in the blank*. Answers are not published here.

### A — Git and GitHub (terminal)

**Word bank:** `add` · `clone` · `diff` · `fork` · `main` · `origin` · `push` · `status` · `switch`

1. A fork is your own online copy of someone else's repository.
2. `git clone  <url>` downloads a repository to your computer.
3. `git switch  -c task-2` creates a new branch and moves you onto it.
4. `git status` lists the files you changed, and `git diff` shows the exact lines you added or removed.
5. `git add  student/10_task2_types/main.go` puts a file in the box before you seal it with `git commit`.
6. `git push  -u origin task-2` sends your commits to your fork on GitHub.
7. `git log --oneline main..task-2` lists only the commits you made on your branch.
8. The remote that points to **your** fork is normally called origin.

### B — Go types, constants and operators

9. `7 / 2` is 3 because both numbers are integer values, so the fraction is dropped.
10. `7 % 2` is 1. The `%` operator gives the remainder of a division.
11. `float64(7) / float64(2)` is 3.5. Go does **not** convert types implicitly (automatically).
12. Go has no `char` type. A single character is stored as a rune (or as a byte for one ASCII byte).
13. A student ID such as `"00123"` should be a string because we never do arithmetic on it and the leading zeros must stay.
14. The keyword for a value that can never change is const. Reassigning it gives a compile-time error.
15. `true && false` is false, `true || false` is true, and `!true` is false.
16. In `x > 0 && y % x == 0`, when `x` is `0` the right side is not evaluated, so the program does not divide by zero.
17. In Go, `=` assigns a value and `==` compares two values.
18. `2 + 3 * 4` is 14 but `(2 + 3) * 4` is 20.
19. With `x, y := 0.1, 0.2`, the expression `x + y == 0.3` is false because `float64` stores some decimals only approximately.
20. `string(65)` gives "A", so use `strconv.Itoa(65)` when you want the text `"65"`.

### C - Complete the program
package main

import "fmt"

var taxRate = 0.1

func main() {
	price := 20000.0
	qty := 3
	subtotal := price * float64(qty)
	total := subtotal + subtotal*taxRate
	isBig := total > 50000
	fmt.Println(total, isBig)
}
- Fill the three blanks so the program compiles: const, float64, >
- Before running, predict the output: 66000 true
- Which line would stop the program from compiling if you removed the conversion blank? Explain in one sentence: The line subtotal := price * float64(qty) would cause a compilation error because Go cannot multiply a float64 value by an int value without type conversion.
